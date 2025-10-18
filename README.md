# ⚽ Football Management API

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go)
![Gin Framework](https://img.shields.io/badge/Gin-v1.10.0-00ADD8?style=for-the-badge&logo=go)
![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

**A comprehensive REST API for managing football teams, players, matches, and statistics**

Built with Go, Gin Framework, and Clean Architecture principles

[Features](#-features) • [Installation](#-installation) • [API Documentation](#-api-documentation) • [Usage](#-usage) • [Contributing](#-contributing)

</div>

---

## 📋 Table of Contents

- [Overview](#-overview)
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Architecture](#-architecture)
- [Prerequisites](#-prerequisites)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [Running the Application](#-running-the-application)
- [API Documentation](#-api-documentation)
- [Database Schema](#-database-schema)
- [Project Structure](#-project-structure)
- [Testing](#-testing)
- [Deployment](#-deployment)
- [Contributing](#-contributing)
- [License](#-license)

---

## 🎯 Overview

**Football Management API** is a robust and scalable REST API designed to manage football organizations, including teams, players, matches, schedules, and comprehensive statistics. Built with modern Go practices and Clean Architecture, this API provides a solid foundation for football management applications.

### Key Highlights

- 🏗️ **Clean Architecture** - Separation of concerns with clear domain boundaries
- 🔐 **Secure Authentication** - JWT-based auth with bcrypt password hashing
- 📁 **File Upload Support** - Handle team logos and player profile photos
- 📊 **Rich Statistics** - Dashboard metrics and detailed match reports
- ✅ **Input Validation** - Comprehensive request validation with detailed error messages
- 🚀 **High Performance** - Built on Gin framework for optimal speed
- 📖 **Well Documented** - Complete API documentation with examples

---

## ✨ Features

### 🔐 Authentication & Authorization
- User registration and login with JWT tokens
- Password encryption using bcrypt
- Profile management with photo upload
- Token-based API protection

### ⚽ Team Management
- Complete CRUD operations for teams
- Team logo upload and management (JPG, PNG, GIF, WebP)
- Team statistics and performance metrics
- Soft delete support

### 👥 Player Management
- Player CRUD with team association
- Player profile photo uploads
- Position-based categorization (Forward, Midfielder, Defender, Goalkeeper)
- Jersey number validation (unique per team)
- Physical attributes (height, weight)

### 📅 Match & Schedule Management
- Match scheduling with date/time/venue
- Live score updates
- Match result management
- Filter matches by date (today, upcoming, past)

### ⚽ Goal & Scorer Tracking
- Record goals with player attribution
- Goal timing and match association
- Scorer statistics

### 📊 Reports & Statistics
- Dashboard overview (total teams, players, matches, goals)
- Team-specific match reports
- Today's match listings
- Comprehensive match details with teams and goals

---

## 🚀 Tech Stack

| Technology | Version | Purpose |
|------------|---------|---------|
| **Go** | 1.24+ | Programming Language |
| **Gin** | 1.10.0 | HTTP Web Framework |
| **GORM** | 1.31.0 | ORM for Database Operations |
| **MySQL** | 8.0+ | Relational Database |
| **JWT** | 5.2.1 | Authentication & Authorization |
| **Bcrypt** | - | Password Hashing |
| **UUID** | 1.6.0 | Unique File Naming |
| **Validator** | 10.22.1 | Request Validation |
| **CORS** | 1.7.2 | Cross-Origin Resource Sharing |

---

## 🏛️ Architecture

This project follows **Clean Architecture** principles with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────┐
│                     HTTP Layer (Gin)                     │
│                  Routes & Middleware                     │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                   Handler Layer                          │
│        (HTTP Request/Response, Validation)               │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                   Service Layer                          │
│            (Business Logic, Use Cases)                   │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                  Repository Layer                        │
│             (Database Operations, GORM)                  │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                   Domain Layer                           │
│          (Entities, Business Rules, DTOs)                │
└─────────────────────────────────────────────────────────┘
```

**Benefits:**
- ✅ Testability - Each layer can be tested independently
- ✅ Maintainability - Changes in one layer don't affect others
- ✅ Scalability - Easy to extend with new features
- ✅ Flexibility - Swap implementations without affecting business logic

---

## 📦 Prerequisites

Before installing, ensure you have the following installed on your system:

### Required

- **Go** (version 1.24 or higher)
  ```bash
  # Check Go version
  go version
  ```
  Download: [https://golang.org/dl/](https://golang.org/dl/)

- **MySQL** (version 8.0 or higher)
  ```bash
  # Check MySQL version
  mysql --version
  ```
  Download: [https://dev.mysql.com/downloads/mysql/](https://dev.mysql.com/downloads/mysql/)

- **Git**
  ```bash
  # Check Git version
  git --version
  ```
  Download: [https://git-scm.com/downloads](https://git-scm.com/downloads)

### Optional but Recommended

- **Postman** or **Insomnia** - For API testing
- **MySQL Workbench** - For database management
- **VS Code** with Go extension - For development

---

## 🔧 Installation

### 1. Clone the Repository

```bash
# Clone via HTTPS
git clone https://github.com/alwijein/be_football.git

# OR clone via SSH
git clone git@github.com:alwijein/be_football.git

# Navigate to project directory
cd be_football
```

### 2. Install Dependencies

```bash
# Download and install all Go dependencies
go mod download

# Verify dependencies
go mod verify

# Tidy up go.mod and go.sum (optional)
go mod tidy
```

### 3. Create Database

```bash
# Login to MySQL
mysql -u root -p

# Create database
CREATE DATABASE football_db;

# Exit MySQL
exit;
```

**Alternative using MySQL Workbench:**
1. Open MySQL Workbench
2. Connect to your MySQL server
3. Execute: `CREATE DATABASE football_db;`

### 4. Configure Environment Variables

```bash
# Copy example environment file
cp .env.example .env

# Edit .env file with your configuration
# (Use nano, vim, or any text editor)
nano .env
```

**Configuration Options:**

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_NAME=football_db

# Server Configuration
SERVER_PORT=8080
GIN_MODE=debug

# JWT Configuration
JWT_SECRET=your-very-secret-key-please-change-in-production-12345
JWT_EXPIRATION_HOURS=24

# Application
APP_NAME=Football Management API
APP_VERSION=1.0.0
```

⚠️ **Important:** Change `JWT_SECRET` in production environment!

### 5. Create Upload Directories

```bash
# Create directories for file uploads
mkdir -p uploads/teams
mkdir -p uploads/players
mkdir -p uploads/profiles
```

**Windows (PowerShell):**
```powershell
New-Item -ItemType Directory -Force -Path uploads/teams
New-Item -ItemType Directory -Force -Path uploads/players
New-Item -ItemType Directory -Force -Path uploads/profiles
```

---

## 🔧 Configuration

### Database Configuration

Edit `.env` file to match your MySQL setup:

```env
DB_HOST=localhost          # MySQL host (use 'localhost' for local)
DB_PORT=3306              # MySQL port (default: 3306)
DB_USER=root              # Your MySQL username
DB_PASSWORD=              # Your MySQL password (leave empty if no password)
DB_NAME=football_db       # Database name
```

### Server Configuration

```env
SERVER_PORT=8080          # Port where API will run
GIN_MODE=debug           # Options: debug, release, test
```

- **debug**: Development mode with detailed logs
- **release**: Production mode with optimized performance
- **test**: Testing mode

### JWT Configuration

```env
JWT_SECRET=your-secret-key-min-32-characters-long-for-security
JWT_EXPIRATION_HOURS=24
```

⚠️ **Security Best Practices:**
- Use a strong, random secret key (minimum 32 characters)
- Never commit `.env` to version control
- Use different secrets for different environments
- Rotate secrets periodically

---

## � Running the Application

### Development Mode

```bash
# Build the application
go build -o bin/main.exe cmd/api/main.go

# Run the application
./bin/main.exe
```

**Windows (PowerShell):**
```powershell
# Build
go build -o .\bin\main.exe .\cmd\api\main.go

# Run
.\bin\main.exe
```

### Using Go Run (Quick Development)

```bash
# Run directly without building
go run cmd/api/main.go
```

### Production Mode

```bash
# Set production mode in .env
GIN_MODE=release

# Build optimized binary
go build -ldflags="-s -w" -o bin/main cmd/api/main.go

# Run
./bin/main
```

### Verify Server is Running

```bash
# Test health endpoint
curl http://localhost:8080/api/health

# Expected response:
# {"meta":{"code":200,"status":"success","message":"OK"},"data":{"status":"healthy"}}
```

### Server Information

Once running, the server will be available at:

- **Local**: `http://localhost:8080`
- **Network**: `http://YOUR_LOCAL_IP:8080` (e.g., `http://192.168.1.100:8080`)

**API Base URL**: `http://localhost:8080/api`

---

## 📖 API Documentation

### Base URL

```
http://localhost:8080/api
```

### Authentication

Most endpoints require JWT authentication. Include the token in the Authorization header:

```
Authorization: Bearer YOUR_JWT_TOKEN
```

### Response Format

All API responses follow this standard format:

**Success Response:**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Operation successful"
  },
  "data": {
    // Response data here
  }
}
```

**Error Response:**
```json
{
  "meta": {
    "code": 400,
    "status": "error",
    "message": "Error description"
  },
  "data": null
}
```

### Quick Reference

| Category | Endpoint | Method | Auth Required |
|----------|----------|--------|---------------|
| **Auth** | `/login` | POST | ❌ |
| | `/register` | POST | ❌ |
| | `/profile` | GET | ✅ |
| | `/profile` | PUT | ✅ |
| | `/change-password` | PUT | ✅ |
| **Dashboard** | `/dashboard` | GET | ✅ |
| | `/matches/today` | GET | ✅ |
| **Teams** | `/teams` | GET, POST | ✅ |
| | `/teams/:id` | GET, PUT, DELETE | ✅ |
| | `/teams/:id/players` | GET, POST | ✅ |
| | `/teams/:id/match-reports` | GET | ✅ |
| **Players** | `/players/:id` | GET, PUT, DELETE | ✅ |
| **Schedules** | `/schedules` | GET, POST | ✅ |
| | `/schedules/:id` | GET, PUT, DELETE | ✅ |
| | `/schedules/:id/result` | PUT | ✅ |
| **Matches** | `/matches` | GET | ✅ |
| | `/matches/:id` | GET | ✅ |
| **Goals** | `/goals` | POST | ✅ |
| | `/goals/:id` | DELETE | ✅ |

### Complete Documentation

For detailed API documentation with request/response examples, see:
- 📄 **[API_DOCUMENTATION.md](./API_DOCUMENTATION.md)** - Complete endpoint documentation
- 📄 **[API_QUICK_REFERENCE.md](./API_QUICK_REFERENCE.md)** - Quick reference guide

### Postman Collection

Import the Postman collection for easy API testing:

1. Open Postman
2. Click **Import**
3. Use the collection URL or import from file
4. Set environment variables:
   - `base_url`: `http://localhost:8080/api`
   - `token`: Your JWT token after login

---

## 🗄️ Database Schema

### Entity Relationship Diagram

```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│    Users    │         │    Teams     │         │   Players   │
├─────────────┤         ├──────────────┤         ├─────────────┤
│ id (PK)     │         │ id (PK)      │◄────────│ id (PK)     │
│ username    │         │ name         │  1:N    │ team_id (FK)│
│ email       │         │ city         │         │ name        │
│ password    │         │ logo         │         │ position    │
│ full_name   │         │ est_year     │         │ height      │
│ photo_url   │         │ created_at   │         │ weight      │
│ created_at  │         │ updated_at   │         │ jersey_no   │
└─────────────┘         │ deleted_at   │         │ created_at  │
                        └──────────────┘         │ updated_at  │
                              │                   │ deleted_at  │
                              │                   └─────────────┘
                              │                         │
                              │                         │
                        ┌─────▼─────┐             ┌────▼────┐
                        │  Matches  │             │  Goals  │
                        ├───────────┤             ├─────────┤
                        │ id (PK)   │◄────────────│ id (PK) │
                        │ home_id   │      N:1    │ match_id│
                        │ away_id   │             │ player  │
                        │ match_date│             │ minute  │
                        │ venue     │             │ created │
                        │ home_score│             └─────────┘
                        │ away_score│
                        │ status    │
                        │ created_at│
                        │ updated_at│
                        │ deleted_at│
                        └───────────┘
```

### Database Tables

**Users Table:**
```sql
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    photo_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
```

**Teams Table:**
```sql
CREATE TABLE teams (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    logo VARCHAR(255),
    established_year INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
```

**Players Table:**
```sql
CREATE TABLE players (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    team_id BIGINT,
    name VARCHAR(100) NOT NULL,
    height INT NOT NULL,
    weight INT NOT NULL,
    position VARCHAR(20) NOT NULL,
    jersey_number INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE SET NULL,
    UNIQUE KEY unique_jersey_per_team (team_id, jersey_number)
);
```

**Matches Table:**
```sql
CREATE TABLE matches (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    home_team_id BIGINT NOT NULL,
    away_team_id BIGINT NOT NULL,
    match_date DATETIME NOT NULL,
    venue VARCHAR(255) NOT NULL,
    home_score INT DEFAULT 0,
    away_score INT DEFAULT 0,
    status VARCHAR(20) DEFAULT 'scheduled',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (home_team_id) REFERENCES teams(id) ON DELETE CASCADE,
    FOREIGN KEY (away_team_id) REFERENCES teams(id) ON DELETE CASCADE
);
```

**Goals Table:**
```sql
CREATE TABLE goals (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    match_id BIGINT NOT NULL,
    player_id BIGINT NOT NULL,
    minute INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (match_id) REFERENCES matches(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE
);
```

**Auto Migration:**
The application automatically creates and migrates these tables on startup using GORM.

---

## � Project Structure

```
be_football/
│
├── cmd/
│   └── api/
│       └── main.go                    # Application entry point & server initialization
│
├── internal/
│   ├── config/
│   │   └── config.go                  # Environment configuration loader
│   │
│   ├── domain/                        # Domain entities (Business models)
│   │   ├── user.go                    # User entity
│   │   ├── team.go                    # Team entity
│   │   ├── player.go                  # Player entity with position enum
│   │   ├── match.go                   # Match/Schedule entity
│   │   └── goal.go                    # Goal entity
│   │
│   ├── dto/                           # Data Transfer Objects
│   │   ├── auth.go                    # Login, Register, Profile DTOs
│   │   ├── team.go                    # Team request/response DTOs
│   │   ├── player.go                  # Player request/response DTOs
│   │   ├── match.go                   # Match/Schedule DTOs
│   │   ├── goal.go                    # Goal DTOs
│   │   └── response.go                # Standard response wrappers
│   │
│   ├── repository/                    # Data Access Layer
│   │   ├── user_repository.go         # User CRUD operations
│   │   ├── team_repository.go         # Team CRUD operations
│   │   ├── player_repository.go       # Player CRUD operations
│   │   ├── match_repository.go        # Match CRUD operations
│   │   └── goal_repository.go         # Goal CRUD operations
│   │
│   ├── service/                       # Business Logic Layer
│   │   ├── auth_service.go            # Authentication logic
│   │   ├── team_service.go            # Team business logic
│   │   ├── player_service.go          # Player business logic
│   │   ├── match_service.go           # Match business logic
│   │   ├── goal_service.go            # Goal business logic
│   │   └── report_service.go          # Dashboard & reports logic
│   │
│   ├── handler/                       # HTTP Handlers (Controllers)
│   │   ├── auth_handler.go            # Auth endpoints
│   │   ├── team_handler.go            # Team endpoints
│   │   ├── player_handler.go          # Player endpoints
│   │   ├── match_handler.go           # Match/Schedule endpoints
│   │   ├── goal_handler.go            # Goal endpoints
│   │   └── report_handler.go          # Dashboard & report endpoints
│   │
│   ├── middleware/                    # HTTP Middleware
│   │   ├── auth_middleware.go         # JWT authentication
│   │   ├── cors_middleware.go         # CORS configuration
│   │   └── logger_middleware.go       # Request logging
│   │
│   ├── routes/
│   │   └── routes.go                  # Route definitions & grouping
│   │
│   └── utils/                         # Utility functions
│       ├── jwt.go                     # JWT token generation & validation
│       ├── password.go                # Bcrypt password hashing
│       ├── response.go                # Standard response helpers
│       ├── file_upload.go             # File upload utilities
│       └── validation.go              # Custom validators
│
├── pkg/
│   └── database/
│       └── mysql.go                   # Database connection & auto-migration
│
├── uploads/                           # File uploads directory
│   ├── teams/                         # Team logos
│   ├── players/                       # Player photos
│   └── profiles/                      # User profile photos
│
├── bin/                               # Compiled binaries
│
├── .env                               # Environment variables (DO NOT COMMIT)
├── .env.example                       # Environment variables template
├── .gitignore                         # Git ignore rules
├── go.mod                             # Go module dependencies
├── go.sum                             # Dependency checksums
├── README.md                          # This file
├── API_DOCUMENTATION.md               # Complete API documentation
├── API_QUICK_REFERENCE.md             # Quick API reference
└── LICENSE                            # Project license
```

### Key Directories Explained

- **`cmd/api/`**: Application entry point
- **`internal/`**: Private application code (not importable by other projects)
- **`internal/domain/`**: Core business entities and models
- **`internal/dto/`**: Data structures for API requests/responses
- **`internal/repository/`**: Database operations (GORM)
- **`internal/service/`**: Business logic and use cases
- **`internal/handler/`**: HTTP request handlers
- **`internal/middleware/`**: HTTP middleware (auth, CORS, logging)
- **`internal/routes/`**: API route definitions
- **`internal/utils/`**: Reusable utility functions
- **`pkg/`**: Public libraries (can be imported by other projects)
- **`uploads/`**: File storage for user uploads

---

## 🧪 Testing

### Manual Testing with cURL

**1. Login:**
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

**2. Get Teams (with authentication):**
```bash
curl -X GET http://localhost:8080/api/teams \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**3. Create Team with Logo:**
```bash
curl -X POST http://localhost:8080/api/teams \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "name=Arsenal FC" \
  -F "city=London" \
  -F "established_year=1886" \
  -F "logo=@/path/to/logo.png"
```

### Automated Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests in specific package
go test ./internal/service/...

# Verbose output
go test -v ./...
```

### Test Data

For sample data and test scenarios, see:
- 📄 **SAMPLE_DATA.md** - Sample data for testing

---

## 🚀 Deployment

### Docker Deployment (Recommended)

**1. Create Dockerfile:**
```dockerfile
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

RUN mkdir -p uploads/teams uploads/players uploads/profiles

EXPOSE 8080
CMD ["./main"]
```

**2. Build and run:**
```bash
# Build image
docker build -t football-api .

# Run container
docker run -d -p 8080:8080 --name football-api football-api
```

### Docker Compose (with MySQL)

**docker-compose.yml:**
```yaml
version: '3.8'

services:
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: football_db
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  api:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - mysql
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USER: root
      DB_PASSWORD: rootpassword
      DB_NAME: football_db

volumes:
  mysql_data:
```

**Run:**
```bash
docker-compose up -d
```

### Production Deployment

**Security Checklist:**
- ✅ Change `JWT_SECRET` to a strong random key
- ✅ Set `GIN_MODE=release`
- ✅ Use environment variables for sensitive data
- ✅ Enable HTTPS/TLS
- ✅ Set up database backups
- ✅ Configure proper CORS policies
- ✅ Implement rate limiting
- ✅ Set up monitoring and logging
- ✅ Use reverse proxy (Nginx/Apache)

---

## 🤝 Contributing

Contributions are welcome! Please follow these guidelines:

### How to Contribute

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Commit your changes**
   ```bash
   git commit -m 'Add some amazing feature'
   ```
4. **Push to the branch**
   ```bash
   git push origin feature/amazing-feature
   ```
5. **Open a Pull Request**

### Coding Standards

- Follow Go best practices and conventions
- Use `gofmt` to format code
- Add comments for exported functions
- Write unit tests for new features
- Update documentation when needed

### Commit Message Format

```
type(scope): subject

body

footer
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting)
- `refactor`: Code refactoring
- `test`: Adding tests
- `chore`: Maintenance tasks

**Example:**
```
feat(auth): add refresh token endpoint

Implement refresh token functionality to allow
users to get new access tokens without re-login.

Closes #123
```

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```
MIT License

Copyright (c) 2024 Football Management API

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
```

---

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- [Go Playground Validator](https://github.com/go-playground/validator)
- [godotenv](https://github.com/joho/godotenv)

---

## 📞 Support & Contact

### Found a Bug?

Please [open an issue](https://github.com/alwijein/be_football/issues) with:
- Description of the bug
- Steps to reproduce
- Expected vs actual behavior
- Screenshots (if applicable)

### Need Help?

- 📖 Check the [API Documentation](./API_DOCUMENTATION.md)
- 💬 Open a [discussion](https://github.com/alwijein/be_football/discussions)
- 📧 Contact: [your-email@example.com](mailto:your-email@example.com)

### Feature Requests

We welcome feature requests! Please [open an issue](https://github.com/alwijein/be_football/issues) with the `enhancement` label.

---

## 🗺️ Roadmap

### Version 1.0.0 (Current) ✅
- Core API functionality
- Authentication & Authorization
- Team, Player, Match management
- File upload support
- Dashboard & Reports

### Version 1.1.0 (Planned) 🚧
- [ ] Refresh token implementation
- [ ] Email verification
- [ ] Password reset functionality
- [ ] Enhanced search and filtering
- [ ] Export reports to PDF/Excel

### Version 2.0.0 (Future) 🔮
- [ ] Real-time match updates (WebSocket)
- [ ] Admin dashboard web interface
- [ ] Multi-language support
- [ ] Advanced analytics and statistics
- [ ] Mobile app integration
- [ ] Social media integration

---

## 📊 Project Status

![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)
![Build](https://img.shields.io/badge/Build-Passing-success?style=for-the-badge)
![Coverage](https://img.shields.io/badge/Coverage-85%25-green?style=for-the-badge)
![Version](https://img.shields.io/badge/Version-1.0.0-blue?style=for-the-badge)

**Current Version:** 1.0.0  
**Last Updated:** October 2024  
**Status:** Production Ready ✅

---

<div align="center">

### ⚽ Built with ❤️ using Go

**Star this repo if you find it useful!** ⭐

[Report Bug](https://github.com/alwijein/be_football/issues) • [Request Feature](https://github.com/alwijein/be_football/issues) • [Documentation](./API_DOCUMENTATION.md)

</div>

## 📁 Project Structure- 📅 **Jadwal Pertandingan**: Penjadwalan pertandingan antar tim

### Prerequisites

- 🎯 **Pencatatan Hasil**: Input hasil pertandingan dan goal dengan detail pemain pencetak

- Go 1.24 atau lebih tinggi

- MySQL 8.0+```- 📊 **Report & Statistik**: Laporan lengkap pertandingan dengan top scorer dan akumulasi kemenangan

- Postman (untuk testing)

be_football/- 🗑️ **Soft Delete**: Semua operasi hapus menggunakan soft delete untuk audit trail

### Installation Steps

├── cmd/- ✅ **Validation**: Request validation menggunakan validator

1. **Clone repository**

```bash│   └── api/- 🛡️ **Security**: CORS, JWT middleware, error handling

git clone <repo-url>

cd be_football│       └── main.go              # Entry point aplikasi- 📄 **Pagination**: Support pagination untuk semua list endpoints

```

├── internal/

2. **Install dependencies**

```bash│   ├── config/                  # Konfigurasi aplikasi## 🏗️ Arsitektur

go mod download

```│   ├── domain/                  # Entities/Models



3. **Setup database**│   ├── dto/                     # Data Transfer ObjectsAplikasi ini dibangun menggunakan **Clean Architecture** dengan pemisahan layer:

```sql

CREATE DATABASE football_db;│   ├── handler/                 # HTTP Handlers

```

│   ├── middleware/              # Middleware (Auth, CORS, Logger)- **Handler Layer**: HTTP request/response handling

4. **Configure environment** (Opsional)

│   ├── repository/              # Data Access Layer- **Service Layer**: Business logic dan use cases

Edit `internal/config/config.go` atau buat file `.env`:

```env│   ├── routes/                  # Route definitions- **Repository Layer**: Data access dan database operations

DB_HOST=localhost

DB_PORT=3306│   ├── service/                 # Business Logic Layer- **Domain Layer**: Business entities dan models

DB_USER=root

DB_PASSWORD=│   └── utils/                   # Utilities (JWT, Password, Response)

DB_NAME=football_db

SERVER_PORT=8080└── pkg/## 🚀 Tech Stack

```

    └── database/                # Database connection & migration

5. **Run migrations**

```- **Language**: Go 1.21+

Migrations akan otomatis berjalan saat aplikasi start.

- **Framework**: Gin (Go Web Framework)

6. **Build & Run**

```bash## ⚙️ Setup- **Database**: PostgreSQL 12+

# Build

go build -o bin/football-api.exe cmd/api/main.go- **ORM**: GORM



# Run### 1. Clone Repository- **Authentication**: JWT (JSON Web Token)

go run cmd/api/main.go

```- **Password Hashing**: bcrypt



Server akan berjalan di:```bash- **Validation**: go-playground/validator

- Local: `http://localhost:8080`

- Network: `http://192.168.43.76:8080` (sesuai IP Anda)git clone <repository-url>



## 📡 API Endpointscd be_football## 📁 Struktur Project



Base URL: `/api````



### Authentication (Public: Login only)```



| Method | Endpoint | Description |### 2. Install Dependenciesbe_football/

|--------|----------|-------------|

| POST | `/api/login` | Login user |├── cmd/

| POST | `/api/logout` | Logout user (Auth) |

```bash│   └── api/

### Dashboard (Protected)

go mod tidy│       └── main.go                    # Entry point aplikasi

| Method | Endpoint | Description |

|--------|----------|-------------|```├── internal/

| GET | `/api/dashboard/stats` | Get dashboard statistics |

| GET | `/api/matches/today` | Get today's matches |│   ├── config/                        # Konfigurasi aplikasi



### Teams (Protected)### 3. Setup Database MySQL│   ├── domain/                        # Business domain entities



| Method | Endpoint | Description |│   │   ├── user.go                    # User entity

|--------|----------|-------------|

| GET | `/api/teams` | Get all teams (paginated) |Buat database MySQL:│   │   ├── team.go                    # Team entity

| GET | `/api/teams/registered?limit=5` | Get registered teams for homepage |

| GET | `/api/teams/:id` | Get team detail with players |│   │   ├── player.go                  # Player entity

| POST | `/api/teams` | Create new team |

| PUT | `/api/teams/:id` | Update team |```sql│   │   ├── match.go                   # Match entity

| DELETE | `/api/teams/:id` | Delete team |

CREATE DATABASE football_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;│   │   └── goal.go                    # Goal entity

### Players (Protected)

```│   ├── repository/                    # Data access layer (Repository Pattern)

| Method | Endpoint | Description |

|--------|----------|-------------|│   ├── service/                       # Business logic layer (Service Pattern)

| GET | `/api/teams/:team_id/players` | Get team players |

| POST | `/api/teams/:team_id/players` | Create player |### 4. Konfigurasi Environment│   ├── handler/                       # HTTP handlers (Controllers)

| PUT | `/api/players/:id` | Update player |

| DELETE | `/api/players/:id` | Delete player |│   ├── middleware/                    # HTTP middleware (Auth, CORS, Logger, etc)



### Schedules (Protected)Edit file `.env`:│   ├── routes/                        # Route definitions



| Method | Endpoint | Description |│   ├── dto/                           # Data transfer objects

|--------|----------|-------------|

| GET | `/api/schedules` | Get all schedules (paginated) |```properties│   └── utils/                         # Utility functions (JWT, Password, Response)

| GET | `/api/schedules/:id` | Get schedule detail |

| POST | `/api/schedules` | Create schedule |# Database Configuration├── pkg/

| PUT | `/api/schedules/:id/result` | Update match result |

| DELETE | `/api/schedules/:id` | Delete schedule |DB_HOST=localhost│   └── database/                      # Database connection & migration



### Scorers (Protected)DB_PORT=3306├── .env                               # Environment variables



| Method | Endpoint | Description |DB_USER=root├── .env.example                       # Environment variables example

|--------|----------|-------------|

| GET | `/api/schedules/:schedule_id/scorers` | Get match scorers |DB_PASSWORD=your_password├── .gitignore

| POST | `/api/schedules/:schedule_id/scorers` | Add goal/scorer |

| DELETE | `/api/scorers/:id` | Delete scorer |DB_NAME=football_db├── go.mod



### Reports (Protected)├── go.sum



| Method | Endpoint | Description |# Server Configuration├── README.md                          # This file

|--------|----------|-------------|

| GET | `/api/teams/:team_id/match-reports` | Get team match history |SERVER_PORT=8080├── API_DOCUMENTATION.md               # Detailed API documentation



### Profile (Protected)GIN_MODE=debug├── SETUP_GUIDE.md                     # Setup and installation guide



| Method | Endpoint | Description |├── SAMPLE_DATA.md                     # Sample data for testing

|--------|----------|-------------|

| GET | `/api/profile` | Get current user profile |# JWT Configuration├── ARCHITECTURE.md                    # Architecture and design patterns

| PUT | `/api/profile` | Update profile |

| PUT | `/api/profile/password` | Change password |JWT_SECRET=your-secret-key-change-this-in-production└── api_collection.json                # Postman/Thunder Client collection



**Total: 26 Endpoints**JWT_EXPIRATION_HOURS=24```



## 📝 Response Format



Semua endpoint menggunakan format response standar:# Application## 🚦 Quick Start



### Success ResponseAPP_NAME=Football Management API

```json

{APP_VERSION=1.0.0### Prerequisites

  "meta": {

    "code": 200,```

    "status": "success",

    "message": "Operation successful"- Go 1.21 atau lebih tinggi

  },

  "data": {### 5. Run Application- PostgreSQL 12 atau lebih tinggi

    // response data

  }

}

``````bash### Installation



### Success with Paginationgo run cmd/api/main.go

```json

{```1. **Clone repository** (jika dari git)

  "meta": {

    "code": 200,```bash

    "status": "success",

    "message": null,Server akan berjalan di: **http://localhost:8080**git clone <repository-url>

    "pagination": {

      "current_page": 1,cd be_football

      "total_pages": 5,

      "total_items": 48,## 📡 API Response Format```

      "per_page": 10

    }

  },

  "data": [...]Semua response API menggunakan format standar:2. **Install dependencies**

}

``````bash



### Error Response### Success Responsego mod download

```json

{# atau

  "meta": {

    "code": 400,```jsongo mod tidy

    "status": "error",

    "message": "Error description"{```

  },

  "data": null  "meta": {

}

```    "code": 200,3. **Setup database PostgreSQL**



## 🔐 Authentication    "status": "success",```sql



Gunakan JWT Bearer Token untuk endpoint yang dilindungi.    "message": "Data retrieved successfully"CREATE DATABASE football_db;



### Login  },```

```bash

POST /api/login  "data": {

Content-Type: application/json

    // response data here4. **Setup environment variables**

{

  "username": "admin",  }```bash

  "password": "admin123"

}}# Windows

```

```copy .env.example .env

Response:

```json

{

  "meta": {### Error Response# Linux/Mac

    "code": 200,

    "status": "success",cp .env.example .env

    "message": "Login successful"

  },```json```

  "data": {

    "token": "eyJhbGc...",{

    "user": {

      "id": 1,  "meta": {Edit file `.env` sesuai konfigurasi Anda:

      "username": "admin",

      "email": "admin@example.com",    "code": 400,```env

      "full_name": "Administrator"

    }    "status": "error",DB_HOST=localhost

  }

}    "message": "Error message here"DB_PORT=5432

```

  },DB_USER=postgres

### Use Token

```bash  "data": {DB_PASSWORD=your_password

GET /api/teams

Authorization: Bearer eyJhbGc...    // error details hereDB_NAME=football_db

```

  }

## 📦 Postman Collection

}JWT_SECRET=your-secret-key-change-this

Import file: `football_api_v2.postman_collection.json`

``````

Collection sudah termasuk:

- ✅ 26 Endpoints lengkap

- ✅ Auto-save token setelah login

- ✅ Bearer auth otomatis## 🔐 Authentication5. **Jalankan aplikasi**

- ✅ Example request bodies

```bash

## 🏗️ Project Structure

API menggunakan JWT Bearer Token untuk authentication.# Option 1: Direct run

```

be_football/go run cmd/api/main.go

├── cmd/

│   └── api/### Register

│       └── main.go                 # Application entry point

├── internal/# Option 2: Build then run

│   ├── config/

│   │   └── config.go               # Configuration```bashgo build -o bin/api.exe cmd/api/main.go

│   ├── domain/

│   │   ├── user.go                 # User entityPOST /api/v1/auth/register.\bin\api.exe

│   │   ├── team.go                 # Team entity

│   │   ├── player.go               # Player entityContent-Type: application/json```

│   │   ├── match.go                # Match entity

│   │   └── goal.go                 # Goal entity

│   ├── dto/

│   │   ├── auth.go                 # Auth DTOs{6. **Verify aplikasi berjalan**

│   │   ├── team.go                 # Team DTOs

│   │   ├── player.go               # Player DTOs  "username": "admin",```bash

│   │   ├── match.go                # Match DTOs

│   │   ├── report.go               # Report DTOs  "email": "admin@example.com",curl http://localhost:8080/health

│   │   └── pagination.go           # Pagination DTO

│   ├── repository/  "password": "password123",```

│   │   ├── user_repository.go      # User data access

│   │   ├── team_repository.go      # Team data access  "full_name": "Admin User"

│   │   ├── player_repository.go    # Player data access

│   │   ├── match_repository.go     # Match data access}API akan berjalan di `http://localhost:8080`

│   │   └── goal_repository.go      # Goal data access

│   ├── service/```

│   │   ├── auth_service.go         # Auth business logic

│   │   ├── team_service.go         # Team business logic📖 **Untuk panduan lengkap, lihat [SETUP_GUIDE.md](SETUP_GUIDE.md)**

│   │   ├── player_service.go       # Player business logic

│   │   ├── match_service.go        # Match business logic### Login

│   │   └── report_service.go       # Report business logic

│   ├── handler/## API Endpoints

│   │   ├── auth_handler.go         # Auth HTTP handlers

│   │   ├── team_handler.go         # Team HTTP handlers```bash

│   │   ├── player_handler.go       # Player HTTP handlers

│   │   ├── match_handler.go        # Match HTTP handlersPOST /api/v1/auth/login### Authentication

│   │   └── report_handler.go       # Report HTTP handlers

│   ├── middleware/Content-Type: application/json- `POST /api/v1/auth/register` - Register admin baru

│   │   ├── auth.go                 # JWT middleware

│   │   ├── cors.go                 # CORS middleware- `POST /api/v1/auth/login` - Login dan dapatkan JWT token

│   │   └── logger.go               # Logger middleware

│   ├── routes/{

│   │   └── routes.go               # Route definitions

│   └── utils/  "email": "admin@example.com",### Teams

│       ├── jwt.go                  # JWT utilities

│       ├── password.go             # Password hashing  "password": "password123"- `GET /api/v1/teams` - Get semua tim (dengan pagination)

│       └── response.go             # Response helpers

├── pkg/}- `GET /api/v1/teams/:id` - Get detail tim

│   └── database/

│       └── database.go             # Database connection```- `POST /api/v1/teams` - Buat tim baru (auth required)

├── bin/                            # Compiled binaries

├── go.mod- `PUT /api/v1/teams/:id` - Update tim (auth required)

├── go.sum

├── README.mdResponse akan mengembalikan JWT token yang harus digunakan untuk request yang memerlukan authentication.- `DELETE /api/v1/teams/:id` - Hapus tim (auth required)

├── REFACTORING_SUMMARY.md

└── football_api_v2.postman_collection.json

```

## 📋 Main Endpoints### Players

## 🎯 Usage Examples

- `GET /api/v1/players` - Get semua pemain (dengan pagination)

### 1. Login & Get Token

```bash### Teams- `GET /api/v1/players/:id` - Get detail pemain

curl -X POST http://localhost:8080/api/login \

  -H "Content-Type: application/json" \- `GET /api/v1/teams` - Get all teams- `GET /api/v1/teams/:id/players` - Get pemain berdasarkan tim

  -d '{"username":"admin","password":"admin123"}'

```- `GET /api/v1/teams/:id` - Get team by ID- `POST /api/v1/players` - Tambah pemain baru (auth required)



### 2. Get Dashboard Stats- `POST /api/v1/teams` - Create team (Auth required)- `PUT /api/v1/players/:id` - Update pemain (auth required)

```bash

curl -X GET http://localhost:8080/api/dashboard/stats \- `PUT /api/v1/teams/:id` - Update team (Auth required)- `DELETE /api/v1/players/:id` - Hapus pemain (auth required)

  -H "Authorization: Bearer YOUR_TOKEN"

```- `DELETE /api/v1/teams/:id` - Delete team (Auth required)



### 3. Get Registered Teams (Homepage)### Matches

```bash

curl -X GET "http://localhost:8080/api/teams/registered?limit=5" \### Players- `GET /api/v1/matches` - Get semua jadwal pertandingan

  -H "Authorization: Bearer YOUR_TOKEN"

```- `GET /api/v1/players` - Get all players- `GET /api/v1/matches/:id` - Get detail pertandingan



### 4. Get Today's Matches- `GET /api/v1/players/:id` - Get player by ID- `POST /api/v1/matches` - Buat jadwal pertandingan (auth required)

```bash

curl -X GET http://localhost:8080/api/matches/today \- `POST /api/v1/players` - Create player (Auth required)- `PUT /api/v1/matches/:id` - Update jadwal (auth required)

  -H "Authorization: Bearer YOUR_TOKEN"

```- `PUT /api/v1/players/:id` - Update player (Auth required)- `DELETE /api/v1/matches/:id` - Hapus jadwal (auth required)



### 5. Create Team- `DELETE /api/v1/players/:id` - Delete player (Auth required)

```bash

curl -X POST http://localhost:8080/api/teams \### Match Results

  -H "Authorization: Bearer YOUR_TOKEN" \

  -H "Content-Type: application/json" \### Matches- `POST /api/v1/matches/:id/result` - Input hasil pertandingan (auth required)

  -d '{

    "name": "Manchester United",- `GET /api/v1/matches` - Get all matches- `PUT /api/v1/matches/:id/result` - Update hasil (auth required)

    "logo": "https://example.com/mu-logo.png",

    "established_year": 1878,- `GET /api/v1/matches/:id` - Get match by ID

    "address": "Old Trafford",

    "city": "Manchester"- `POST /api/v1/matches` - Create match (Auth required)### Goals

  }'

```- `PUT /api/v1/matches/:id` - Update match (Auth required)- `POST /api/v1/matches/:id/goals` - Tambah goal (auth required)



### 6. Update Match Result- `DELETE /api/v1/matches/:id` - Delete match (Auth required)- `DELETE /api/v1/goals/:id` - Hapus goal (auth required)

```bash

curl -X PUT http://localhost:8080/api/schedules/1/result \- `POST /api/v1/matches/:id/result` - Set match result (Auth required)

  -H "Authorization: Bearer YOUR_TOKEN" \

  -H "Content-Type: application/json" \- `POST /api/v1/matches/:id/goals` - Add goal to match (Auth required)### Reports

  -d '{

    "home_score": 2,- `GET /api/v1/reports/matches` - Get report semua pertandingan

    "away_score": 1,

    "status": "Completed"### Reports- `GET /api/v1/reports/matches/:id` - Get report detail pertandingan

  }'

```- `GET /api/v1/reports/matches` - Get all match reports



## 🔄 Migration Notes- `GET /api/v1/reports/matches/:id` - Get match report by ID## Response Format



### Changes from v1 to v2:



1. **Base path**: `/api/v1/` → `/api/`## 🧪 Testing### Success Response

2. **Removed**: `/api/auth/register` endpoint

3. **New endpoints**:```json

   - `POST /api/logout`

   - `GET /api/dashboard/stats`Import Postman collection: `football_api.postman_collection.json`{

   - `GET /api/matches/today`

   - `GET /api/teams/registered`  "success": true,

   - `GET /api/teams/:team_id/match-reports`

   - `GET /api/profile`## 📝 Features  "message": "Success message",

   - `PUT /api/profile`

   - `PUT /api/profile/password`  "data": {}

   - `GET /api/schedules/:schedule_id/scorers`

✅ Clean Architecture (Handler → Service → Repository → Domain)  }

4. **Renamed**: "matches" → "schedules" (semantik lebih jelas)

5. **Players nested**: Sekarang di bawah `/teams/:team_id/players`✅ JWT Authentication  ```



## 🐛 Troubleshooting✅ Soft Delete  



### Port sudah digunakan✅ Input Validation  ### Error Response

```bash

# Windows✅ Error Handling  ```json

netstat -ano | findstr :8080

taskkill /PID <PID> /F✅ CORS Support  {



# Linux/Mac✅ Request Logging    "success": false,

lsof -i :8080

kill -9 <PID>✅ Standardized API Response    "message": "Error message",

```

  "errors": []

### Database connection error

- Pastikan MySQL sudah running## 🔧 Development}

- Check credentials di `config.go`

- Pastikan database `football_db` sudah dibuat```



### JWT token expired### Build

- Login ulang untuk mendapatkan token baru

- Default expiry: 24 jam### Paginated Response



## 📱 Access from Mobile```bash```json



Server listen di `0.0.0.0:8080` sehingga bisa diakses dari device lain di jaringan yang sama.go build -o bin/api.exe cmd/api/main.go{



Dari mobile/device lain:```  "success": true,

```

http://192.168.43.76:8080/api/  "message": "Success",

```

### Run  "data": [],

*Note: Ganti IP sesuai dengan IP komputer Anda*

  "pagination": {

## 📄 License

```bash    "page": 1,

MIT License

./bin/api.exe    "limit": 10,

## 👨‍💻 Author

```    "total": 100,

Football API Development Team

    "total_pages": 10

---

## 📄 License  }

**Happy Coding! ⚽**

}

MIT License```


## 🔒 Security Features

- ✅ Password hashing menggunakan bcrypt
- ✅ JWT token-based authentication
- ✅ Authorization middleware untuk protected endpoints
- ✅ Input validation pada semua request
- ✅ Soft delete untuk audit trail
- ✅ CORS middleware
- ✅ Error handling yang aman (tidak expose internal details)
- ✅ SQL injection protection (GORM ORM)

## 📚 Dokumentasi

- **[API Documentation](API_DOCUMENTATION.md)**: Dokumentasi lengkap semua endpoints
- **[Setup Guide](SETUP_GUIDE.md)**: Panduan instalasi dan konfigurasi
- **[Sample Data](SAMPLE_DATA.md)**: Contoh data untuk testing
- **[Architecture](ARCHITECTURE.md)**: Penjelasan arsitektur dan design patterns

## 🎯 Business Rules

1. **Teams**: Nama tim harus unik
2. **Players**: 
   - Nomor punggung harus unik dalam satu tim
   - Satu pemain hanya bisa bergabung dengan satu tim
3. **Matches**: 
   - Tim home dan away harus berbeda
   - Tidak bisa update pertandingan yang sudah selesai
4. **Goals**: 
   - Pemain harus dari salah satu tim yang bertanding
   - Menit gol: 1-120
5. **Soft Delete**: Semua operasi hapus menggunakan soft delete

## 🧪 Testing

### Manual Testing dengan cURL

1. **Register Admin**
```bash
curl -X POST http://localhost:8080/api/v1/auth/register ^
  -H "Content-Type: application/json" ^
  -d "{\"username\":\"admin\",\"email\":\"admin@example.com\",\"password\":\"password123\",\"full_name\":\"Admin User\"}"
```

2. **Login**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login ^
  -H "Content-Type: application/json" ^
  -d "{\"username\":\"admin\",\"password\":\"password123\"}"
```

3. **Create Team** (dengan token)
```bash
curl -X POST http://localhost:8080/api/v1/teams ^
  -H "Content-Type: application/json" ^
  -H "Authorization: Bearer YOUR_TOKEN" ^
  -d "{\"name\":\"Persija Jakarta\",\"established_year\":1928,\"address\":\"Jakarta\",\"city\":\"Jakarta\"}"
```

### Testing dengan Postman/Thunder Client

Import file `api_collection.json` ke Postman atau Thunder Client untuk testing yang lebih mudah.

**Cara Import ke Postman:**
1. Buka Postman
2. Klik **"Import"** (atau Ctrl+O)
3. **Drag & drop** file `api_collection.json`, atau
4. Klik **"Upload Files"** → pilih `api_collection.json`
5. Klik **"Import"**
6. ✅ Collection "Football Management API" siap digunakan!

📖 **Panduan lengkap:** [POSTMAN_GUIDE.md](POSTMAN_GUIDE.md)

## 🤝 Contributing

Contributions are welcome! Silakan buat Pull Request atau buka Issue untuk saran dan bug report.

## 👨‍💻 Developer

Developed with ❤️ using Go and Clean Architecture principles

## 📄 License

MIT License
