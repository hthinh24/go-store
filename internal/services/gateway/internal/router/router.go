package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/gateway/internal/config"
)

type Gateway struct {
	config *config.GatewayConfig
	logger logger.Logger
	client *http.Client
}

type VerifyResponse struct {
	UserID      string   `json:"user_id"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

func NewGateway(cfg *config.GatewayConfig, logger logger.Logger) *Gateway {
	return &Gateway{
		config: cfg,
		logger: logger,
		client: &http.Client{},
	}
}

func (g *Gateway) SetupRoutes(r *gin.Engine) {
	r.Any("/*path", g.handleRequest)
}

// Public endpoints that don't require authentication
var publicEndpoints = map[string][]string{
	"/api/v1/auth/login":    {"POST"},
	"/api/v1/auth/register": {"POST"},
	"/api/v1/auth/refresh":  {"POST"},
	"/api/v1/products/:id":  {"GET"}, // Public product details
	"/api/v1/products":      {"GET"},
}

func (g *Gateway) handleRequest(c *gin.Context) {
	path := c.Request.URL.Path
	method := c.Request.Method

	g.logger.Info("Received request, ", "path: ", path, " | ", "method: ", method)

	// Check if endpoint is public
	if g.isPublicEndpoint(path, method) {
		g.forwardToService(c, path)
		return
	}

	// For non-public endpoints, verify auth with identity service
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	// Verify token with identity service
	authResp, err := g.verifyWithIdentityService(authToken)
	if err != nil {
		g.logger.Error("Auth verification failed, error: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Add user info to headers
	c.Request.Header.Set("X-User-ID", authResp.UserID)
	c.Request.Header.Set("X-User-Roles", strings.Join(authResp.Roles, ","))
	c.Request.Header.Set("X-User-Permissions", strings.Join(authResp.Permissions, ","))

	// Forward to appropriate service
	g.forwardToService(c, path)
}

func (g *Gateway) isPublicEndpoint(path, method string) bool {
	for publicPath, methods := range publicEndpoints {
		if strings.HasPrefix(path, publicPath) {
			for _, acceptMethod := range methods {
				if acceptMethod == method {
					return true
				}
			}
		}
	}

	return false
}

func (g *Gateway) verifyWithIdentityService(authToken string) (*VerifyResponse, error) {
	req, err := http.NewRequest("GET",
		g.config.IdentityServiceURL+"/"+config.ApiVersionV1+"/auth/verify", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", authToken)

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("auth verification failed with status: %d", resp.StatusCode)
	}

	var verifyResponse VerifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&verifyResponse); err != nil {
		return nil, err
	}

	return &verifyResponse, nil
}

func (g *Gateway) forwardToService(c *gin.Context, path string) {
	var targetURL string

	// Route to appropriate service based on path
	if strings.HasPrefix(path, "/"+config.ApiVersionV1+"/auth") ||
		strings.HasPrefix(path, "/"+config.ApiVersionV1+"/users") {
		targetURL = g.config.IdentityServiceURL + path
	} else if strings.HasPrefix(path, "/"+config.ApiVersionV1+"/products") {
		targetURL = g.config.ProductServiceURL + path
	} else if strings.HasPrefix(path, "/"+config.ApiVersionV1+"/cart") {
		targetURL = g.config.CartServiceURL + path
	} else {
		g.logger.Warn("No service found for path", " path: ", path)
		c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	// Create new request
	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		g.logger.Error("Failed to create request", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Copy headers
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Forward request
	resp, err := g.client.Do(req)
	if err != nil {
		g.logger.Error("Failed to forward request", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service unavailable"})
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// Copy response
	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}
