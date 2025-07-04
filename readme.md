
# 🎓 Student API - Built with GoLang, Docker & Modern Tools

A robust, containerized RESTful API built in Go to manage student records with clean architecture, efficient routing, and scalable design patterns. This project demonstrates a production-grade setup using Go, Docker, and essential third-party libraries.

---

## 🚀 Features

- 🔁 Full CRUD for student data
- 🧰 Modular Go architecture (handlers, services, repositories)
- 📦 Dependency management using Go Modules
- 🐳 Dockerized for containerized deployment
- 🔐 JWT-based authentication (optional)
- 📄 Swagger/OpenAPI integration for API docs
- 🔧 Configurable via environment variables
- 🧪 Unit and integration test-ready

---

## 📁 Project Structure

```

student-api/
├── cmd/                # Entry point
│   └── main.go
├── internal/
│   ├── handler/        # HTTP Handlers
│   ├── service/        # Business Logic
│   ├── repository/     # Database Interactions
│   └── model/          # Data Models
├── config/             # Environment Configuration
├── docker/             # Docker-related files
├── docs/               # Swagger/OpenAPI Specs
├── go.mod / go.sum     # Dependency Files
├── Dockerfile
└── README.md

````

---

## 📦 Tech Stack

| Category      | Technology            |
| ------------- | --------------------- |
| Language      | Go (Golang)           |
| Database      | PostgreSQL / MySQL    |
| API Docs      | Swagger (Go-Swagger)  |
| Dependency    | Go Modules            |
| HTTP Router   | Chi / Gin             |
| Auth (opt.)   | JWT (github.com/golang-jwt/jwt) |
| Container     | Docker                |
| Config        | Viper / godotenv      |
| Logging       | logrus / zerolog      |

---

## ⚙️ Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/student-api-go.git
cd student-api-go
````

### 2. Set Up Environment

Create a `.env` file:

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=studentdb
JWT_SECRET=your_jwt_secret
```

### 3. Run with Docker

```bash
docker-compose up --build
```

> This will spin up both the API and the database containers.

---

## 📮 API Endpoints

| Method | Endpoint            | Description            |
| ------ | ------------------- | ---------------------- |
| GET    | `/api/students`     | Get all students       |
| GET    | `/api/students/:id` | Get a specific student |
| POST   | `/api/students`     | Create a new student   |
| PUT    | `/api/students/:id` | Update student details |
| DELETE | `/api/students/:id` | Delete a student       |

---

## 🔐 Authentication (Optional)

To enable JWT:

* Pass the token as `Authorization: Bearer <token>` header.
* You can configure this in middleware.

---

## 🧪 Testing

You can run unit tests using:

```bash
go test ./...
```

Integration tests can be written using `httptest` and run in CI pipelines.

---

## 📚 API Documentation

Swagger is served at:

```
http://localhost:8080/swagger/index.html
```

You can regenerate the docs using:

```bash
swag init
```

---

## 📦 Docker Compose Structure

```yaml
version: '3'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - .env
    depends_on:
      - db
  db:
    image: postgres:13
    environment:
      POSTGRES_DB: studentdb
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: yourpassword
    ports:
      - "5432:5432"
```

---

## 🧠 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

## ✨ Acknowledgements

* [GoLang](https://golang.org/)
* [Chi](https://github.com/go-chi/chi) / [Gin](https://github.com/gin-gonic/gin)
* [Docker](https://www.docker.com/)
* [Swagger](https://swagger.io/)
* [Logrus](https://github.com/sirupsen/logrus)

---

> ⚡ Built with love by \Vedant Pisal – striving for clean, scalable Go code.

