# 🎬 Govies

[![Go Report Card](https://goreportcard.com/badge/github.com/heisenberg8055/govies)](https://goreportcard.com/report/github.com/heisenberg8055/govies)

Govies is a RESTful API for managing movie data, user authentication, and permissions. It is built with Go and PostgreSQL, providing a robust and scalable solution for movie-related applications.

---

## 📚 Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [API Documentation](#api-documentation)
  - [Movies Endpoints](#movies-endpoints)
  - [User Endpoints](#user-endpoints)
  - [Token Endpoints](#token-endpoints)
  - [Health Check](#health-check)
- [Testing](#testing)
- [Docker](#docker)
- [Metrics](#metrics)
- [Security](#security)
- [License](#license)
- [Contributing](#contributing)
- [Contact](#contact)

---

## 🚀 Features

- **Movie Management**: CRUD operations for movies.
- **User Authentication**: Token-based authentication for secure access.
- **Permissions System**: Role-based access control for users.
- **Rate Limiting**: Prevent abuse with configurable rate limits.
- **CORS Support**: Enable cross-origin requests for trusted origins.
- **Health Check**: Monitor the API's availability.
- **Metrics**: Track API performance and usage statistics.

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Libraries**:
  - `github.com/julienschmidt/httprouter` - Lightweight HTTP router.
  - `github.com/lib/pq` - PostgreSQL driver.
  - `github.com/wneessen/go-mail` - Email sending library.
  - `github.com/joho/godotenv` - Environment variable management.

---

## 📂 Project Structure

```
govies/
├── cmd/                  # Main application entrypoint
│   └── govies/           # API server code
├── internal/             # Internal packages
│   ├── data/             # Data models and database logic
│   ├── mailer/           # Email sending logic
│   ├── realip/           # Real IP extraction
│   ├── validator/        # Input validation
│   └── vcs/              # Version info
├── migrations/           # SQL migrations for database setup
├── Makefile              # Build and test commands
├── Dockerfile            # API server Docker build
├── db.Dockerfile         # Database Docker build
├── docker-compose.yml    # Multi-service orchestration
├── .env                  # Environment variables
├── LICENSE               # License file
└── README.md             # This file
```

---

## 🏁 Getting Started

### Prerequisites

- Go 1.25+
- Docker & Docker Compose

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/heisenberg8055/govies.git
   cd govies
   ```

2. **Configure environment variables:**
   Create a `.env` file:
   ```
   POSTGRES_USER=youruser
   POSTGRES_PASSWORD=yourpassword
   POSTGRES_DB=govies
   POSTGRES_HOST=localhost
   ```

3. **Start services with Docker Compose:**
   ```sh
   docker compose up
   ```

4. **Run the API server locally:**
   ```sh
   make run
   ```

---

## 📖 API Documentation

### Movies Endpoints

- `GET /v1/movies` — List movies (supports filtering, sorting, pagination)
- `POST /v1/movies` — Create a new movie
- `GET /v1/movies/:id` — Get details of a movie
- `PATCH /v1/movies/:id` — Update a movie
- `DELETE /v1/movies/:id` — Delete a movie

### User Endpoints

- `POST /v1/users` — Register a new user
- `PUT /v1/users/activated` — Activate user account
- `PUT /v1/users/password` — Reset user password

### Token Endpoints

- `POST /v1/tokens/authentication` — Get authentication token
- `POST /v1/tokens/password-reset` — Request password reset token
- `POST /v1/tokens/activation` — Request activation token

### Health Check

- `GET /v1/healthz` — Check API status

---

## 🧪 Testing

- Run all tests:
  ```sh
  make test
  ```
- View coverage report:
  ```sh
  make coverage
  ```

---

## 🐳 Docker

- Build and run all services:
  ```sh
  docker compose up
  ```
- Build only:
  ```sh
  make docker-build
  ```

---

## 📊 Metrics

- `GET /v1/metrics` — Exposes runtime metrics for monitoring.

---

## 🔒 Security

- Rate limiting is enabled by default.
- CORS is configurable via trusted origins.
- Passwords are hashed using bcrypt.

---

## 📄 License

This project is licensed under the Apache 2.0 License. See [LICENSE](./LICENSE).

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check [issues page](https://github.com/heisenberg8055/govies/issues).

---

## 📬 Contact

Created by [@heisenberg8055](https://github.com/heisenberg8055) — feel free to reach out!

---
