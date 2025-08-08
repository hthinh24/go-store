# GoStore - E-commerce Microservices Platform

A modern, scalable e-commerce platform built with Go microservices architecture, featuring comprehensive product management, user authentication, and robust data modeling.

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                           Client                                │
└─────────────────────┬───────────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────────┐
│                      API Gateway (Nginx)                       │
└─┬─────────┬─────────┬─────────┬─────────┬─────────┬─────────────┘
  │         │         │         │         │         │
┌─▼──────┐ ┌▼──────┐ ┌▼──────┐ ┌▼──────┐ ┌▼──────┐ ┌▼──────┐
│Identity│ │  Cart │ │Product│ │Inventory│ │ Order │ │Future │
│Service │ │Service│ │Service│ │ Service │ │Service│ │Services│
└─┬──────┘ └┬──────┘ └┬──────┘ └┬───────┘ └┬──────┘ └───────┘
  │         │         │         │          │
┌─▼──────┐ ┌▼──────┐ ┌▼──────┐ ┌▼───────┐ ┌▼──────┐
│Identity│ │ Cart  │ │Product│ │Inventory│ │ Order │
│   DB   │ │  DB   │ │  DB   │ │   DB    │ │  DB   │
│(Postgres)│(Postgres)│(Postgres)│(Postgres)│(Postgres)
└────────┘ └───────┘ └───────┘ └────────┘ └───────┘
```

## 🚀 Services

### Identity Service
- **Purpose**: User authentication, authorization, and profile management
- **Features**:
  - JWT-based authentication
  - Role-based access control (RBAC)
  - User registration and login
  - Password management
  - Profile updates

### Product Service
- **Purpose**: Product catalog management with complex SKU handling
- **Features**:
  - Product CRUD operations
  - Dynamic SKU generation with multiple options (size, color, material, etc.)
  - Category and brand management
  - Product attributes and options system
  - Sale price management with time-based validity
  - Bulk product seeding capabilities

## 🛠️ Tech Stack

- **Language**: Go 1.24.3
- **Framework**: Gin (HTTP router)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT
- **Password Hashing**: bcrypt
- **Configuration**: godotenv

## 📁 Project Structure

```
go-store/
├── internal/
│   ├── pkg/                    # Shared packages
│   │   ├── entity/            # Base entities
│   │   ├── logger/            # Logging utilities
│   │   └── rest/              # REST utilities & error handling
│   └── services/
│       ├── identity/          # Authentication & user management
│       │   ├── cmd/           # Service entry point
│       │   ├── db/            # Database schemas & seed data
│       │   └── internal/      # Service-specific logic
│       └── product/           # Product catalog management
│           ├── cmd/           # Service entry point & seeder
│           ├── db/            # Database schemas & seed data
│           └── internal/      # Service-specific logic
```

## 🏃‍♂️ Getting Started

### Prerequisites
- Go 1.24.3 or higher
- PostgreSQL 12+
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/hthinh24/go-store.git
   cd go-store
   ```

2. **Setup databases**
   ```bash
   # Create databases for each service
   createdb gostore_identity
   createdb gostore_product
   ```

3. **Environment configuration**
   
   Create `.env` files for each service:
   
   **Identity Service** (`internal/services/identity/.env`):
   ```env
   PORT=8001
   DB_HOST=localhost
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=gostore_identity
   DB_PORT=5432
   JWT_SECRET=your_jwt_secret_key
   ```
   
   **Product Service** (`internal/services/product/.env`):
   ```env
   PORT=8002
   DB_HOST=localhost
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=gostore_product
   DB_PORT=5432
   ```

4. **Database setup**
   ```bash
   # Identity Service
   cd internal/services/identity
   psql -d gostore_identity -f db/shema.sql
   psql -d gostore_identity -f db/roles_permissions_data.sql
   
   # Product Service
   cd ../product
   psql -d gostore_product -f db/schemaV2.sql
   psql -d gostore_product -f db/seed_data.sql
   ```

5. **Install dependencies**
   ```bash
   # Identity Service
   cd internal/services/identity
   go mod tidy
   
   # Product Service
   cd ../product
   go mod tidy
   ```

## 🔥 Running the Services

### Identity Service
```bash
cd internal/services/identity
go run cmd/main.go
# Service will start on port 8001
```

### Product Service
```bash
cd internal/services/product
go run cmd/main.go
# Service will start on port 8002
```

### Product Seeding
The product service includes powerful seeding capabilities:

```bash
cd internal/services/product

# Seed 50 random products
make seed-random COUNT=50

# Seed 20 men's clothing items
make seed-mens COUNT=20

# Seed diverse product mix
make seed-diverse COUNT=30

# Seed in batch mode
make seed-batch COUNT=100
```

## 📊 Database Schema Highlights

### Product Schema Features
- **Complex SKU System**: Automatic SKU generation based on product options
- **Dynamic Pricing**: Base price + SKU-specific price modifiers
- **Sale Management**: Time-based sale prices with start/end dates
- **Rich Attributes**: Flexible product attributes and options system
- **Inventory Tracking**: Stock management per SKU
- **SEO-Friendly**: Automatic slug generation with conflict resolution

### Identity Schema Features
- **Role-Based Access**: Hierarchical permission system
- **Secure Authentication**: Bcrypt password hashing + JWT tokens
- **User Profiles**: Comprehensive user information management

## 🔌 API Endpoints

### Identity Service (Port 8001)
```
POST   /api/auth/register     # User registration
POST   /api/auth/login        # User login
POST   /api/auth/refresh      # Refresh JWT token
GET    /api/users/profile     # Get user profile
PUT    /api/users/profile     # Update user profile
PUT    /api/users/password    # Change password
```

### Product Service (Port 8002)
```
GET    /api/products          # List products (with pagination)
POST   /api/products          # Create product (with auto-SKU generation)
GET    /api/products/:id      # Get product details
PUT    /api/products/:id      # Update product
DELETE /api/products/:id      # Delete product
GET    /api/categories        # List categories
GET    /api/brands           # List brands
```

## 🧪 Testing & Development

### Sample Data
The project includes comprehensive seed data:
- **Categories**: Electronics, Fashion, Home & Garden, Sports, Books, etc.
- **Brands**: Nike, Apple, Samsung, Zara, IKEA, etc.
- **Attributes**: Size, Color, Material, Storage, RAM, etc.
- **Product Variations**: Automatic SKU generation for all combinations

### Product Creation Example
```json
{
  "name": "Premium Cotton T-Shirt",
  "description": "High-quality cotton t-shirt",
  "base_price": 29.99,
  "category_id": 1,
  "brand_id": 1,
  "option_values": {
    "1": ["S", "M", "L", "XL"],
    "2": ["Red", "Blue", "Black"],
    "3": ["Cotton", "Blend"]
  }
}
```
This will automatically generate 24 SKUs (4×3×2) with unique identifiers.

## 🔧 Configuration

### Environment Variables
- `PORT`: Service port number
- `DB_*`: Database connection parameters
- `JWT_SECRET`: Secret key for JWT token signing (Identity Service)

### Database Configuration
- PostgreSQL with GORM ORM
- Automatic migrations on startup
- Connection pooling and optimization

## 🚧 Roadmap

- [ ] **Cart Service**: Shopping cart management
- [ ] **Order Service**: Order processing and management
- [ ] **Inventory Service**: Advanced inventory tracking
- [ ] **Payment Service**: Payment processing integration
- [ ] **Notification Service**: Email and SMS notifications
- [ ] **Search Service**: Elasticsearch integration
- [ ] **File Storage**: Image and media management
- [ ] **API Gateway**: Centralized routing and rate limiting

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👨‍💻 Author

**Hung Thinh** - [hthinh24](https://github.com/hthinh24)

---

*Building a comprehensive e-commerce platform with modern microservices architecture. Each service is designed to be scalable, maintainable, and production-ready.*
