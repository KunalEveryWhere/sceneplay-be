# ScenePlay Backend (sceneplay-be)

## Overview

ScenePlay Backend is a secure and scalable **Go API** that facilitates **donation payments** via **Razorpay Payment Links**. It follows a clean, modular structure for maintainability and extensibility.

## Features

- 🛡️ **Secure Payment Integration** – Uses Razorpay's API to generate payment links dynamically.
- 🌍 **International Support** – Accepts donations in multiple currencies (including USD & INR).
- 🚀 **Modular Architecture** – Clean project structure with separation of concerns.
- 🏗️ **Scalable & Maintainable** – Built with best practices for future growth.
- 🐳 **Docker Support** – Easily run the backend in a containerized environment.

## Project Structure

```sh
sceneplay-be/
│── cmd/
│   └── server/           # Main entry point (main.go)
│── internal/
│   ├── config/           # Environment and config loader
│   ├── handlers/         # API handlers (controllers)
│   ├── models/           # Request & Response structures
│   ├── services/         # Razorpay API logic
│   └── routes/           # Route definitions
│── pkg/
│   └── utils/            # Utility functions (logging, formatting)
│── .env                  # Environment variables
│── Dockerfile            # Docker container configuration
│── docker-compose.yml    # Docker Compose configuration
│── go.mod                # Go module file
│── go.sum                # Dependency versions
│── README.md             # Documentation
```

## Getting Started

### 1️⃣ Prerequisites

- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
- Razorpay Account ([Sign Up](https://razorpay.com/))
- Environment variables set in `.env` file

### 2️⃣ Installation

Clone the repository:

```sh
git clone https://github.com/kunaleverywhere/sceneplay-be.git
cd sceneplay-be
```

### 3️⃣ Configuration

Create a `.env` file in the root directory and add your Razorpay credentials:

```ini
## ServerConfig
SCENEPLAY_SERVER_PORT = 9000
ALLOWED_ORIGINS = https://www.sceneplay.com

## RazorPay Configs
RAZORPAY_KEY_ID = key_id
RAZORPAY_KEY_SECRET = secret_key
RAZORPAY_CALLBACK_URL = https://www.sceneplay.com/payment-success

## Developer Mode
DEVELOPER_MODE = disabled
```

> Note: When `DEVELOPER_MODE` is `enabled`, then requests from all origins are entertained.

### 4️⃣ Running with Docker

Build and start the application using Docker Compose:

```sh
docker-compose up --build -d
```

This will:

- Build the Go application inside a container.
- Expose the application on **port** as per the config file.

To stop the application, run:

```sh
docker-compose down
```

## Running Locally (Without Docker)

If you prefer to run the server without Docker:

```sh
go run cmd/server/main.go
```

The API will be available at `http://localhost:{port}`.

## API Endpoints

### 🩺 Health Check

- **Endpoint:** `GET /health`
- **Description:** Verifies if the server is running.
- **Response:**
  ```json
  {
    "status": "ok"
  }
  ```

### 💳 Create a Payment Link

- **Endpoint:** `POST /create-payment`
- **Description:** Generates a Razorpay Payment Link.
- **Request Body:**
  ```json
  {
    "amount": 1000,
    "currency": "USD"
  }
  ```
- **Response:**
  ```json
  {
    "amount": 1000,
    "created-at": "2024-07-26T12:00:00Z",
    "currency": "USD",
    "payment-link": "https://rzp.io/l/examplelink",
    "expire-by": "2024-07-30T12:00:00Z"
  }
  ```
