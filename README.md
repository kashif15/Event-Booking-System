# 🎟 Event Booking Backend API

> A **production-ready backend system** built with **Go (Golang)** that demonstrates real-world backend engineering: authentication, authorization, transactions, pagination, search, and testing.

---

## 🚀 About This Project

This project focuses on:

- Security
- Scalability
- Real business rules
- Clean architecture
- Production practices
---

## ✨ Features at a Glance

### 🔐 Authentication & Security
- JWT authentication
- Refresh token flow (session-based auth)
- Secure password hashing (bcrypt)
- Logout & token revocation
- Role-based access control (USER / ADMIN)

### 📅 Event Management
- Create & manage events
- Ownership-based authorization
- Pagination & filtering
- Search by title & location

### 🎟 Booking System
- Secure event booking
- Duplicate booking prevention
- Capacity enforcement
- Transaction-safe database operations
- View & cancel bookings

### 🧪 Testing & Quality
- Unit tests (auth, JWT)
- API tests using `httptest`
- Middleware tests
- Clean, maintainable codebase

---


## 🏗 Architecture Overview

```text
cmd/server
 └── main.go            # Application entrypoint

internal/
 ├── auth/              # Auth, JWT, refresh tokens
 ├── event/             # Event domain logic
 ├── booking/           # Booking domain logic
 ├── middleware/        # Auth middleware
 ├── routes/            # HTTP routing

pkg/
 ├── config/            # Environment configuration
 └── database/          # Database connection

migrations/             # SQL migrations
```

## 🧠 Architecture Principles

- Separation of concerns
- Thin HTTP handlers
- Repository pattern
- Middleware-driven security
- Transaction-safe business logic

---

## 🛠 Tech Stack

| Category | Technology |
|--------|------------|
| Language | Go (Golang) |
| Framework | Gin |
| Database | PostgreSQL |
| Auth | JWT + Refresh Tokens |
| Security | bcrypt |
| Testing | testing, httptest |
| Config | godotenv |

---

## 📦 Getting Started

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/kashif15/event-booking-api.git
cd event-booking-api
```

### 2️⃣ Setup Environment Variables
Create a .env file:
```bash
APP_PORT=8080
DB_URL=postgres://user:password@localhost:5432/event_booking?sslmode=disable
JWT_SECRET=your_secret_key
```
### 3️⃣ Run Database Migrations
```bash
psql -d event_booking -f migrations/001_create_users.sql
psql -d event_booking -f migrations/002_create_events.sql
psql -d event_booking -f migrations/003_create_bookings.sql
psql -d event_booking -f migrations/004_create_refresh_tokens.sql
```
### 4️⃣ Install Dependencies
```bash
go mod tidy
```
### 5️⃣ Start the Server
```bash
go run cmd/server/main.go
```


## 🚀 Server

The server runs at:

```text
http://localhost:8080
```
## 🔑 API Endpoints

🔐 Authentication
```text
Method	Endpoint
POST	/auth/register
POST	/auth/login
POST	/auth/refresh
POST	/auth/logout
```
📅 Events (Protected)
```text
Method	Endpoint
POST	/events
GET	/events
GET	/events/:id
DELETE	/events/:id
```
🎟 Bookings (Protected)
```text
Method	Endpoint
POST	/events/:id/book
DELETE	/events/:id/book
GET	/bookings
```
🔎 Pagination, Filtering & Search
```text
Example Request
GET /events?page=1&limit=10&search=go&created_by=me
```

## Supported Query Parameters

- page
- limit
- search
- status
- created_by
- from_date

## 💡 Engineering Highlights

- Secure session-based authentication
- Stateless access tokens with DB-backed refresh tokens
- Transaction-safe booking logic
- Ownership-based authorization
- Production-grade API structure

## 📌 Future Enhancements

- Token rotation
- Rate limiting
- Redis caching
- Admin analytics APIs
- Docker & CI/CD pipeline

## 👨‍💻 Author

Kashif Ahmad

Backend Developer (Go)
- GitHub: https://github.com/kashif15
- Email: kashifahmad599@gmail.com












