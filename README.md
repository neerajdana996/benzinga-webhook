
# Golang Benzinga Webhook

This is a simple Golang-based webhook application that demonstrates key concepts such as:

- **Gin framework** for HTTP routing
- **Docker** for containerization
- **Docker Compose** for managing services
- **Environment variables** loaded from `.env` file
- **Structured logging** using `logrus`
- **Batch processing** with retry logic
- **Unit testing** with `testify` and mocking HTTP requests

## Project Structure

```
benzinga-webhook/
├── controllers/             # HTTP handlers (Gin controllers)
│   ├── log_controller.go
│   └── log_controller_test.go
├── models/                  # Data models
│   └── payload.go
├── services/                # Business logic (batch processing, caching)
│   ├── log_service.go
│   └── log_service_test.go
├── .env                     # Environment variables
├── Dockerfile               # Dockerfile for building the application
├── docker-compose.yml        # Docker Compose file
├── main.go                  # Main application entry point
└── README.md                # This README file
```

## Features

1. **Health Check Endpoint:**
   - `GET /healthz` - Returns `200 OK` with a simple "OK" response.

2. **Log Endpoint:**
   - `POST /log` - Accepts a JSON payload, stores it in an in-memory cache, and processes it in batches.

3. **Batch Processing:**
   - Collects payloads in memory and sends them in batches based on batch size or interval.
   - Retry logic (3 retries with 2-second delays) if the batch fails to send.

4. **Logging:**
   - Uses `logrus` for structured, leveled logging.
   - Logs important events such as application startup, batch sending, and retries.

5. **Environment Variables:**
   - The batch size, batch interval, and POST endpoint are configurable through environment variables loaded from a `.env` file.

6. **Docker and Docker Compose:**
   - The application is containerized using Docker.
   - Docker Compose is used to manage the app's configuration and run tests in containers.

7. **Unit Testing:**
   - Unit tests are written for both controllers and services using Go's `testing` package and `testify` for assertions.

## Getting Started

### Prerequisites

- **Docker** and **Docker Compose** installed on your system.
- **Go 1.20** or later if you want to run locally.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/neerajdana996/benzinga-webhook.git
   cd benzinga-webhook
   ```

2. Create a `.env` file to define environment variables:
   ```bash
   BATCH_SIZE=10
   BATCH_INTERVAL=15s
   POST_ENDPOINT=http://requestbin.net/r/your-endpoint
   ```

3. Build and run the application using Docker Compose:
   ```bash
   docker-compose up --build
   ```

The application will be available at `http://localhost:8080`.

### API Endpoints

- **Health Check**:
  ```bash
  GET /healthz
  ```
  Response:
  ```json
  "OK"
  ```

- **Log Payload**:
  ```bash
  POST /log
  ```
  Example JSON Payload:
  ```json
  {
    "user_id": 1,
    "total": 1.65,
    "title": "example title",
    "meta": {
      "logins": [
        {
          "time": "2020-08-08T01:52:50Z",
          "ip": "0.0.0.0"
        }
      ],
      "phone_numbers": {
        "home": "555-1212",
        "mobile": "123-5555"
      }
    },
    "completed": false
  }
  ```

### Running Unit Tests

1. Run unit tests in the service and controller packages:

```bash
go test ./... -v
```

### Environment Variables

- **BATCH_SIZE**: Defines the number of logs to collect before sending a batch.
- **BATCH_INTERVAL**: Time interval between batch sends.
- **POST_ENDPOINT**: The external endpoint where the batch is sent.

