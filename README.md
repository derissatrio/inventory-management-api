# Inventory & Ticketing Management System

A comprehensive inventory and ticketing management system built with Go, Gin framework, GORM, and PostgreSQL.

## Features

- **Asset Management**: Create, read, update, and delete assets with categories, locations, and status tracking
- **Ticket Management**: Create and manage tickets for assets with severity levels and status tracking
- **User Management**: Role-based authentication (admin/employee) with JWT tokens
- **Location Management**: Manage physical locations with capacity tracking
- **RESTful API**: Clean API following REST principles
- **PostgreSQL Database**: Robust database with proper relationships and constraints

## Tech Stack

- **Backend**: Go 1.22+, Gin-Gonic framework
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT (JSON Web Tokens)
- **Architecture**: Clean Architecture with Domain-Driven Design (DDD)
- **Containerization**: Docker and Docker Compose

## Project Structure

```
├── cmd/                          # Application entry points
│   ├── app/main.go              # Main application
│   └── migration/main.go        # Database migration tool
├── domain/                      # Domain layer (entities, repositories, services)
│   ├── entity/                  # Domain entities
│   ├── enum/                    # Domain enums
│   ├── repository/              # Repository interfaces
│   └── service/                 # Service interfaces
├── application/                 # Application layer (use cases, DTOs)
│   ├── dto/                     # Data Transfer Objects
│   ├── repository/              # Repository implementations
│   ├── service/                 # Service implementations
│   └── usecase/                 # Use cases
├── delivery/                    # Presentation layer
│   └── http/                    # HTTP handlers and middleware
├── infrastructure/              # Infrastructure layer
│   ├── config/                  # Configuration
│   └── jwt/                     # JWT implementation
├── pkg/                         # Shared packages
│   ├── database/                # Database utilities
│   └── common/                  # Common utilities
├── migrations/                  # SQL migration files
├── Dockerfile                   # Docker configuration
├── docker-compose.yml           # Docker Compose configuration
└── .env.example                 # Environment variables template
```

## API Endpoints

### Authentication
- `POST /api/v1/auth/login` - User login

### Assets
- `GET /api/v1/assets` - List all assets (with filtering and pagination)
- `POST /api/v1/assets` - Create new asset (admin only)
- `GET /api/v1/assets/{id}` - Get asset details
- `PUT /api/v1/assets/{id}` - Update asset (admin only)
- `DELETE /api/v1/assets/{id}` - Delete asset (admin only)

### Tickets
- `GET /api/v1/tickets` - List all tickets (with filtering and pagination)
- `POST /api/v1/tickets` - Create new ticket
- `GET /api/v1/tickets/{id}` - Get ticket details
- `PUT /api/v1/tickets/{id}` - Update ticket (admin only)
- `DELETE /api/v1/tickets/{id}` - Delete ticket (admin only)

### Locations
- `GET /api/v1/locations` - List all locations
- `POST /api/v1/locations` - Create new location (admin only)
- `GET /api/v1/locations/{id}` - Get location details
- `PUT /api/v1/locations/{id}` - Update location (admin only)
- `DELETE /api/v1/locations/{id}` - Delete location (admin only)

### Users
- `GET /api/v1/users/me` - Get current user profile

### Health Check
- `GET /api/v1/health` - Health check endpoint

## Getting Started

### Prerequisites

- Go 1.22 or higher
- PostgreSQL 12 or higher
- Docker and Docker Compose (optional)

### Using Docker Compose (Recommended)

1. Clone the repository:
```bash
git clone <repository-url>
cd inventory-ticketing-system
```

2. Run with Docker Compose:
```bash
docker-compose up -d
```

This will start:
- The API server on port 8080
- PostgreSQL database on port 5432
- Adminer (database admin tool) on port 8081

### Manual Setup

1. Install dependencies:
```bash
go mod download
```

2. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your database configuration
```

3. Run database migrations:
```bash
go run cmd/migration/main.go
```

4. Run the application:
```bash
go run cmd/app/main.go
```

Or build and run:
```bash
go build -o main ./cmd/app
./main
```

## Default Credentials

The system comes with a default admin user:
- **Email**: admin@company.com
- **Password**: admin123

## API Documentation

### Example Requests

#### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@company.com",
    "password": "admin123"
  }'
```

#### Create Asset (Admin only, requires JWT token)
```bash
curl -X POST http://localhost:8080/api/v1/assets \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "unique_id": "AST001",
    "name": "Dell Latitude 5420",
    "category": "Laptop",
    "type": "it",
    "qty": 10,
    "brand": "Dell",
    "detail": "Intel i7, 16GB RAM, 512GB SSD"
  }'
```

#### List Assets
```bash
curl -X GET "http://localhost:8080/api/v1/assets?limit=10&offset=0&jenis=it" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

#### Create Ticket
```bash
curl -X POST http://localhost:8080/api/v1/tickets \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "assetId": "YOUR_ASSET_ID",
    "kategori": "Repair",
    "severity": "high",
    "comment": "Laptop screen is flickering"
  }'
```

## Development

### Running Tests
```bash
go test ./...
```

### Database Migrations
The application uses GORM auto-migration, which creates/updates the database schema automatically on startup.

### Environment Variables
- `SERVER_PORT`: HTTP server port (default: 8080)
- `DB_HOST`: PostgreSQL host (default: localhost)
- `DB_PORT`: PostgreSQL port (default: 5432)
- `DB_USER`: PostgreSQL username (default: postgres)
- `DB_PASSWORD`: PostgreSQL password (default: postgres)
- `DB_NAME`: PostgreSQL database name (default: inventory_db)
- `JWT_SECRET`: JWT secret key (change this in production)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.