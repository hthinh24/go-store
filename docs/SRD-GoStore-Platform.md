# Software Requirements Document (SRD)
# GoStore E-commerce Platform

**Version:** 1.0  
**Date:** August 16, 2025  
**Author:** Hung Thinh  
**Project Type:** E-commerce Microservices Platform

---

## 1. Introduction

### 1.1 Purpose
This Software Requirements Document (SRD) defines the technical specifications, system architecture, and implementation details for the GoStore e-commerce platform. It serves as a comprehensive guide for development, testing, and deployment of the microservices ecosystem.

### 1.2 Scope
The document covers all microservices in the GoStore platform:
- Identity Service (Authentication & User Management)
- Product Service (Catalog & SKU Management)
- Cart Service (Shopping Cart Management)
- Order Service (Order Processing)
- Gateway Service (API Gateway & Routing)
- Inventory Service (Stock Management)

### 1.3 Definitions and Acronyms
- **API**: Application Programming Interface
- **JWT**: JSON Web Token
- **RBAC**: Role-Based Access Control
- **SKU**: Stock Keeping Unit
- **ORM**: Object-Relational Mapping
- **CORS**: Cross-Origin Resource Sharing

---

## 2. System Architecture

### 2.1 Overall Architecture Pattern
**Microservices Architecture** with the following characteristics:
- **Service Independence**: Each service is independently deployable
- **Database per Service**: Each service owns its data
- **API Gateway Pattern**: Centralized entry point for client requests
- **Domain-Driven Design**: Services organized around business capabilities

### 2.2 Technology Stack

#### 2.2.1 Backend Technologies
```
Language: Go 1.24.3+
Web Framework: Gin (HTTP router/middleware)
Database: PostgreSQL 12+
ORM: GORM v2
Authentication: JWT with RS256/HS256
Password Hashing: bcrypt
Configuration: godotenv
Caching: Redis (planned)
```

#### 2.2.2 Development Tools
```
Dependency Management: Go Modules
Database Migration: Custom SQL scripts
API Documentation: Swagger/OpenAPI (planned)
Testing Framework: Go testing package + testify
Logging: Custom logger with structured output
```

### 2.3 Service Communication
- **Protocol**: HTTP/REST
- **Data Format**: JSON
- **Authentication**: JWT Bearer tokens
- **Error Handling**: Standardized error response format

---

## 3. Service Specifications

### 3.1 Identity Service

#### 3.1.1 Service Overview
**Port**: 8001  
**Database**: gostore_identity  
**Primary Responsibilities**:
- User authentication and authorization
- JWT token management
- Role-based access control
- User profile management

#### 3.1.2 Database Schema
```sql
-- Core Tables
users (id, email, password, first_name, last_name, phone_number, etc.)
roles (id, name, description)
permissions (id, name, description)
user_roles (user_id, role_id)
role_permissions (role_id, permission_id)
```

#### 3.1.3 Key APIs
```go
// Authentication
POST /api/auth/register
POST /api/auth/login
POST /api/auth/refresh
POST /api/auth/logout

// User Management
GET /api/users/profile
PUT /api/users/profile
PUT /api/users/password
GET /api/users/:id (admin only)

// Role Management (Admin)
GET /api/roles
POST /api/roles
PUT /api/roles/:id
DELETE /api/roles/:id
```

#### 3.1.4 Security Requirements
```go
// Password Policy
- Minimum 8 characters
- bcrypt hashing with cost factor 12
- Password history tracking (prevent reuse)

// JWT Configuration
- Algorithm: HS256 (configurable)
- Expiration: 24 hours (configurable)
- Refresh token: 7 days (configurable)
- Secret key: Environment variable

// Rate Limiting
- Login attempts: 5 per 15 minutes per IP
- Registration: 3 per hour per IP
```

#### 3.1.5 Data Models
```go
type User struct {
    BaseEntity
    Email        string    `json:"email" gorm:"unique;not null"`
    Password     string    `json:"-" gorm:"not null"`
    FirstName    string    `json:"first_name" gorm:"not null"`
    LastName     string    `json:"last_name" gorm:"not null"`
    PhoneNumber  string    `json:"phone_number"`
    DateOfBirth  time.Time `json:"date_of_birth"`
    Gender       string    `json:"gender"`
    Avatar       string    `json:"avatar"`
    Status       string    `json:"status" gorm:"default:ACTIVE"`
    // Provider fields for OAuth (future)
    ProviderID   string    `json:"provider_id"`
    ProviderName string    `json:"provider_name"`
}
```

### 3.2 Product Service

#### 3.2.1 Service Overview
**Port**: 8002  
**Database**: gostore_product  
**Primary Responsibilities**:
- Product catalog management
- Dynamic SKU generation
- Category and brand management
- Pricing and inventory tracking

#### 3.2.2 Database Schema
```sql
-- Core Tables
products (id, name, description, base_price, category_id, brand_id, etc.)
categories (id, name, description, parent_category_id)
brands (id, name, description)
product_skus (id, product_id, sku_code, price_modifier, stock_quantity)
product_options (id, name, type) -- size, color, material
product_option_values (id, option_id, value, display_order)
product_sku_options (sku_id, option_value_id)
product_attributes (id, product_id, attribute_name, attribute_value)
```

#### 3.2.3 SKU Generation Algorithm
```go
// SKU Format: {PRODUCT_PREFIX}-{CATEGORY_CODE}-{OPTION_COMBINATIONS}
// Example: TSHIRT-CLO-S-RED-COT (T-Shirt, Clothing, Small, Red, Cotton)

func GenerateSKU(product Product, optionCombination []OptionValue) string {
    prefix := strings.ToUpper(product.Name[:6]) // First 6 chars
    categoryCode := product.Category.Code       // 3-letter code
    
    var optionCodes []string
    for _, option := range optionCombination {
        optionCodes = append(optionCodes, option.Code)
    }
    
    return fmt.Sprintf("%s-%s-%s", 
        prefix, categoryCode, strings.Join(optionCodes, "-"))
}
```

#### 3.2.4 Key APIs
```go
// Product Management
GET /api/products?page=1&limit=20&category=&brand=&search=
POST /api/products
GET /api/products/:id
PUT /api/products/:id
DELETE /api/products/:id

// Category Management
GET /api/categories
POST /api/categories
GET /api/categories/:id/products

// Brand Management
GET /api/brands
POST /api/brands

// SKU Management
GET /api/products/:id/skus
POST /api/products/:id/skus
PUT /api/skus/:id
GET /api/skus/:id/inventory
```

#### 3.2.5 Product Data Model
```go
type Product struct {
    BaseEntity
    Name             string     `json:"name" gorm:"not null"`
    Description      string     `json:"description"`
    ShortDescription string     `json:"short_description"`
    ImageURL         string     `json:"image_url"`
    Slug             string     `json:"slug" gorm:"unique"`
    BasePrice        float64    `json:"base_price" gorm:"type:decimal(10,2)"`
    SalePrice        *float64   `json:"sale_price,omitempty"`
    SaleStartDate    *time.Time `json:"sale_start_date,omitempty"`
    SaleEndDate      *time.Time `json:"sale_end_date,omitempty"`
    IsFeatured       bool       `json:"is_featured" gorm:"default:false"`
    Status           string     `json:"status" gorm:"default:ACTIVE"`
    CategoryID       int64      `json:"category_id"`
    BrandID          int64      `json:"brand_id"`
    UserID           int64      `json:"user_id"`
    Version          int32      `json:"version" gorm:"default:1"`
}
```

### 3.3 Cart Service

#### 3.3.1 Service Overview
**Port**: 8003  
**Database**: gostore_cart  
**Primary Responsibilities**:
- Shopping cart management
- Cart newItem operations
- Price calculation
- Cart persistence

#### 3.3.2 Database Schema
```sql
-- Core Tables
carts (id, user_id, status, created_at, updated_at)
cart_items (id, cart_id, product_id, product_sku_id, quantity, unit_price, total_price)
```

#### 3.3.3 Key APIs
```go
// Cart Management
GET /api/cart
POST /api/cart/items
PUT /api/cart/items/:id
DELETE /api/cart/items/:id
DELETE /api/cart
GET /api/cart/summary

// Cart Operations
POST /api/cart/merge      // Merge guest cart with user cart
POST /api/cart/validate   // Validate cart before checkout
```

#### 3.3.4 Business Logic
```go
// Cart Item Calculation
func CalculateCartItemTotal(newItem CartItem) float64 {
    return newItem.UnitPrice * float64(newItem.Quantity)
}

// Cart Total Calculation
func CalculateCartTotal(cart Cart) CartSummary {
    var subtotal float64
    for _, newItem := range cart.Items {
        subtotal += CalculateCartItemTotal(newItem)
    }
    
    tax := subtotal * 0.08    // 8% tax rate
    shipping := 10.00         // Flat shipping rate
    total := subtotal + tax + shipping
    
    return CartSummary{
        Subtotal: subtotal,
        Tax:      tax,
        Shipping: shipping,
        Total:    total,
    }
}
```

### 3.4 Order Service (Implementation Specification)

#### 3.4.1 Service Overview
**Port**: 8004  
**Database**: gostore_order  
**Primary Responsibilities**:
- Order creation and management
- Order status tracking
- Order history
- Payment processing coordination

#### 3.4.2 Database Schema
```sql
-- Core Tables
orders (id, user_id, order_number, status, subtotal, tax, shipping, total, etc.)
order_items (id, order_id, product_id, sku_id, quantity, unit_price, total_price)
order_status_history (id, order_id, status, changed_by, changed_at, notes)
shipping_addresses (id, order_id, street, city, state, zip_code, country)
```

#### 3.4.3 Order State Machine
```go
// Order Status Flow
const (
    OrderStatusPending    = "PENDING"
    OrderStatusConfirmed  = "CONFIRMED"
    OrderStatusProcessing = "PROCESSING"
    OrderStatusShipped    = "SHIPPED"
    OrderStatusDelivered  = "DELIVERED"
    OrderStatusCancelled  = "CANCELLED"
    OrderStatusReturned   = "RETURNED"
)

// Valid Transitions
var OrderStatusTransitions = map[string][]string{
    OrderStatusPending:    {OrderStatusConfirmed, OrderStatusCancelled},
    OrderStatusConfirmed:  {OrderStatusProcessing, OrderStatusCancelled},
    OrderStatusProcessing: {OrderStatusShipped, OrderStatusCancelled},
    OrderStatusShipped:    {OrderStatusDelivered, OrderStatusReturned},
    OrderStatusDelivered:  {OrderStatusReturned},
}
```

### 3.5 Gateway Service

#### 3.5.1 Service Overview
**Port**: 8000  
**Primary Responsibilities**:
- Request routing to appropriate services
- Authentication middleware
- CORS handling
- Rate limiting
- Load balancing

#### 3.5.2 Routing Configuration
```go
// Service Discovery Configuration
type ServiceConfig struct {
    Identity ServiceEndpoint `json:"identity"`
    Product  ServiceEndpoint `json:"product"`
    Cart     ServiceEndpoint `json:"cart"`
    Order    ServiceEndpoint `json:"order"`
}

type ServiceEndpoint struct {
    Host string `json:"host"`
    Port int    `json:"port"`
    Path string `json:"path"`
}

// Route Mapping
var RouteMapping = map[string]ServiceEndpoint{
    "/api/auth/*":     services.Identity,
    "/api/users/*":    services.Identity,
    "/api/products/*": services.Product,
    "/api/categories/*": services.Product,
    "/api/brands/*":   services.Product,
    "/api/cart/*":     services.Cart,
    "/api/orders/*":   services.Order,
}
```

---

## 4. Data Models & Schemas

### 4.1 Base Entity
```go
type BaseEntity struct {
    ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
    CreatedBy string    `json:"created_by" gorm:"not null"`
    UpdatedBy string    `json:"updated_by" gorm:"not null"`
}
```

### 4.2 Standard Response Format
```go
type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   *ErrorInfo  `json:"error,omitempty"`
    Meta    *MetaInfo   `json:"meta,omitempty"`
}

type ErrorInfo struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}

type MetaInfo struct {
    Page       int `json:"page,omitempty"`
    Limit      int `json:"limit,omitempty"`
    Total      int `json:"total,omitempty"`
    TotalPages int `json:"total_pages,omitempty"`
}
```

### 4.3 Pagination Standard
```go
type PaginationRequest struct {
    Page  int `json:"page" form:"page" binding:"min=1"`
    Limit int `json:"limit" form:"limit" binding:"min=1,max=100"`
}

type PaginationResponse struct {
    Items      interface{} `json:"items"`
    Page       int         `json:"page"`
    Limit      int         `json:"limit"`
    Total      int64       `json:"total"`
    TotalPages int         `json:"total_pages"`
    HasNext    bool        `json:"has_next"`
    HasPrev    bool        `json:"has_prev"`
}
```

---

## 5. API Specifications

### 5.1 HTTP Status Codes
```
200 OK - Successful GET, PUT, PATCH
201 Created - Successful POST
204 No Content - Successful DELETE
400 Bad Request - Invalid request data
401 Unauthorized - Authentication required
403 Forbidden - Insufficient permissions
404 Not Found - Resource not found
409 Conflict - Resource conflict
422 Unprocessable Entity - Validation errors
500 Internal Server Error - Server errors
```

### 5.2 Error Response Format
```json
{
  "success": false,
  "message": "Validation failed",
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid input data",
    "details": "Email is required and must be valid"
  }
}
```

### 5.3 Authentication Header
```
Authorization: Bearer <JWT_TOKEN>
```

---

## 6. Security Specifications

### 6.1 JWT Token Structure
```json
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "user_id": "12345",
    "email": "user@example.com",
    "roles": ["customer"],
    "permissions": ["read:products", "write:cart"],
    "iat": 1692000000,
    "exp": 1692086400
  }
}
```

### 6.2 Input Validation Rules
```go
// User Registration Validation
type RegisterRequest struct {
    Email           string    `json:"email" binding:"required,email"`
    Password        string    `json:"password" binding:"required,min=8"`
    FirstName       string    `json:"first_name" binding:"required,min=2,max=50"`
    LastName        string    `json:"last_name" binding:"required,min=2,max=50"`
    PhoneNumber     string    `json:"phone_number" binding:"required,phone"`
    DateOfBirth     time.Time `json:"date_of_birth" binding:"required"`
    Gender          string    `json:"gender" binding:"required,oneof=MALE FEMALE OTHER"`
}

// Product Creation Validation
type CreateProductRequest struct {
    Name        string             `json:"name" binding:"required,min=3,max=255"`
    Description string             `json:"description" binding:"required,min=10"`
    BasePrice   float64            `json:"base_price" binding:"required,min=0"`
    CategoryID  int64              `json:"category_id" binding:"required"`
    BrandID     int64              `json:"brand_id" binding:"required"`
    Options     map[string][]string `json:"options"`
}
```

### 6.3 SQL Injection Prevention
```go
// Use GORM for all database operations
// Parameterized queries are automatically handled
db.Where("email = ?", email).First(&user)

// Input sanitization for search queries
func SanitizeSearchQuery(query string) string {
    // Remove special SQL characters
    query = strings.ReplaceAll(query, "'", "")
    query = strings.ReplaceAll(query, "\"", "")
    query = strings.ReplaceAll(query, ";", "")
    return strings.TrimSpace(query)
}
```

---

## 7. Performance Specifications

### 7.1 Response Time Requirements
```
Authentication APIs: < 200ms
Product Listing: < 300ms
Product Search: < 500ms
Cart Operations: < 200ms
Order Creation: < 1000ms
```

### 7.2 Database Optimization
```sql
-- Essential Indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_brand_id ON products(brand_id);
CREATE INDEX idx_products_status ON products(status);
CREATE INDEX idx_product_skus_product_id ON product_skus(product_id);
CREATE INDEX idx_cart_items_cart_id ON cart_items(cart_id);
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_status ON orders(status);

-- Composite Indexes for Common Queries
CREATE INDEX idx_products_category_status ON products(category_id, status);
CREATE INDEX idx_products_search ON products USING gin(to_tsvector('english', name || ' ' || description));
```

### 7.3 Caching Strategy (Planned)
```go
// Redis Caching Configuration
type CacheConfig struct {
    ProductCache    CacheSettings `json:"product_cache"`
    CategoryCache   CacheSettings `json:"category_cache"`
    UserCache       CacheSettings `json:"user_cache"`
}

type CacheSettings struct {
    TTL     time.Duration `json:"ttl"`
    Enabled bool          `json:"enabled"`
}

// Cache Keys
const (
    ProductCacheKey   = "product:%d"
    CategoryCacheKey  = "category:%d"
    UserProfileKey    = "user:profile:%d"
    CartCacheKey      = "cart:user:%d"
)
```

---

## 8. Testing Specifications

### 8.1 Unit Testing Standards
```go
// Test Coverage Requirements
- Service Layer: 90%+ coverage
- Repository Layer: 85%+ coverage
- Handler Layer: 80%+ coverage

// Test Naming Convention
func TestServiceName_MethodName_Scenario(t *testing.T) {
    // Arrange
    // Act
    // Assert
}

// Example Unit Test
func TestUserService_CreateUser_ValidInput_ReturnsUser(t *testing.T) {
    // Arrange
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)
    request := CreateUserRequest{
        Email:     "test@example.com",
        Password:  "password123",
        FirstName: "John",
        LastName:  "Doe",
    }
    
    // Act
    user, err := service.CreateUser(request)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, request.Email, user.Email)
}
```

### 8.2 Integration Testing
```go
// Database Integration Tests
func TestUserRepository_Create_ValidUser_SavesSuccessfully(t *testing.T) {
    // Setup test database
    db := setupTestDB(t)
    defer teardownTestDB(t, db)
    
    repo := NewUserRepository(db)
    user := &User{
        Email:     "integration@test.com",
        FirstName: "Integration",
        LastName:  "Test",
    }
    
    err := repo.Create(user)
    
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)
}
```

### 8.3 API Testing
```go
// HTTP Handler Tests
func TestAuthHandler_Register_ValidRequest_Returns201(t *testing.T) {
    // Setup
    router := setupTestRouter()
    
    requestBody := `{
        "email": "test@example.com",
        "password": "password123",
        "first_name": "Test",
        "last_name": "User"
    }`
    
    req := httptest.NewRequest("POST", "/api/auth/register", 
        strings.NewReader(requestBody))
    req.Header.Set("Content-Type", "application/json")
    
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    assert.Equal(t, http.StatusCreated, w.Code)
    
    var response APIResponse
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.True(t, response.Success)
}
```

---

## 9. Deployment Specifications

### 9.1 Environment Configuration
```bash
# Development Environment
PORT=8001
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=gostore_identity_dev
DB_PORT=5432
JWT_SECRET=development-secret-key
LOG_LEVEL=debug

# Production Environment
PORT=8001
DB_HOST=prod-db-host
DB_USER=app_user
DB_PASSWORD=${DB_PASSWORD}  # From environment
DB_NAME=gostore_identity
DB_PORT=5432
JWT_SECRET=${JWT_SECRET}    # From environment
LOG_LEVEL=info
```

### 9.2 Docker Configuration (Planned)
```dockerfile
# Dockerfile template for services
FROM golang:1.24.3-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"]
```

### 9.3 Database Migration Strategy
```go
// Migration Management
type Migration struct {
    Version     string    `json:"version"`
    Description string    `json:"description"`
    Script      string    `json:"script"`
    Applied     bool      `json:"applied"`
    AppliedAt   time.Time `json:"applied_at"`
}

// Migration Runner
func RunMigrations(db *gorm.DB) error {
    migrations := loadMigrationFiles()
    
    for _, migration := range migrations {
        if !migration.Applied {
            if err := executeMigration(db, migration); err != nil {
                return fmt.Errorf("failed to apply migration %s: %w", 
                    migration.Version, err)
            }
        }
    }
    
    return nil
}
```

---

## 10. Monitoring & Logging

### 10.1 Logging Standards
```go
// Structured Logging Format
type LogEntry struct {
    Timestamp string            `json:"timestamp"`
    Level     string            `json:"level"`
    Service   string            `json:"service"`
    Message   string            `json:"message"`
    Fields    map[string]interface{} `json:"fields,omitempty"`
    Error     string            `json:"error,omitempty"`
    TraceID   string            `json:"trace_id,omitempty"`
}

// Usage Example
logger.Info("User created successfully", 
    logger.Fields{
        "user_id": user.ID,
        "email":   user.Email,
        "action":  "user_creation",
    })
```

### 10.2 Health Check Endpoints
```go
// Health Check Response
type HealthResponse struct {
    Status    string                 `json:"status"`
    Service   string                 `json:"service"`
    Version   string                 `json:"version"`
    Timestamp time.Time              `json:"timestamp"`
    Checks    map[string]HealthCheck `json:"checks"`
}

type HealthCheck struct {
    Status  string        `json:"status"`
    Message string        `json:"message,omitempty"`
    Latency time.Duration `json:"latency,omitempty"`
}

// Endpoint: GET /health
func HealthCheckHandler(c *gin.Context) {
    checks := map[string]HealthCheck{
        "database": checkDatabase(),
        "cache":    checkCache(),
    }
    
    status := "healthy"
    for _, check := range checks {
        if check.Status != "healthy" {
            status = "unhealthy"
            break
        }
    }
    
    response := HealthResponse{
        Status:    status,
        Service:   "identity-service",
        Version:   "1.0.0",
        Timestamp: time.Now(),
        Checks:    checks,
    }
    
    c.JSON(http.StatusOK, response)
}
```

---

## 11. Future Enhancements

### 11.1 Scalability Improvements
- **Event-Driven Architecture**: Implement message queues (RabbitMQ/Apache Kafka)
- **CQRS Pattern**: Separate read/write operations for better performance
- **Database Sharding**: Horizontal partitioning for large datasets
- **Microservice Mesh**: Service mesh for advanced traffic management

### 11.2 Additional Features
- **Real-time Notifications**: WebSocket implementation
- **Advanced Search**: Elasticsearch integration
- **File Storage**: S3-compatible object storage
- **Analytics**: Data warehouse and reporting system
- **Mobile API**: GraphQL endpoint for mobile applications

### 11.3 DevOps Enhancements
- **CI/CD Pipeline**: Automated testing and deployment
- **Container Orchestration**: Kubernetes deployment
- **Infrastructure as Code**: Terraform configurations
- **Monitoring Stack**: Prometheus + Grafana
- **Distributed Tracing**: Jaeger implementation

---

*This SRD provides the technical foundation for implementing the GoStore e-commerce platform and should be updated as the system evolves and new requirements emerge.*
