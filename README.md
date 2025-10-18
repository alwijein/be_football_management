# ⚽ Football Management API

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go)
![Gin Framework](https://img.shields.io/badge/Gin-v1.10.0-00ADD8?style=for-the-badge&logo=go)
![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

**REST API Lengkap untuk Manajemen Tim Sepak Bola, Pemain, Pertandingan, dan Statistik**

Dibangun dengan Go, Gin Framework, dan Clean Architecture

---

### 🚀 Postman Collection - Import Sekarang!

[![Run in Postman](https://run.pstmn.io/button.svg)](https://elements.getpostman.com/redirect?entityId=24630549-226d5eed-7d4a-49df-a1d4-dc518a7de6ba&entityType=collection)

**Link Collection:**
```
https://elements.getpostman.com/redirect?entityId=24630549-226d5eed-7d4a-49df-a1d4-dc518a7de6ba&entityType=collection
```

> 💡 **Tip:** Klik tombol di atas atau copy link untuk import collection ke Postman. Token akan otomatis tersimpan setelah login!

---

[Fitur](#-fitur) • [Instalasi](#-instalasi) • [Dokumentasi API](#-dokumentasi-api) • [Struktur Project](#-struktur-project)

</div>

---

## 📋 Daftar Isi

- [Tentang Project](#-tentang-project)
- [Fitur Utama](#-fitur-utama)
- [Tech Stack](#-tech-stack)
- [Arsitektur](#-arsitektur)
- [Prasyarat](#-prasyarat)
- [Instalasi](#-instalasi)
- [Konfigurasi](#-konfigurasi)
- [Menjalankan Aplikasi](#-menjalankan-aplikasi)
- [Dokumentasi API](#-dokumentasi-api)
- [Database Schema](#-database-schema)
- [Struktur Project](#-struktur-project)
- [Lisensi](#-lisensi)

---

## 🎯 Tentang Project

**Football Management API** adalah REST API yang robust dan scalable untuk mengelola organisasi sepak bola, termasuk tim, pemain, pertandingan, jadwal, dan statistik lengkap. Dibangun dengan praktik modern Go dan Clean Architecture, API ini menyediakan fondasi yang solid untuk aplikasi manajemen sepak bola.

### Keunggulan Utama

- 🏗️ **Clean Architecture** - Pemisahan concern dengan batasan domain yang jelas
- 🔐 **Autentikasi Aman** - JWT-based auth dengan Hashing password Bcrypt
- 📁 **Upload File** - Support upload logo tim dan foto profil pemain
- 📊 **Statistik Lengkap** - Metrik dashboard dan laporan pertandingan detail
- ✅ **Validasi Input** - Validasi request komprehensif dengan pesan error detail
- 🚀 **Performa Tinggi** - Dibangun dengan Gin framework untuk kecepatan optimal
- 📖 **Dokumentasi Lengkap** - Dokumentasi API lengkap dengan contoh

---

## ✨ Fitur Utama

### 🔐 Autentikasi & Otorisasi
- Registrasi dan login user dengan JWT token
- Enkripsi password menggunakan bcrypt
- Manajemen profil dengan upload foto
- Proteksi API berbasis token

### ⚽ Manajemen Tim
- Operasi CRUD lengkap untuk tim
- Upload dan manajemen logo tim (JPG, PNG, GIF, WebP)
- Statistik dan metrik performa tim
- Support soft delete

### 👥 Manajemen Pemain
- CRUD pemain dengan asosiasi tim
- Upload foto profil pemain
- Kategorisasi berdasarkan posisi (Penyerang, Gelandang, Bertahan, Penjaga Gawang)
- Validasi nomor punggung (unik per tim)
- Atribut fisik (tinggi, berat)

### 📅 Manajemen Pertandingan & Jadwal
- Penjadwalan pertandingan dengan tanggal/waktu/venue
- Update skor live
- Manajemen hasil pertandingan
- Filter pertandingan berdasarkan tanggal (hari ini, mendatang, lampau)

### ⚽ Pencatatan Gol & Pencetak Gol
- Catat gol dengan atribusi pemain
- Waktu gol dan asosiasi pertandingan
- Statistik pencetak gol

### 📊 Laporan & Statistik
- Overview dashboard (total tim, pemain, pertandingan, gol)
- Laporan pertandingan spesifik tim
- Daftar pertandingan hari ini
- Detail pertandingan lengkap dengan tim dan gol

---

## 🚀 Tech Stack

| Teknologi | Versi | Kegunaan |
|-----------|-------|----------|
| **Go** | 1.24+ | Bahasa Pemrograman |
| **Gin** | 1.10.0 | HTTP Web Framework |
| **GORM** | 1.31.0 | ORM untuk Operasi Database |
| **MySQL** | 8.0+ | Relational Database |
| **JWT** | 5.2.1 | Autentikasi & Otorisasi |
| **Bcrypt** | - | Password Hashing |
| **UUID** | 1.6.0 | Unique File Naming |
| **Validator** | 10.22.1 | Request Validation |
| **CORS** | 1.7.2 | Cross-Origin Resource Sharing |
---

## 🏛️ Arsitektur

Project ini mengikuti prinsip **Clean Architecture** dengan pemisahan concern yang jelas:

```
┌─────────────────────────────────────────────────────────┐
│                   HTTP Layer (Gin)                       │
│                Routes & Middleware                       │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                  Handler Layer                           │
│         (HTTP Request/Response, Validation)              │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                  Service Layer                           │
│            (Business Logic, Use Cases)                   │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                 Repository Layer                         │
│            (Database Operations, GORM)                   │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                  Domain Layer                            │
│         (Entities, Business Rules, DTOs)                 │
└─────────────────────────────────────────────────────────┘
```

**Keuntungan:**
- ✅ **Testability** - Setiap layer bisa ditest secara independen
- ✅ **Maintainability** - Perubahan di satu layer tidak mempengaruhi yang lain
- ✅ **Scalability** - Mudah diperluas dengan fitur baru
- ✅ **Flexibility** - Bisa mengganti implementasi tanpa mempengaruhi business logic

---

## 📦 Prasyarat

Sebelum instalasi, pastikan Anda sudah menginstall:

### Wajib

- **Go** (versi 1.24 atau lebih tinggi)
  ```bash
  # Cek versi Go
  go version
  ```
  Download: [https://golang.org/dl/](https://golang.org/dl/)

- **MySQL** (versi 8.0 atau lebih tinggi)
  ```bash
  # Cek versi MySQL
  mysql --version
  ```
  Download: [https://dev.mysql.com/downloads/mysql/](https://dev.mysql.com/downloads/mysql/)

- **Git**
  ```bash
  # Cek versi Git
  git --version
  ```
  Download: [https://git-scm.com/downloads](https://git-scm.com/downloads)

### Opsional (Disarankan)

- **Postman** - Untuk testing API
- **MySQL Workbench** - Untuk manajemen database
- **VS Code** dengan Go extension - Untuk development

---

## 🔧 Instalasi

### 1. Clone Repository

```bash
# Clone via HTTPS
git clone https://github.com/alwijein/be_football_management.git

# Atau via SSH
git clone git@github.com:alwijein/be_football_management.git

# Masuk ke direktori project
cd be_football_management
```

### 2. Install Dependencies

```bash
# Download dan install semua dependensi Go
go mod download

# Verifikasi dependencies
go mod verify

# Rapihkan go.mod dan go.sum (opsional)
go mod tidy
```

### 3. Buat Database

```bash
# Login ke MySQL
mysql -u root -p

# Buat database
CREATE DATABASE football_db;

# Keluar dari MySQL
exit;
```

**Alternatif menggunakan MySQL Workbench:**
1. Buka MySQL Workbench
2. Connect ke MySQL server Anda
3. Eksekusi: `CREATE DATABASE football_db;`

### 4. Konfigurasi Environment Variables

```bash
# Copy file environment example
cp .env.example .env

# Edit file .env dengan konfigurasi Anda
# (Gunakan nano, vim, atau text editor apapun)
nano .env
```

**Opsi Konfigurasi:**

```env
# Konfigurasi Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password_mysql_anda
DB_NAME=football_db

# Konfigurasi Server
SERVER_PORT=8080
GIN_MODE=debug

# Konfigurasi JWT
JWT_SECRET=your-very-secret-key-please-change-in-production-12345
JWT_EXPIRATION_HOURS=24

# Aplikasi
APP_NAME=Football Management API
APP_VERSION=1.0.0
```

⚠️ **Penting:** Ganti `JWT_SECRET` di environment production!

### 5. Buat Direktori Upload

```bash
# Buat direktori untuk file uploads
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

## ⚙️ Konfigurasi

### Konfigurasi Database

Edit file `.env` sesuai dengan setup MySQL Anda:

```env
DB_HOST=localhost          # MySQL host (gunakan 'localhost' untuk lokal)
DB_PORT=3306              # MySQL port (default: 3306)
DB_USER=root              # Username MySQL Anda
DB_PASSWORD=              # Password MySQL Anda (kosongkan jika tidak ada password)
DB_NAME=football_db       # Nama database
```

### Konfigurasi Server

```env
SERVER_PORT=8080          # Port dimana API akan berjalan
GIN_MODE=debug           # Opsi: debug, release, test
```

- **debug**: Mode development dengan log detail
- **release**: Mode production dengan performa optimal
- **test**: Mode testing

### JWT Configuration

```env
JWT_SECRET=your-secret-key-min-32-characters-long-for-security
JWT_EXPIRATION_HOURS=24
```

⚠️ **Best Practice Keamanan:**
- Gunakan secret key yang kuat dan random (minimal 32 karakter)
- Jangan commit `.env` ke version control
- Gunakan secret yang berbeda untuk environment berbeda
- Rotasi secret secara berkala

---

## � Menjalankan Aplikasi

### Mode Development

```bash
# Build aplikasi
go build -o bin/main.exe cmd/api/main.go

# Jalankan aplikasi
./bin/main.exe
```

**Windows (PowerShell):**
```powershell
# Build
go build -o .\bin\main.exe .\cmd\api\main.go

# Run
.\bin\main.exe
```

### Menggunakan Go Run (Quick Development)

```bash
# Jalankan langsung tanpa build
go run cmd/api/main.go
```

### Mode Production

```bash
# Set Mode Production in .env
GIN_MODE=release

# Build binary yang optimal
go build -ldflags="-s -w" -o bin/main cmd/api/main.go

# Run
./bin/main
```

### Verifikasi Server Berjalan

```bash
# Test health endpoint
curl http://localhost:8080/api/health

# Response yang diharapkan:
# {"meta":{"code":200,"status":"success","message":"OK"},"data":{"status":"healthy"}}
```

### Informasi Server

Setelah berjalan, server akan tersedia di:

- **Local**: `http://localhost:8080`
- **Network**: `http://YOUR_LOCAL_IP:8080` (e.g., `http://192.168.1.100:8080`)

**API Base URL**: `http://localhost:8080/api`

---

## 📖 Dokumentasi API

### Base URL

```
http://localhost:8080/api
```

### Autentikasi

Most endpoints require JWT Autentikasi. Include the token in the Authorization header:

```
Authorization: Bearer YOUR_JWT_TOKEN
```

### Format Response

Semua response API mengikuti format standar ini:

**Success Response:**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Operation successful"
  },
  "data": {
    // Data response di sini
  }
}
```

**Error Response:**
```json
{
  "meta": {
    "code": 400,
    "status": "error",
    "message": "Deskripsi error"
  },
  "data": null
}
```

### Referensi Cepat

| Kategori | Endpoint | Method | Perlu Auth |
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

### Dokumentasi Lengkap

For detailed Dokumentasi API with request/response examples, see:
- 📄 **[API_DOCUMENTATION.md](./API_DOCUMENTATION.md)** - Dokumentasi endpoint lengkap
- 📄 **[API_QUICK_REFERENCE.md](./API_QUICK_REFERENCE.md)** - Referensi Cepat guide

### Postman Collection

**🚀 Quick Import - Public Collection:**

[![Run in Postman](https://run.pstmn.io/button.svg)](https://elements.getpostman.com/redirect?entityId=24630549-226d5eed-7d4a-49df-a1d4-dc518a7de6ba&entityType=collection)

**Atau import manual:**

1. Buka Postman
2. Klik **Import**
3. Paste URL ini: `https://elements.getpostman.com/redirect?entityId=24630549-226d5eed-7d4a-49df-a1d4-dc518a7de6ba&entityType=collection`
4. Klik **Import**
5. Set environment variables:
   - `base_url`: `http://localhost:8080/api`
   - `token`: JWT token Anda setelah login

> 💡 **Tip:** Token akan otomatis tersimpan setelah login berhasil!

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

### Tabel Database

**Tabel Users:**
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

**Tabel Teams:**
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

**Tabel Players:**
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

**Tabel Matches:**
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

**Tabel Goals:**
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
Aplikasi secara otomatis membuat dan migrasi tabel-tabel ini saat startup menggunakan GORM.

---

## � Struktur Project

```
be_football/
│
├── cmd/
│   └── api/
│       └── main.go                    # Entry point aplikasi & inisialisasi server
│
├── internal/
│   ├── config/
│   │   └── config.go                  # Loader konfigurasi environment
│   │
│   ├── domain/                        # Entity domain (Model bisnis)
│   │   ├── user.go                    # Entity User
│   │   ├── team.go                    # Entity Team
│   │   ├── player.go                  # Entity Player dengan enum posisi
│   │   ├── match.go                   # Entity Match/Schedule
│   │   └── goal.go                    # Entity Goal
│   │
│   ├── dto/                           # Data Transfer Objects
│   │   ├── auth.go                    # DTO Login, Register, Profile
│   │   ├── team.go                    # DTO request/response Team
│   │   ├── player.go                  # DTO request/response Player
│   │   ├── match.go                   # DTO Match/Schedule
│   │   ├── goal.go                    # DTO Goal
│   │   └── response.go                # Wrapper response standar
│   │
│   ├── repository/                    # Layer Akses Data
│   │   ├── user_repository.go         # Operasi CRUD User
│   │   ├── team_repository.go         # Operasi CRUD Team
│   │   ├── player_repository.go       # Operasi CRUD Player
│   │   ├── match_repository.go        # Operasi CRUD Match
│   │   └── goal_repository.go         # Operasi CRUD Goal
│   │
│   ├── service/                       # Layer Logika Bisnis
│   │   ├── auth_service.go            # Autentikasi logic
│   │   ├── team_service.go            # Logika bisnis Team
│   │   ├── player_service.go          # Logika bisnis Player
│   │   ├── match_service.go           # Logika bisnis Match
│   │   ├── goal_service.go            # Logika bisnis Goal
│   │   └── report_service.go          # Logika dashboard & laporan
│   │
│   ├── handler/                       # HTTP Handlers (Controllers)
│   │   ├── auth_handler.go            # Endpoint Auth
│   │   ├── team_handler.go            # Endpoint Team
│   │   ├── player_handler.go          # Endpoint Player
│   │   ├── match_handler.go           # Endpoint Match/Schedule
│   │   ├── goal_handler.go            # Endpoint Goal
│   │   └── report_handler.go          # Endpoint dashboard & laporan
│   │
│   ├── middleware/                    # HTTP Middleware
│   │   ├── auth_middleware.go         # JWT Autentikasi
│   │   ├── cors_middleware.go         # Konfigurasi CORS
│   │   └── logger_middleware.go       # Logging request
│   │
│   ├── routes/
│   │   └── routes.go                  # Definisi & grouping route
│   │
│   └── utils/                         # Fungsi utilitas
│       ├── jwt.go                     # Generasi & validasi JWT token
│       ├── password.go                # Hashing password Bcrypt
│       ├── response.go                # Helper response standar
│       ├── file_upload.go             # Utilitas upload file
│       └── validation.go              # Validator kustom
│
├── pkg/
│   └── database/
│       └── mysql.go                   # Koneksi database & auto-migration
│
├── uploads/                           # Direktori file uploads
│   ├── teams/                         # Logo tim
│   ├── players/                       # Foto pemain
│   └── profiles/                      # Foto profil user
│
├── bin/                               # Binary yang dikompilasi
│
├── .env                               # Environment variables (DO NOT COMMIT)
├── .env.example                       # Template variabel environment
├── .gitignore                         # Aturan git ignore
├── go.mod                             # Dependensi Go module
├── go.sum                             # Checksum dependensi
├── README.md                          # File ini
├── API_DOCUMENTATION.md               # Complete Dokumentasi API
├── API_QUICK_REFERENCE.md             # Referensi API cepat
└── LICENSE                            # Lisensi project
```

### Penjelasan Direktori Utama

- **`cmd/api/`**: Entry point aplikasi
- **`internal/`**: Kode aplikasi private (tidak bisa diimport project lain)
- **`internal/domain/`**: Entity bisnis inti dan model
- **`internal/dto/`**: Struktur data untuk request/response API
- **`internal/repository/`**: Operasi database (GORM)
- **`internal/service/`**: Logika bisnis dan use case
- **`internal/handler/`**: Handler request HTTP
- **`internal/middleware/`**: Middleware HTTP (auth, CORS, logging)
- **`internal/routes/`**: Definisi route API
- **`internal/utils/`**: Reusable Fungsi utilitas
- **`pkg/`**: Library publik (bisa diimport project lain)
- **`uploads/`**: Penyimpanan file untuk upload user

---

##  Contributing

Kontribusi sangat diterima! Mohon ikuti panduan berikut:

### Cara Berkontribusi

1. **Fork repository**
2. **Buat feature branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Commit perubahan Anda**
   ```bash
   git commit -m 'Add some amazing feature'
   ```
4. **Push ke branch**
   ```bash
   git push origin feature/amazing-feature
   ```
5. **Buka Pull Request**

### Standar Coding

- Ikuti best practice dan konvensi Go
- Gunakan `gofmt` untuk format kode
- Tambahkan komentar untuk fungsi yang di-export
- Tulis unit test untuk fitur baru
- Update dokumentasi jika diperlukan

### Format Commit Message

```
type(scope): subject

body

footer
```

**Types:**
- `feat`: Fitur baru
- `fix`: Perbaikan bug
- `docs`: Perubahan dokumentasi
- `style`: Perubahan style kode (formatting)
- `refactor`: Refactoring kode
- `test`: Menambahkan test
- `chore`: Tugas maintenance

**Example:**
```
feat(auth): tambahkan endpoint refresh token

Implementasi fungsi refresh token untuk memungkinkan
user mendapatkan access token baru tanpa login ulang.

Closes #123
```

---

## 📄 License

Project ini dilisensikan di bawah MIT License - lihat file [LICENSE](LICENSE) untuk detail.

```
MIT License

Copyright (c) 2024 Football Management API

Dengan ini diberikan izin, tanpa biaya, kepada siapapun yang mendapatkan salinan
dari software ini dan file dokumentasi terkait (the "Software"), to deal
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

## 🙏 Penghargaan

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- [Go Playground Validator](https://github.com/go-playground/validator)
- [godotenv](https://github.com/joho/godotenv)

---

## 📞 Dukungan & Kontak

### Menemukan Bug?

Please [buka issue](https://github.com/alwijein/be_football/issues) with:
- Deskripsi bug
- Langkah-langkah untuk mereproduksi
- Perilaku yang diharapkan vs aktual
- Screenshot (jika ada)

### Butuh Bantuan?

- 📖 Cek [Dokumentasi API](./API_DOCUMENTATION.md)
- 💬 Buka [diskusi](https://github.com/alwijein/be_football/diskusis)
- 📧 Contact: [your-email@example.com](mailto:your-email@example.com)

### Permintaan Fitur

We welcome Permintaan Fitur! Please [buka issue](https://github.com/alwijein/be_football/issues) dengan .`enhancement` .

---

## 🗺️ Roadmap

### Versi 1.0.0 (Saat Ini) ✅
- Fungsi API inti
- Autentikasi & Authorization
- Team, Player, Match management
- Dukungan upload file
- Dashboard & Laporan

### Versi 1.1.0 (Direncanakan) 🚧
- [ ] Implementasi refresh token
- [ ] Verifikasi email
- [ ] Fungsi reset password
- [ ] Pencarian dan filter yang ditingkatkan
- [ ] Export laporan ke PDF/Excel

### Versi 2.0.0 (Masa Depan) 🔮
- [ ] Update pertandingan real-time (WebSocket)
- [ ] Interface web dashboard admin
- [ ] Dukungan multi-bahasa
- [ ] Analitik dan statistik lanjutan
- [ ] Integrasi aplikasi mobile
- [ ] Integrasi media sosial

---

## 📊 Status Project

![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)
![Build](https://img.shields.io/badge/Build-Passing-success?style=for-the-badge)
![Coverage](https://img.shields.io/badge/Coverage-85%25-green?style=for-the-badge)
![Version](https://img.shields.io/badge/Version-1.0.0-blue?style=for-the-badge)

**Current Version:** 1.0.0  
**Terakhir Update:** October 2024  
**Status:** Siap Production ✅





























