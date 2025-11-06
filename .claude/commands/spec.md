
# Backend Specification - Inventory & Ticketing Management System

## 1. Overview
This document outlines the technical specifications for the backend of the Inventory & Ticketing Management System. The backend is built using Go, Gin framework, GORM for ORM, and PostgreSQL as the database.

## 2. Technology Stack
- **Language**: Go 1.22+
- **Framework**: Gin-Gonic
- **ORM**: GORM
- **Database**: PostgreSQL
- **Authentication**: JWT
- **Configuration**: ENV (.env)
- **Migration**: Golang Migrate

## 3. Project Structure
```
project_root/
├── cmd/
│   └── app/
│       └── main.go
│   └── migration/
│       └── main.go
├── domain/
│   ├── entity/
│   │   ├── user.go
│   │   ├── asset.go
│   │   ├── ticket.go
│   │   └── location.go
│   ├── valueobject/
│   │   ├── user_role.go
│   │   ├── asset_status.go
│   │   ├── asset_type.go
│   │   ├── ticket_severity.go
│   │   └── ticket_status.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   ├── asset_repository.go
│   │   ├── ticket_repository.go
│   │   └── location_repository.go
│   ├── service/
│   │   ├── auth_service.go
│   │   ├── asset_service.go
│   │   ├── ticket_service.go
│   │   └── location_service.go
│   ├── event/
│   │   ├── user_created.go
│   │   ├── asset_created.go
│   │   ├── ticket_created.go
│   │   └── location_created.go
│   └── enum/
│       ├── user_role.go
│       ├── asset_status.go
│       ├── asset_type.go
│       ├── ticket_severity.go
│       └── ticket_status.go
├── application/
│   ├── usecase/
│   │   ├── auth/
│   │   │   ├── login.go
│   │   │   └── register.go
│   │   ├── asset/
│   │   │   ├── create_asset.go
│   │   │   ├── get_asset.go
│   │   │   ├── list_assets.go
│   │   │   ├── update_asset.go
│   │   │   └── delete_asset.go
│   │   ├── ticket/
│   │   │   ├── create_ticket.go
│   │   │   ├── get_ticket.go
│   │   │   ├── list_tickets.go
│   │   │   ├── update_ticket.go
│   │   │   └── delete_ticket.go
│   │   └── location/
│   │       ├── create_location.go
│   │       ├── get_location.go
│   │       ├── list_locations.go
│   │       ├── update_location.go
│   │       └── delete_location.go
│   ├── repository/
│   │   ├── user_repository_impl.go
│   │   ├── asset_repository_impl.go
│   │   ├── ticket_repository_impl.go
│   │   └── location_repository_impl.go
│   ├── service/
│   │   ├── auth_service_impl.go
│   │   ├── asset_service_impl.go
│   │   ├── ticket_service_impl.go
│   │   └── location_service_impl.go
│   └── dto/
│       ├── auth/
│       │   ├── login_request.go
│       │   └── login_response.go
│       ├── asset/
│       │   ├── asset_request.go
│       │   └── asset_response.go
│       ├── ticket/
│       │   ├── ticket_request.go
│       │   └── ticket_response.go
│       └── location/
│           ├── location_request.go
│           └── location_response.go
├── delivery/
│   └── http/
│       ├── handler/
│       │   ├── auth_handler.go
│       │   ├── asset_handler.go
│       │   ├── ticket_handler.go
│       │   └── location_handler.go
│       ├── middleware/
│       │   ├── auth.go
│       │   ├── logging.go
│       │   └── error.go
│       └── router.go
├── infrastructure/
│   ├── config/
│   │   └── config.go
│   ├── persistence/
│   │   └── postgres/
│   │       ├── user_repository.go
│   │       ├── asset_repository.go
│   │       ├── ticket_repository.go
│   │       └── location_repository.go
│   └── jwt/
│       └── jwt.go
├── migrations/
│   ├── 000001_init_schema.up.sql
│   └── 000002_seed_data.up.sql
├── pkg/
│   ├── database/
│   │   └── database.go
│   ├── context/
│   │   └── context.go
│   └── common/
│       ├── errors.go
│       ├── response.go
│       └── validator.go
├── go.mod
├── go.sum
└── .env.example
```

## 4. Models

### User Model
```go
type User struct {
    ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Name         string    `json:"name" gorm:"not null"`
    Email        string    `json:"email" gorm:"unique;not null"`
    PasswordHash string    `json:"-" gorm:"not null"`
    Role         string    `json:"role" gorm:"not null;check:role IN ('admin', 'employee')"`
    CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
```

### Asset Model
```go
type Asset struct {
    ID            uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    UniqueID      string     `json:"unique_id" gorm:"not null"`
    Name          string     `json:"name" gorm:"not null"`
    Comment       string     `json:"comment"`
    Detail        string     `json:"detail"`
    Qty           int        `json:"qty" gorm:"default:1"`
    Brand         string     `json:"brand"`
    Type          string     `json:"type" gorm:"check:type IN ('it', 'non_it')"`
    Status        string     `json:"status" gorm:"default:'available';check:status IN ('available', 'booked', 'broken', 'repair')"`
    Category      string     `json:"category"`
    LocationID    *uuid.UUID `json:"location_id" gorm:"type:uuid"`
    LocationLabel string     `json:"location_label"`
    CreatedAt     time.Time  `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt     time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
```

### Ticket Model
```go
type Ticket struct {
    ID          uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    AssetID     uuid.UUID  `json:"asset_id" gorm:"type:uuid;not null"`
    Category    string     `json:"category" gorm:"not null"`
    Severity    string     `json:"severity" gorm:"check:severity IN ('low', 'medium', 'high', 'critical')"`
    Duration    string     `json:"duration"`
    DueDate     time.Time  `json:"due_date"`
    Reporting   uuid.UUID  `json:"reporting" gorm:"type:uuid;not null"`
    AssignedTo  *uuid.UUID `json:"assigned_to" gorm:"type:uuid"`
    Comment     string     `json:"comment"`
    Status      string     `json:"status" gorm:"default:'open';check:status IN ('open', 'in_progress', 'resolved', 'closed')"`
    CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
```

### Location Model
```go
type Location struct {
    ID          uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Name        string     `json:"name" gorm:"not null"`
    Area        string     `json:"area" gorm:"not null"`
    Description string     `json:"description"`
    Capacity    int        `json:"capacity" gorm:"default:0"`
    CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
```

## 5. Routes

| Path           | Method | Handler           | Description                 | Access Level |
| -------------- | ------ | ----------------- | --------------------------- | ------------ |
| `/auth/login`  | POST   | AuthHandler.Login | User authentication         | Public       |
| `/assets`      | GET    | AssetHandler.List | List all assets             | All          |
| `/assets`      | POST   | AssetHandler.Create | Create new asset           | Admin        |
| `/assets/:id`  | GET    | AssetHandler.Get  | Get asset details           | All          |
| `/assets/:id`  | PUT    | AssetHandler.Update | Update asset              | Admin        |
| `/assets/:id`  | DELETE | AssetHandler.Delete | Delete asset              | Admin        |
| `/tickets`     | GET    | TicketHandler.List | List all tickets            | All          |
| `/tickets`     | POST   | TicketHandler.Create | Create new ticket         | All          |
| `/tickets/:id` | GET    | TicketHandler.Get  | Get ticket details         | All          |
| `/tickets/:id` | PUT    | TicketHandler.Update | Update ticket             | Admin        |
| `/tickets/:id` | DELETE | TicketHandler.Delete | Delete ticket             | Admin        |
| `/locations`   | GET    | LocationHandler.List | List all locations         | All          |
| `/locations`   | POST   | LocationHandler.Create | Create new location      | Admin        |
| `/locations/:id`| GET    | LocationHandler.Get  | Get location details      | All          |
| `/locations/:id`| PUT    | LocationHandler.Update | Update location          | Admin        |
| `/locations/:id`| DELETE | LocationHandler.Delete | Delete location          | Admin        |

## 6. Middleware

### Authentication Middleware
```go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }
        
        tokenString := strings.TrimPrefix(token, "Bearer ")
        claims, err := jwt.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        c.Set("user_id", claims.UserID)
        c.Set("user_role", claims.Role)
        c.Next()
    }
}
```

### Role-based Authorization Middleware
```go
func RoleMiddleware(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("user_role")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
            c.Abort()
            return
        }
        
        role := userRole.(string)
        authorized := false
        for _, r := range roles {
            if r == role {
                authorized = true
                break
            }
        }
        
        if !authorized {
            c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

## 7. Database Schema

### Users Table
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'employee')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Assets Table
```sql
CREATE TABLE assets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    unique_id VARCHAR(100) NOT NULL,
    name VARCHAR(255) NOT NULL,
    comment TEXT,
    detail TEXT,
    qty INTEGER DEFAULT 1,
    brand VARCHAR(100),
    type VARCHAR(20) NOT NULL CHECK (type IN ('it', 'non_it')),
    status VARCHAR(20) DEFAULT 'available' CHECK (status IN ('available', 'booked', 'broken', 'repair')),
    category VARCHAR(100),
    location_id UUID REFERENCES locations(id),
    location_label VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Tickets Table
```sql
CREATE TABLE tickets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID NOT NULL REFERENCES assets(id),
    category VARCHAR(100) NOT NULL,
    severity VARCHAR(20) NOT NULL CHECK (severity IN ('low', 'medium', 'high', 'critical')),
    duration VARCHAR(50),
    due_date DATE,
    reporting UUID NOT NULL REFERENCES users(id),
    assigned_to UUID REFERENCES users(id),
    comment TEXT,
    status VARCHAR(20) DEFAULT 'open' CHECK (status IN ('open', 'in_progress', 'resolved', 'closed')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Locations Table
```sql
CREATE TABLE locations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    area VARCHAR(100) NOT NULL,
    description TEXT,
    capacity INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

## 8. Seed Data

```sql
-- Seed locations
INSERT INTO locations (name, area, description, capacity) VALUES
('Smart Solution', '1st Floor', 'Located on the 1st Floor', 8),
('Technology', '1st Floor', 'Located on the 1st Floor', 4),
('Integrity', '2nd Floor', 'Located on the 2nd Floor', 6),
('Innovation', '3rd Floor', 'Located on the 3rd Floor', 8),
('Loyalty', '3rd Floor', 'Located on the 3rd Floor', 6),
('Quality', '3rd Floor', 'Located on the 3rd Floor', 6),
('Team Work (Open Area)', '3rd Floor', 'Located on the 3rd Floor', 50),
('Excellent', '4th Floor', 'Located on the 4th Floor', 4),
('Open Communication', '4th Floor', 'Located on the 4th Floor', 8),
('General', 'Outside Floor', 'General area', 0);

-- Seed admin user
INSERT INTO users (name, email, password_hash, role) VALUES
('Admin User', 'admin@company.com', '$2a$10$hashedpasswordhere', 'admin');
```

## 9. Deployment

### Dockerfile
```dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env

EXPOSE 8080
CMD ["./main"]
```

### Docker Compose
```yaml
version: '3.8'

services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=postgres
      - DB_USER=postgres
      - DB_PASSWORD=postgres
      - DB_NAME=inventory_db
      - JWT_SECRET=your-jwt-secret
    depends_on:
      - postgres
    networks:
      - app-network

  postgres:
    image: postgres:15
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=inventory_db
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"
    networks:
      - app-network

volumes:
  postgres_data:

networks:
  app-network:
    driver: bridge
```

---