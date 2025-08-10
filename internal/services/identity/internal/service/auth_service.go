package service

import (
	"github.com/hthinh24/go-store/services/identity/internal/middleware"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/internal/pkg/rest"
	"github.com/hthinh24/go-store/services/identity"
	"github.com/hthinh24/go-store/services/identity/internal/config"
	"github.com/hthinh24/go-store/services/identity/internal/dto/request"
	"github.com/hthinh24/go-store/services/identity/internal/dto/response"
	"github.com/hthinh24/go-store/services/identity/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	logger         logger.Logger
	userRepository identity.UserRepository
	authRepository identity.AuthRepository
	config         *config.AppConfig
}

func NewAuthService(logger logger.Logger, userRepository identity.UserRepository, authRepository identity.AuthRepository, cfg *config.AppConfig) identity.AuthService {
	return &authService{
		logger:         logger,
		userRepository: userRepository,
		authRepository: authRepository,
		config:         cfg,
	}
}

func (a *authService) Login(request request.AuthRequest) (*response.AuthResponse, error) {
	user, err := a.userRepository.FindUserByEmail(request.Email)
	if err != nil {
		return &response.AuthResponse{}, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		a.logger.Error("Invalid password for user with email:", request.Email)
		return &response.AuthResponse{}, rest.AuthenticationError{}
	}

	token, err := a.generateToken(user)
	if err != nil {
		a.logger.Error("Error generating token for user:", user.Email, err)
		return &response.AuthResponse{}, err
	}

	a.logger.Info("User logged in successfully with email:", request.Email)
	return a.createAuthResponse(token)
}

func (a *authService) Verify(token string) (*response.VerifyResponse, error) {
	claims, err := a.validateToken(token)
	if err != nil {
		a.logger.Error("Failed to validate JWT token:", err)
		return nil, rest.AuthenticationError{}
	}

	return &response.VerifyResponse{
		UserID:      strconv.FormatInt(claims.UserID, 10),
		Roles:       claims.Roles,
		Permissions: claims.Permissions,
	}, nil
}

func (a *authService) Logout(token string) error {
	//TODO implement me
	panic("implement me")
}

func (a *authService) createAuthResponse(token string) (*response.AuthResponse, error) {
	return &response.AuthResponse{Token: token}, nil
}

func (a *authService) generateToken(user *entity.User) (string, error) {
	roles, err := a.authRepository.FindAllUserRolesByUserID(user.ID)
	if err != nil {
		a.logger.Error("Error fetching user roles for user ID %d: %v", user.ID, err)
		return "", err
	}

	//Extract role names and collect all permissions
	roleIDs := make([]int64, len(*roles))
	roleNames := make([]string, len(*roles))
	for i, role := range *roles {
		roleIDs[i] = role.ID
		roleNames[i] = role.Name
	}

	permissionList, err := a.authRepository.FindAllPermissionsByRoleIDs(roleIDs)
	if err != nil {
		return "", err
	}

	// Create a set to avoid duplicate permissions
	permissionSet := make(map[string]bool)
	for _, permission := range *permissionList {
		permissionSet[permission.Name] = true
	}

	// Convert permission set to slice
	permissions := make([]string, 0, len(permissionSet))
	for permission := range permissionSet {
		permissions = append(permissions, permission)
	}

	// Create JWT claims with configurable expiration
	claims := middleware.JWTClaims{
		UserID:      user.ID,
		Email:       user.Email,
		Roles:       roleNames,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.config.JWTExpiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.config.JWTSecret))
}

func (a *authService) validateToken(tokenString string) (*middleware.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &middleware.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// Convert claims to our custom JWTClaims type & check validity
	claims, ok := token.Claims.(*middleware.JWTClaims)
	if ok {
		return claims, nil
	}

	return nil, jwt.ErrTokenMalformed
}
