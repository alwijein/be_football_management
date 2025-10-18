# 📦 Installation Guide - Football Management API

<div align="center">

**Step-by-Step Installation Instructions**

Complete guide from zero to running server

</div>

---

## 📋 Table of Contents

1. [Prerequisites](#1-prerequisites)
2. [Installation Steps](#2-installation-steps)
3. [Configuration](#3-configuration)
4. [Database Setup](#4-database-setup)
5. [Running the Application](#5-running-the-application)
6. [Verification](#6-verification)
7. [Troubleshooting](#7-troubleshooting)

---

## 1. Prerequisites

### Required Software

Before you begin, ensure you have the following installed:

#### ✅ Go (Golang)

**Version Required:** 1.24 or higher

**Check if installed:**
```bash
go version
# Expected output: go version go1.24.0 windows/amd64 (or similar)
```

**Install Go:**
- **Windows:** Download from [https://golang.org/dl/](https://golang.org/dl/)
  - Download the `.msi` installer
  - Run installer and follow wizard
  - Restart terminal/PowerShell

- **macOS:**
  ```bash
  # Using Homebrew
  brew install go
  ```

- **Linux (Ubuntu/Debian):**
  ```bash
  sudo apt update
  sudo apt install golang-go
  ```

**Verify Installation:**
```bash
go version
go env GOPATH
```

---

#### ✅ MySQL

**Version Required:** 8.0 or higher

**Check if installed:**
```bash
mysql --version
# Expected output: mysql Ver 8.0.30 (or higher)
```

**Install MySQL:**

- **Windows:**
  - Download from [https://dev.mysql.com/downloads/mysql/](https://dev.mysql.com/downloads/mysql/)
  - OR use [XAMPP](https://www.apachefriends.org/) / [Laragon](https://laragon.org/)
  - Run installer and set root password

- **macOS:**
  ```bash
  # Using Homebrew
  brew install mysql
  brew services start mysql
  ```

- **Linux (Ubuntu/Debian):**
  ```bash
  sudo apt update
  sudo apt install mysql-server
  sudo systemctl start mysql
  sudo mysql_secure_installation
  ```

**Verify MySQL is Running:**
```bash
# Windows (PowerShell)
Get-Service MySQL* | Select-Object Status, Name

# Linux/macOS
sudo systemctl status mysql
```

---

#### ✅ Git

**Check if installed:**
```bash
git --version
# Expected output: git version 2.40.0 (or similar)
```

**Install Git:**
- **Windows:** Download from [https://git-scm.com/download/win](https://git-scm.com/download/win)
- **macOS:** `brew install git`
- **Linux:** `sudo apt install git`

---

### Optional but Recommended

- **Postman** - For API testing ([Download](https://www.postman.com/downloads/))
- **MySQL Workbench** - For database management ([Download](https://dev.mysql.com/downloads/workbench/))
- **VS Code** - Code editor with Go extension ([Download](https://code.visualstudio.com/))

---

## 2. Installation Steps

### Step 1: Clone the Repository

```bash
# Using HTTPS
git clone https://github.com/alwijein/be_football.git

# OR using SSH (if you have SSH keys configured)
git clone git@github.com:alwijein/be_football.git

# Navigate to project directory
cd be_football
```

**Verify:**
```bash
# List files
ls
# You should see: cmd/, internal/, go.mod, README.md, etc.
```

---

### Step 2: Install Go Dependencies

```bash
# Download all dependencies
go mod download

# Verify dependencies
go mod verify

# Optional: Clean up dependencies
go mod tidy
```

**Expected Output:**
```
go: downloading github.com/gin-gonic/gin v1.10.0
go: downloading gorm.io/gorm v1.31.0
go: downloading gorm.io/driver/mysql v1.6.0
...
```

**Verify all dependencies are installed:**
```bash
go list -m all
```

---

### Step 3: Create Upload Directories

The application needs these directories for file uploads:

**Windows (PowerShell):**
```powershell
# Create all upload directories
New-Item -ItemType Directory -Force -Path uploads/teams
New-Item -ItemType Directory -Force -Path uploads/players
New-Item -ItemType Directory -Force -Path uploads/profiles
```

**Linux/macOS:**
```bash
mkdir -p uploads/teams
mkdir -p uploads/players
mkdir -p uploads/profiles
```

**Verify:**
```bash
# List directory structure
ls -R uploads/

# Expected output:
# uploads/teams
# uploads/players
# uploads/profiles
```

---

## 3. Configuration

### Step 1: Create Environment File

```bash
# Windows (PowerShell)
Copy-Item .env.example .env

# Linux/macOS
cp .env.example .env
```

### Step 2: Edit Configuration

Open `.env` file in your text editor:

**Windows:**
```powershell
notepad .env
# OR
code .env  # if using VS Code
```

**Linux/macOS:**
```bash
nano .env
# OR
vim .env
# OR
code .env  # if using VS Code
```

### Step 3: Configure Database

Update these lines in `.env`:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_mysql_password_here
DB_NAME=football_db
```

**Example configurations:**

**For XAMPP/Laragon (Windows):**
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=football_db
```

**For MySQL with password:**
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=mySecretPassword123
DB_NAME=football_db
```

**For remote MySQL:**
```env
DB_HOST=192.168.1.100
DB_PORT=3306
DB_USER=football_user
DB_PASSWORD=StrongPassword456
DB_NAME=football_db
```

### Step 4: Configure JWT Secret

⚠️ **IMPORTANT for Production:**

```env
# NEVER use this default value in production!
JWT_SECRET=your-very-secret-key-please-change-in-production-12345

# Generate a strong secret:
# Method 1: Random string (min 32 characters)
JWT_SECRET=a8f7d9e2b4c6d8e0f2a4b6c8d0e2f4a6b8c0d2e4f6a8b0c2d4e6f8a0b2c4d6e8

# Method 2: Using OpenSSL (recommended)
# Run: openssl rand -base64 32
JWT_SECRET=3xAmPl3B@s364St0ngS3cr3tK3yG3n3r@t3d!
```

### Step 5: Configure Server

```env
# Server Configuration
SERVER_PORT=8080          # Default port
GIN_MODE=debug           # Use 'release' for production
```

**GIN_MODE Options:**
- `debug` - Development (detailed logs, reload on change)
- `release` - Production (optimized, minimal logs)
- `test` - Testing mode

---

## 4. Database Setup

### Method 1: Using MySQL Command Line

```bash
# Step 1: Login to MySQL
mysql -u root -p
# Enter your password when prompted

# Step 2: Create database
CREATE DATABASE football_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# Step 3: Verify database created
SHOW DATABASES;

# Step 4: Exit
exit;
```

**Expected Output:**
```
Query OK, 1 row affected (0.01 sec)
```

---

### Method 2: Using MySQL Workbench

1. **Open MySQL Workbench**
2. **Connect to MySQL Server**
   - Click on your connection (usually "Local instance MySQL")
   - Enter password if prompted

3. **Create Database**
   - Click **"Create Schema" icon** (cylinder with plus)
   - Name: `football_db`
   - Charset: `utf8mb4`
   - Collation: `utf8mb4_unicode_ci`
   - Click **"Apply"**

4. **Verify**
   - You should see `football_db` in the left sidebar

---

### Method 3: Using One-Line Command

**Windows (PowerShell):**
```powershell
mysql -u root -p -e "CREATE DATABASE football_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
```

**Linux/macOS:**
```bash
mysql -u root -p -e "CREATE DATABASE football_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
```

---

### Database Migrations

**Good News:** The application automatically creates all tables when it starts!

GORM will auto-migrate these tables:
- ✅ `users`
- ✅ `teams`
- ✅ `players`
- ✅ `matches`
- ✅ `goals`

You'll see this in the logs when the server starts:
```
[2024-10-18 10:00:00] Database connected successfully
[2024-10-18 10:00:00] Auto-migration completed
```

---

## 5. Running the Application

### Development Mode (Recommended)

#### Option 1: Using `go run`

```bash
# Run directly without building
go run cmd/api/main.go
```

**Advantages:**
- ✅ Faster development
- ✅ No need to rebuild
- ✅ Good for testing changes

---

#### Option 2: Build and Run

**Windows (PowerShell):**
```powershell
# Step 1: Build
go build -o .\bin\main.exe .\cmd\api\main.go

# Step 2: Run
.\bin\main.exe
```

**Linux/macOS:**
```bash
# Step 1: Build
go build -o ./bin/main ./cmd/api/main.go

# Step 2: Run
./bin/main
```

**Advantages:**
- ✅ Faster startup
- ✅ Single executable file
- ✅ Can be distributed

---

### Production Mode

**Step 1: Set production mode in `.env`:**
```env
GIN_MODE=release
```

**Step 2: Build optimized binary:**

**Windows:**
```powershell
go build -ldflags="-s -w" -o .\bin\main.exe .\cmd\api\main.go
```

**Linux/macOS:**
```bash
go build -ldflags="-s -w" -o ./bin/main ./cmd/api/main.go
```

**Flags explained:**
- `-s` - Omit symbol table (smaller binary)
- `-w` - Omit DWARF debug info (smaller binary)

**Step 3: Run:**
```bash
# Windows
.\bin\main.exe

# Linux/macOS
./bin/main
```

---

### Running in Background

**Windows:**
```powershell
# Using PowerShell job
Start-Job -ScriptBlock { .\bin\main.exe }

# Check status
Get-Job

# Stop
Stop-Job -Id 1
```

**Linux/macOS:**
```bash
# Using nohup
nohup ./bin/main > app.log 2>&1 &

# Check if running
ps aux | grep main

# Stop (replace PID)
kill <PID>
```

---

## 6. Verification

### Step 1: Check Server Logs

After running the application, you should see:

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.

[2024-10-18 10:00:00] Database connected successfully
[2024-10-18 10:00:00] Auto-migration completed
[2024-10-18 10:00:00] Server starting on :8080

[GIN-debug] GET    /api/login                --> handler
[GIN-debug] GET    /api/teams                --> handler
[GIN-debug] POST   /api/teams                --> handler
...
[GIN-debug] Listening and serving HTTP on :8080
```

✅ **Success indicators:**
- Database connected
- Auto-migration completed
- All routes registered
- Server listening on port 8080

---

### Step 2: Test API Endpoints

#### Test 1: Health Check

**Using cURL:**
```bash
curl http://localhost:8080/api/health
```

**Expected Response:**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "OK"
  },
  "data": {
    "status": "healthy"
  }
}
```

---

#### Test 2: Login

**Using cURL:**
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"admin\",\"password\":\"admin123\"}"
```

**Expected Response:**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Login successful"
  },
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@football.com",
      "full_name": "Administrator",
      "photo_url": ""
    }
  }
}
```

---

#### Test 3: Get Teams (with Auth)

**Step 1: Copy token from login response**

**Step 2: Use token in request:**
```bash
curl -X GET http://localhost:8080/api/teams \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

---

### Step 3: Test in Browser

1. Open browser
2. Navigate to: `http://localhost:8080/api/health`
3. You should see JSON response

---

### Step 4: Import to Postman

1. **Open Postman**
2. **Import Collection**
   - Click **Import** button
   - Select file or URL
3. **Setup Environment**
   - Create new environment "Football API - Local"
   - Add variables:
     - `base_url`: `http://localhost:8080/api`
     - `token`: (leave empty, will auto-populate)
4. **Test Login**
   - Open "Login" request
   - Click **Send**
   - Token should auto-save to environment

---

## 7. Troubleshooting

### Problem: "go: command not found"

**Solution:**
```bash
# Check if Go is installed
go version

# If not installed, install Go from https://golang.org/dl/
# After installation, restart terminal/PowerShell
```

---

### Problem: "Error 1045: Access denied for user 'root'@'localhost'"

**Solution 1: Wrong password**
```env
# Update .env with correct password
DB_PASSWORD=your_correct_password
```

**Solution 2: Reset MySQL password**
```bash
# Login to MySQL as root
sudo mysql

# Change password
ALTER USER 'root'@'localhost' IDENTIFIED BY 'new_password';
FLUSH PRIVILEGES;
exit;
```

---

### Problem: "Unknown database 'football_db'"

**Solution:**
```bash
# Create database
mysql -u root -p -e "CREATE DATABASE football_db;"
```

---

### Problem: "bind: address already in use"

**Solution 1: Change port in .env**
```env
SERVER_PORT=8081  # Use different port
```

**Solution 2: Kill process using port 8080**

**Windows:**
```powershell
# Find process
netstat -ano | findstr :8080

# Kill process (replace PID)
taskkill /PID <PID> /F
```

**Linux/macOS:**
```bash
# Find process
lsof -i :8080

# Kill process (replace PID)
kill -9 <PID>
```

---

### Problem: "no such file or directory: uploads/"

**Solution:**
```bash
# Create upload directories
# Windows
New-Item -ItemType Directory -Force -Path uploads/teams
New-Item -ItemType Directory -Force -Path uploads/players
New-Item -ItemType Directory -Force -Path uploads/profiles

# Linux/macOS
mkdir -p uploads/{teams,players,profiles}
```

---

### Problem: Database connection timeout

**Solution 1: Check MySQL is running**
```bash
# Windows
Get-Service MySQL*

# Linux/macOS
sudo systemctl status mysql
```

**Solution 2: Update .env with correct host**
```env
DB_HOST=127.0.0.1  # Try this instead of localhost
```

---

### Problem: "Unauthorized" error

**Solution:**
```bash
# 1. Login again to get fresh token
# 2. Copy full token including "Bearer " prefix
# 3. Check token hasn't expired (default: 24 hours)
```

---

### Problem: Build errors

**Solution:**
```bash
# Clean module cache
go clean -modcache

# Re-download dependencies
go mod download

# Verify
go mod verify

# Try build again
go build cmd/api/main.go
```

---

## ✅ Installation Checklist

Use this checklist to verify your installation:

- [ ] Go 1.24+ installed and verified
- [ ] MySQL 8.0+ installed and running
- [ ] Git installed
- [ ] Repository cloned
- [ ] Dependencies downloaded (`go mod download`)
- [ ] `.env` file created from `.env.example`
- [ ] Database credentials configured in `.env`
- [ ] JWT secret configured (changed from default)
- [ ] Upload directories created
- [ ] Database `football_db` created
- [ ] Application builds without errors
- [ ] Server starts successfully
- [ ] Health check endpoint responds
- [ ] Login endpoint works
- [ ] Postman collection imported (optional)
- [ ] Documentation reviewed

---

## 🎉 Success!

If you've reached this point, congratulations! 🎊

Your **Football Management API** is now running and ready to use!

### Next Steps:

1. 📖 Read the [API Documentation](./API_DOCUMENTATION.md)
2. 🧪 Test endpoints using [Postman](./POSTMAN_OVERVIEW.md)
3. 💻 Start building your frontend application
4. 🚀 Deploy to production when ready

### Need Help?

- 📧 Email: support@footballapi.com
- 🐛 Issues: [GitHub Issues](https://github.com/alwijein/be_football/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/alwijein/be_football/discussions)

---

<div align="center">

**Happy Coding!** ⚽🚀

</div>
