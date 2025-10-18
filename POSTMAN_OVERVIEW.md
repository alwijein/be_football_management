# ⚽ Football Management API - Postman Documentation

<div align="center">

**Complete REST API for Football Team Management**

Built with Go, Gin Framework, and MySQL

**Base URL:** `http://localhost:8080/api`

[Quick Start](#quick-start) • [Authentication](#authentication) • [Endpoints](#api-endpoints) • [Examples](#request-examples)

</div>

---

## 📖 Overview

The **Football Management API** is a comprehensive REST API designed to manage football organizations, including teams, players, matches, schedules, and detailed statistics. This API provides all the essential features needed for a complete football management system.

### Key Features

- 🔐 **JWT Authentication** - Secure authentication with token-based access
- ⚽ **Team Management** - Complete CRUD operations with logo upload support
- 👥 **Player Management** - Player profiles with photos and team associations
- 📅 **Match Scheduling** - Schedule and track matches with live updates
- ⚽ **Goal Tracking** - Record goals with player attribution
- 📊 **Statistics Dashboard** - Comprehensive statistics and reports
- 📁 **File Uploads** - Support for team logos, player photos, and profile pictures
- ✅ **Input Validation** - Robust request validation with detailed error messages
- 📱 **Mobile-Ready** - RESTful design perfect for mobile apps

---

## 🎯 Quick Start

### 1. Setup Environment

```bash
# Clone the repository
git clone https://github.com/alwijein/be_football.git
cd be_football

# Install dependencies
go mod download

# Create database
CREATE DATABASE football_db;

# Configure .env file
cp .env.example .env
```

### 2. Run the Server

```bash
# Build and run
go build -o bin/main.exe cmd/api/main.go
./bin/main.exe

# Server will start at:
# http://localhost:8080
```

### 3. Import to Postman

1. Open Postman
2. Click **Import**
3. Select **Link** or **File**
4. Import the collection
5. Create environment with variables:
   - `base_url`: `http://localhost:8080/api`
   - `token`: (will be auto-set after login)

---

## 🔐 Authentication

### How Authentication Works

1. **Register** or **Login** to get a JWT token
2. Token is automatically saved to environment variable `{{token}}`
3. Protected endpoints automatically use `Bearer {{token}}`
4. Token expires in 24 hours (configurable)

### Authentication Flow

```
┌──────────┐          ┌──────────┐          ┌──────────┐
│  Client  │          │   API    │          │ Database │
└────┬─────┘          └────┬─────┘          └────┬─────┘
     │                     │                     │
     │  POST /login        │                     │
     │ ──────────────────> │                     │
     │                     │  Verify credentials │
     │                     │ ──────────────────> │
     │                     │                     │
     │                     │ <─────────────────  │
     │  JWT Token          │      User data      │
     │ <────────────────── │                     │
     │                     │                     │
     │  GET /teams         │                     │
     │  (with Bearer token)│                     │
     │ ──────────────────> │                     │
     │                     │  Verify token       │
     │                     │  Get teams          │
     │                     │ ──────────────────> │
     │                     │                     │
     │  Teams data         │ <─────────────────  │
     │ <────────────────── │                     │
     │                     │                     │
```

### Default Test User

```json
{
  "username": "admin",
  "password": "admin123"
}
```

---

## 📋 API Endpoints

### Summary

| Category | Endpoints | Description |
|----------|-----------|-------------|
| 🔐 **Authentication** | 5 | Login, register, profile management |
| 📊 **Dashboard** | 2 | Statistics and today's matches |
| ⚽ **Teams** | 7 | Team CRUD and team-specific operations |
| 👥 **Players** | 4 | Player CRUD and team association |
| 📅 **Schedules** | 5 | Match scheduling and results |
| 🎯 **Matches** | 2 | Match listings and details |
| ⚽ **Goals** | 2 | Goal recording and management |

**Total:** 26 Endpoints

---

## 🔐 1. Authentication Endpoints

### 1.1 Login
Authenticate user and receive JWT token.

**Endpoint:** `POST /api/login`  
**Auth Required:** ❌ No

**Request Body:**
```json
{
  "username": "admin",
  "password": "admin123"
}
```

**Success Response (200):**
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
      "photo_url": "http://localhost:8080/uploads/profiles/photo.jpg"
    }
  }
}
```

---

### 1.2 Register
Create a new user account.

**Endpoint:** `POST /api/register`  
**Auth Required:** ❌ No

**Request Body:**
```json
{
  "username": "newuser",
  "email": "user@example.com",
  "password": "Password123!",
  "full_name": "John Doe"
}
```

---

### 1.3 Get Profile
Get current user profile information.

**Endpoint:** `GET /api/profile`  
**Auth Required:** ✅ Yes

**Success Response (200):**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Profile retrieved successfully"
  },
  "data": {
    "id": 1,
    "username": "admin",
    "email": "admin@football.com",
    "full_name": "Administrator",
    "photo_url": "http://localhost:8080/uploads/profiles/photo.jpg"
  }
}
```

---

### 1.4 Update Profile
Update user profile with optional photo upload.

**Endpoint:** `PUT /api/profile`  
**Auth Required:** ✅ Yes  
**Content-Type:** `multipart/form-data`

**Form Data:**
- `email` (optional): New email address
- `full_name` (optional): New full name
- `photo` (optional): Profile photo file (max 5MB, JPG/PNG/GIF/WebP)

---

### 1.5 Change Password
Change user password.

**Endpoint:** `PUT /api/change-password`  
**Auth Required:** ✅ Yes

**Request Body:**
```json
{
  "old_password": "OldPassword123!",
  "new_password": "NewPassword123!",
  "confirm_password": "NewPassword123!"
}
```

---

## 📊 2. Dashboard Endpoints

### 2.1 Get Dashboard Statistics
Get overview statistics of the system.

**Endpoint:** `GET /api/dashboard`  
**Auth Required:** ✅ Yes

**Success Response (200):**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Dashboard data retrieved successfully"
  },
  "data": {
    "total_teams": 12,
    "total_players": 264,
    "total_matches": 45,
    "total_goals": 127
  }
}
```

---

### 2.2 Get Today's Matches
Get list of matches scheduled for today.

**Endpoint:** `GET /api/matches/today`  
**Auth Required:** ✅ Yes

**Success Response (200):**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Today's matches retrieved successfully"
  },
  "data": [
    {
      "id": 15,
      "home_team": {
        "id": 3,
        "name": "Arsenal FC",
        "logo": "http://localhost:8080/uploads/teams/arsenal.png"
      },
      "away_team": {
        "id": 7,
        "name": "Chelsea FC",
        "logo": "http://localhost:8080/uploads/teams/chelsea.png"
      },
      "match_date": "2024-10-18T15:00:00Z",
      "venue": "Emirates Stadium",
      "home_score": 0,
      "away_score": 0,
      "status": "scheduled"
    }
  ]
}
```

---

## ⚽ 3. Team Endpoints

### 3.1 Get All Teams
Get paginated list of all teams.

**Endpoint:** `GET /api/teams`  
**Auth Required:** ✅ Yes

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Example:** `GET /api/teams?page=1&limit=10`

**Success Response (200):**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": null,
    "pagination": {
      "current_page": 1,
      "total_pages": 2,
      "total_items": 12,
      "per_page": 10
    }
  },
  "data": [
    {
      "id": 1,
      "name": "Manchester United",
      "city": "Manchester",
      "logo": "http://localhost:8080/uploads/teams/manutd.png",
      "established_year": 1878,
      "created_at": "2024-10-01T10:00:00Z",
      "player_count": 25
    }
  ]
}
```

---

### 3.2 Get Team Detail
Get detailed information of a specific team.

**Endpoint:** `GET /api/teams/:id`  
**Auth Required:** ✅ Yes

**Example:** `GET /api/teams/1`

**Success Response (200):**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Team retrieved successfully"
  },
  "data": {
    "id": 1,
    "name": "Manchester United",
    "city": "Manchester",
    "logo": "http://localhost:8080/uploads/teams/manutd.png",
    "established_year": 1878,
    "created_at": "2024-10-01T10:00:00Z",
    "players": [
      {
        "id": 1,
        "name": "Marcus Rashford",
        "position": "Penyerang",
        "jersey_number": 10,
        "height": 185,
        "weight": 80
      }
    ]
  }
}
```

---

### 3.3 Create Team
Create a new team with optional logo upload.

**Endpoint:** `POST /api/teams`  
**Auth Required:** ✅ Yes  
**Content-Type:** `multipart/form-data`

**Form Data:**
- `name` (required): Team name
- `city` (required): Team city
- `established_year` (required): Year established (1800-2100)
- `logo` (optional): Team logo file (max 5MB, JPG/PNG/GIF/WebP)

**Example:**
```
name: Arsenal FC
city: London
established_year: 1886
logo: [file]
```

**Success Response (201):**
```json
{
  "meta": {
    "code": 201,
    "status": "success",
    "message": "Team created successfully"
  },
  "data": {
    "id": 13,
    "name": "Arsenal FC",
    "city": "London",
    "logo": "http://localhost:8080/uploads/teams/1729174500_a1b2c3d4.jpg",
    "established_year": 1886,
    "created_at": "2024-10-18T10:15:00Z"
  }
}
```

---

### 3.4 Update Team
Update team information with optional logo change.

**Endpoint:** `PUT /api/teams/:id`  
**Auth Required:** ✅ Yes  
**Content-Type:** `multipart/form-data`

**Form Data:** (all optional)
- `name`: New team name
- `city`: New city
- `established_year`: New year
- `logo`: New logo file

---

### 3.5 Delete Team
Soft delete a team.

**Endpoint:** `DELETE /api/teams/:id`  
**Auth Required:** ✅ Yes

**Example:** `DELETE /api/teams/13`

**Success Response (200):**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Team deleted successfully"
  },
  "data": null
}
```

---

### 3.6 Get Team Players
Get all players belonging to a specific team.

**Endpoint:** `GET /api/teams/:id/players`  
**Auth Required:** ✅ Yes

**Example:** `GET /api/teams/1/players`

---

### 3.7 Create Team Player
Add a new player to a specific team.

**Endpoint:** `POST /api/teams/:id/players`  
**Auth Required:** ✅ Yes

**Request Body:**
```json
{
  "name": "Bruno Fernandes",
  "height": 179,
  "weight": 69,
  "position": "Gelandang",
  "jersey_number": 8
}
```

**Note:** `team_id` is automatically set from the URL parameter.

**Positions Available:**
- `Penyerang` (Forward)
- `Gelandang` (Midfielder)
- `Bertahan` (Defender)
- `Penjaga Gawang` (Goalkeeper)

---

### 3.8 Get Team Match Reports
Get match history and statistics for a specific team.

**Endpoint:** `GET /api/teams/:id/match-reports`  
**Auth Required:** ✅ Yes

**Example:** `GET /api/teams/1/match-reports`

**Success Response (200):**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Team match reports retrieved successfully"
  },
  "data": [
    {
      "id": 5,
      "home_team": {
        "id": 1,
        "name": "Manchester United",
        "logo": "http://localhost:8080/uploads/teams/manutd.png"
      },
      "away_team": {
        "id": 3,
        "name": "Arsenal FC",
        "logo": "http://localhost:8080/uploads/teams/arsenal.png"
      },
      "match_date": "2024-10-15T15:00:00Z",
      "venue": "Old Trafford",
      "home_score": 2,
      "away_score": 1,
      "status": "finished",
      "goals": [
        {
          "id": 8,
          "player_name": "Marcus Rashford",
          "minute": 23,
          "team": "home"
        }
      ]
    }
  ]
}
```

---

## 👥 4. Player Endpoints

### 4.1 Get Player Detail
Get detailed information of a specific player.

**Endpoint:** `GET /api/players/:id`  
**Auth Required:** ✅ Yes

**Success Response (200):**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Player retrieved successfully"
  },
  "data": {
    "id": 1,
    "name": "Marcus Rashford",
    "height": 185,
    "weight": 80,
    "position": "Penyerang",
    "jersey_number": 10,
    "team": {
      "id": 1,
      "name": "Manchester United",
      "logo": "http://localhost:8080/uploads/teams/manutd.png"
    }
  }
}
```

---

### 4.2 Update Player
Update player information.

**Endpoint:** `PUT /api/players/:id`  
**Auth Required:** ✅ Yes

**Request Body:** (all fields optional)
```json
{
  "name": "Marcus Rashford Jr.",
  "height": 186,
  "weight": 81,
  "position": "Penyerang",
  "jersey_number": 10
}
```

---

### 4.3 Delete Player
Soft delete a player.

**Endpoint:** `DELETE /api/players/:id`  
**Auth Required:** ✅ Yes

---

## 📅 5. Schedule/Match Endpoints

### 5.1 Get All Schedules
Get paginated list of all match schedules.

**Endpoint:** `GET /api/schedules`  
**Auth Required:** ✅ Yes

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

---

### 5.2 Get Schedule Detail
Get detailed information of a specific match schedule.

**Endpoint:** `GET /api/schedules/:id`  
**Auth Required:** ✅ Yes

---

### 5.3 Create Schedule
Create a new match schedule.

**Endpoint:** `POST /api/schedules`  
**Auth Required:** ✅ Yes

**Request Body:**
```json
{
  "home_team_id": 1,
  "away_team_id": 3,
  "match_date": "2024-10-25T15:00:00Z",
  "venue": "Old Trafford"
}
```

**Validation Rules:**
- `home_team_id` and `away_team_id` must be different
- `match_date` must be in the future
- Both teams must exist

---

### 5.4 Update Match Result
Update the score and status of a match.

**Endpoint:** `PUT /api/schedules/:id/result`  
**Auth Required:** ✅ Yes

**Request Body:**
```json
{
  "home_score": 2,
  "away_score": 1,
  "status": "finished"
}
```

**Status Options:**
- `scheduled` - Match not yet played
- `ongoing` - Match in progress
- `finished` - Match completed
- `postponed` - Match delayed
- `cancelled` - Match cancelled

---

### 5.5 Delete Schedule
Soft delete a match schedule.

**Endpoint:** `DELETE /api/schedules/:id`  
**Auth Required:** ✅ Yes

---

## 🎯 6. Match Endpoints

### 6.1 Get All Matches
Get paginated list of all matches with full details.

**Endpoint:** `GET /api/matches`  
**Auth Required:** ✅ Yes

**Query Parameters:**
- `page` (optional): Page number
- `limit` (optional): Items per page

---

### 6.2 Get Match Report
Get detailed match report with teams, scores, and goals.

**Endpoint:** `GET /api/matches/:id`  
**Auth Required:** ✅ Yes

**Success Response (200):**
```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": "Match report retrieved successfully"
  },
  "data": {
    "id": 5,
    "home_team": {
      "id": 1,
      "name": "Manchester United",
      "logo": "http://localhost:8080/uploads/teams/manutd.png"
    },
    "away_team": {
      "id": 3,
      "name": "Arsenal FC",
      "logo": "http://localhost:8080/uploads/teams/arsenal.png"
    },
    "match_date": "2024-10-15T15:00:00Z",
    "venue": "Old Trafford",
    "home_score": 2,
    "away_score": 1,
    "status": "finished",
    "goals": [
      {
        "id": 8,
        "player": {
          "id": 1,
          "name": "Marcus Rashford",
          "jersey_number": 10
        },
        "minute": 23,
        "team": "home"
      },
      {
        "id": 9,
        "player": {
          "id": 15,
          "name": "Bukayo Saka",
          "jersey_number": 7
        },
        "minute": 45,
        "team": "away"
      },
      {
        "id": 10,
        "player": {
          "id": 5,
          "name": "Bruno Fernandes",
          "jersey_number": 8
        },
        "minute": 78,
        "team": "home"
      }
    ]
  }
}
```

---

## ⚽ 7. Goal Endpoints

### 7.1 Create Goal
Record a goal for a match.

**Endpoint:** `POST /api/goals`  
**Auth Required:** ✅ Yes

**Request Body:**
```json
{
  "match_id": 5,
  "player_id": 1,
  "minute": 23
}
```

**Validation:**
- Match must exist and not be deleted
- Player must exist and belong to one of the teams in the match
- Minute must be between 1 and 120 (including extra time)

**Success Response (201):**
```json
{
  "meta": {
    "code": 201,
    "status": "success",
    "message": "Goal recorded successfully"
  },
  "data": {
    "id": 11,
    "match_id": 5,
    "player_id": 1,
    "player_name": "Marcus Rashford",
    "minute": 23,
    "created_at": "2024-10-18T15:25:00Z"
  }
}
```

---

### 7.2 Delete Goal
Remove a goal record.

**Endpoint:** `DELETE /api/goals/:id`  
**Auth Required:** ✅ Yes

---

## ✅ Response Format

### Success Response Structure

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

### Success with Pagination

```json
{
  "meta": {
    "code": 200,
    "status": "success",
    "message": null,
    "pagination": {
      "current_page": 1,
      "total_pages": 5,
      "total_items": 48,
      "per_page": 10
    }
  },
  "data": []
}
```

### Error Response Structure

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

---

## 🔢 HTTP Status Codes

| Code | Status | Description |
|------|--------|-------------|
| 200 | OK | Request successful |
| 201 | Created | Resource created successfully |
| 400 | Bad Request | Invalid request data |
| 401 | Unauthorized | Authentication required or failed |
| 403 | Forbidden | Access denied |
| 404 | Not Found | Resource not found |
| 409 | Conflict | Resource conflict (e.g., duplicate) |
| 422 | Unprocessable Entity | Validation error |
| 500 | Internal Server Error | Server error |

---

## 📁 File Upload Guidelines

### Supported File Types
- **Images**: JPG, JPEG, PNG, GIF, WebP

### File Size Limits
- **Maximum**: 5 MB per file

### Endpoints with File Upload
1. **POST /api/teams** - Team logo
2. **PUT /api/teams/:id** - Update team logo
3. **PUT /api/profile** - Profile photo

### Upload Example (cURL)

```bash
curl -X POST http://localhost:8080/api/teams \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "name=Arsenal FC" \
  -F "city=London" \
  -F "established_year=1886" \
  -F "logo=@/path/to/logo.png"
```

### Upload Example (Postman)

1. Select **POST** method
2. Choose **Body** tab
3. Select **form-data**
4. Add fields:
   - `name`: Arsenal FC (Text)
   - `city`: London (Text)
   - `established_year`: 1886 (Text)
   - `logo`: [Select File]

---

## 🧪 Testing Guide

### 1. Basic Authentication Test

```bash
# Step 1: Login
POST /api/login
Body: {
  "username": "admin",
  "password": "admin123"
}

# Step 2: Copy the token from response
# Step 3: Use token in subsequent requests
GET /api/teams
Header: Authorization: Bearer YOUR_TOKEN
```

### 2. Create Team with Logo

```bash
POST /api/teams
Header: Authorization: Bearer YOUR_TOKEN
Form-Data:
  - name: Test FC
  - city: Test City
  - established_year: 2000
  - logo: [upload file]
```

### 3. Complete Match Flow

```bash
# 1. Create two teams
POST /api/teams (Team A)
POST /api/teams (Team B)

# 2. Add players to teams
POST /api/teams/1/players
POST /api/teams/2/players

# 3. Schedule a match
POST /api/schedules
Body: {
  "home_team_id": 1,
  "away_team_id": 2,
  "match_date": "2024-10-25T15:00:00Z",
  "venue": "Stadium Name"
}

# 4. Record goals
POST /api/goals
Body: {
  "match_id": 1,
  "player_id": 1,
  "minute": 23
}

# 5. Update match result
PUT /api/schedules/1/result
Body: {
  "home_score": 2,
  "away_score": 1,
  "status": "finished"
}

# 6. View match report
GET /api/matches/1
```

---

## ❗ Common Errors

### 401 Unauthorized
**Cause:** Missing or invalid token  
**Solution:** Login again and use the new token

```json
{
  "meta": {
    "code": 401,
    "status": "error",
    "message": "Unauthorized"
  },
  "data": null
}
```

### 400 Bad Request
**Cause:** Invalid request data  
**Solution:** Check request body format and required fields

```json
{
  "meta": {
    "code": 400,
    "status": "error",
    "message": "Key: 'CreateTeamRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
  },
  "data": null
}
```

### 404 Not Found
**Cause:** Resource doesn't exist  
**Solution:** Verify the ID in the URL

```json
{
  "meta": {
    "code": 404,
    "status": "error",
    "message": "Team not found"
  },
  "data": null
}
```

---

## 🔄 Version History

### Version 1.0.0 (Current)
- ✅ Complete authentication system
- ✅ Team management with logo upload
- ✅ Player management
- ✅ Match scheduling and results
- ✅ Goal tracking
- ✅ Dashboard and statistics
- ✅ Match reports
- ✅ File upload support

---

## 📞 Support

### Need Help?
- 📖 Check the [Complete API Documentation](./API_DOCUMENTATION.md)
- 📧 Contact: support@footballapi.com
- 🐛 Report bugs: [GitHub Issues](https://github.com/alwijein/be_football/issues)

### Useful Resources
- [GitHub Repository](https://github.com/alwijein/be_football)
- [API Documentation](./API_DOCUMENTATION.md)
- [Quick Reference](./API_QUICK_REFERENCE.md)

---

<div align="center">

### ⚽ Built with Go & Gin Framework

**Happy Testing!** 🚀

</div>
