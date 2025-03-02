# go-stack-yourself

**go-stack-yourself** is a **Proof of Concept (PoC)** project demonstrating how to integrate **OpenTelemetry** in a *
*Golang** application and use the **OpenTelemetry Collector** to route logs, traces, and metrics to various
observability backends, including **Loki**, **Tempo**, and **Prometheus**.

---

## Features

- **OpenTelemetry Integration**: Instrumentation for Golang applications using the OpenTelemetry SDK.
- **Observability Backends**:
    - **Loki**: For centralized logging.
    - **Tempo**: For distributed tracing.
    - **Prometheus**: For metrics collection and monitoring.
- **Easy Setup**: Quickly spin up the entire stack with `docker`.

---

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go (Golang)](https://go.dev/) (optional, for direct development)

---

## Setup Instructions

Follow the steps below to get the project up and running:

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-stack-yourself.git
cd go-stack-yourself
```

### 2. Start the Observability Stack

Run the following command to start all required components (OpenTelemetry Collector, Loki, Tempo, Prometheus, etc.):

```bash
docker compose up -d
```

This will run all containers defined in the `docker-compose.yml` file in detached mode.

### 3. Generate Go code from the templ file

Run the templ generate command.

```bash
docker compose exec app sh -c "templ generate"
```

### 4. Run the Application

To run the Golang application with OpenTelemetry instrumentation, use the following command:

```bash
docker compose exec app sh -c "go run ./cmd/dice"
```

### Running Tests in the Application Container

To run all the tests located in the `tests/` folder using Docker Compose, you can execute the following command:

```bash
docker compose exec app sh -c "go test ./tests/..."
```

This command runs the `go test` tool inside the `app` container, targeting all test files (`*_test.go`) located within
the `tests/` folder and its subdirectories. It ensures that tests are executed in the same environment as your
application container, matching the dependencies and configurations defined in the container's setup.

### 5. Application

#### Test Page
http://localhost:8081/app/

The application will send logs, metrics, and traces to the respective backends.

---

#### 6. Monitoring

- [Grafana](http://localhost:8081/grafana)
- [Prometheus](http://localhost:8081/prometheus)