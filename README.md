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