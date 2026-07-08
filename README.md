# Helpdesk Ticket System API

A lightweight, high-performance RESTful API for managing support tickets, built using **Go (Golang)**, the **Gin Web Framework**, and secured with **JWT Authentication**. 

The system leverages a completely thread-safe, in-memory repository layer decoupled from business execution, implementing clean architecture boundaries.

---

## 🌐 Live Production Deployment

The API is containerized, pushed to a Docker registry, and actively hosted live on Render.

* **Base Production URL**: `https://ticket-system-api-v1-0.onrender.com`
* **Live Health Check**: [https://ticket-system-api-v1-0.onrender.com/health](https://ticket-system-api-v1-0.onrender.com/health)

> **Note on Render Free Tier**: If the app hasn't received traffic recently, Render spins down the container. The first incoming request might take up to 50 seconds to complete while the instance spins back up.

---

## 🚀 Key Features

* **User Management & Security**: User registration and login utilizing state-of-the-art **Bcrypt** cryptographic password hashing.
* **JWT State Synchronization**: Stateless session tokens containing encrypted credentials (`user_id`) to validate route permissions.
* **Ticket Lifecycle Engine**: Complete management suite enforcing explicit ownership and state-transition rules (`open` -> `in_progress` -> `closed`).
* **Concurrency Safe**: Thread-safe database emulation layer utilizing `sync.RWMutex` to block race conditions during high-volume reads/writes.
* **Containerized Architecture**: Multi-stage dockerized environment optimized for minimal deployment footprints.

---

## 🛠️ Technology Stack

* **Backend**: Go (v1.26+)
* **HTTP Routing & Middleware**: Gin Web Framework
* **Security & Encryption**: golang-jwt/jwt/v5, x/crypto/bcrypt
* **Deployment & DevOps**: Docker, Docker Registry, Render Cloud Platform

---

## 🏃 Getting Started & Local Run

### Prerequisites
* Go installed locally OR Docker Desktop.

### Option 1: Running Locally
1. Clone the repository and install dependency paths:
   ```bash
   go mod tidy


# ticket-system

![GitHub stars](https://img.shields.io/github/stars/Hardik585/ticket-system?style=for-the-badge&logo=github) ![GitHub forks](https://img.shields.io/github/forks/Hardik585/ticket-system?style=for-the-badge&logo=github) ![GitHub issues](https://img.shields.io/github/issues/Hardik585/ticket-system?style=for-the-badge&logo=github) ![Last commit](https://img.shields.io/github/last-commit/Hardik585/ticket-system?style=for-the-badge&logo=github)

## 📑 Table of Contents

- [Description](#description)
- [Tech Stack](#tech-stack)
- [Quick Start](#quick-start)
- [Key Dependencies](#key-dependencies)
- [Available Scripts](#available-scripts)
- [Project Structure](#project-structure)
- [Development Setup](#development-setup)
- [Deployment](#deployment)
- [Contributors](#contributors)
- [Contributing](#contributing)

## 📝 Description

ticket-system — a software project built with Docker, Go.

## 🛠️ Tech Stack

![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white) ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

## ⚡ Quick Start

```bash

# 1. Clone the repository
git clone https://github.com/Hardik585/ticket-system.git

# Run the project
go run .
```

## 📦 Key Dependencies

```
github.com/gin-gonic/gin: v1.12.0
github.com/golang-jwt/jwt/v5: v5.3.1
golang.org/x/crypto: v0.48.0
github.com/bytedance/gopkg: v0.1.3
github.com/bytedance/sonic: v1.15.0
github.com/bytedance/sonic/loader: v0.5.0
github.com/cloudwego/base64x: v0.1.6
github.com/gabriel-vasile/mimetype: v1.4.12
github.com/gin-contrib/sse: v1.1.0
github.com/go-playground/locales: v0.14.1
github.com/go-playground/universal-translator: v0.18.1
github.com/go-playground/validator/v10: v10.30.1
github.com/goccy/go-json: v0.10.5
github.com/goccy/go-yaml: v1.19.2
github.com/json-iterator/go: v1.1.12
```

## 🚀 Available Scripts

- **run** — `go run .`
- **build** — `go build`
- **test** — `go test ./...`

## 📁 Project Structure

```
.
├── Dockerfile
├── cmd
│   └── api
│       └── main.go
├── go.mod
├── go.sum
└── internal
    ├── auth
    │   ├── handler.go
    │   └── service.go
    ├── models
    │   └── models.go
    ├── repository
    │   └── memory.go
    └── ticket
        ├── handler.go
        └── service.go
```

## 🛠️ Development Setup

### Docker
1. `docker build -t my-app .`
2. `docker run -p 3000:3000 my-app`

### Go
1. Install Go (v1.21+ recommended)
2. `go mod download && go run .`

## 🚢 Deployment

### Docker
```bash
docker build -t ticket-system .
docker run -p 3000:3000 ticket-system
```

## 👥 Contributors

Thanks to everyone who has contributed to this project:

<p align="left">
<a href="https://github.com/Hardik585" title="Hardik585"><img src="https://avatars.githubusercontent.com/u/102424028?v=4&s=64" width="64" height="64" alt="Hardik585" style="border-radius:50%" /></a>
</p>

[See the full list of contributors →](https://github.com/Hardik585/ticket-system/graphs/contributors)

## 👥 Contributing

Contributions are welcome! Here's the standard flow:

1. **Fork** the repository
2. **Clone** your fork: `git clone https://github.com/Hardik585/ticket-system.git`
3. **Branch**: `git checkout -b feature/your-feature`
4. **Commit**: `git commit -m 'feat: add some feature'`
5. **Push**: `git push origin feature/your-feature`
6. **Open** a pull request

Please follow the existing code style and include tests for new behavior where applicable.
