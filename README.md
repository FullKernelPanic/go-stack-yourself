# Go stack yourself

This project is a Proof of Concept (PoC) aimed at demonstrating how to implement OpenTelemetry in a Go-based application. It integrates logs, traces, and metrics collection using Loki, Tempo, and Prometheus, visualized through Grafana. The backend is built using Go and `templ`, while the frontend leverages `htmx` and `scss`.

---

## Tech Stack

The project consists of the following components:

- **Backend**: Written in Go using [templ](https://github.com/a-h/templ).
- **Frontend**: Built with HTML templates, [htmx](https://htmx.org/) for dynamic interactions, and SCSS for styling.
- **Logs**: Collected using [Loki](https://grafana.com/oss/loki/).
- **Traces**: Managed using [Tempo](https://grafana.com/oss/tempo/).
- **Metrics**: Collected using [Prometheus](https://prometheus.io/).
- **Visualization**: Centralized on [Grafana](https://grafana.com/grafana/).

---

## Features

- **Logs**: Aggregated and stored in Loki.
- **Traces**: OpenTelemetry-powered distributed tracing visualized in Grafana via Tempo.
- **Metrics**: Collected from the Go application using Prometheus.
- **Real-time Monitoring**: Visualize logs, traces, and metrics with Grafana's powerful UI.
- **Hot Reloading**: Automatically rebuild the Go application inside the Docker container using [Air](https://github.com/air-verse/air).

---

## Getting Started

### Prerequisites

Before setting up the project, ensure you have the following tools installed:

- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)
- [Make](https://www.gnu.org/software/make/)

---

### Setup Instructions

1. **Clone the Repository**:
   ```bash
   git clone git@github.com:FullKernelPanic/go-stack-yourself.git
   cd go-stack-yourself
   ```

2. **Build the Project**:
   ```bash
   make build
   ```

3. **Start the Stack**:
   ```bash
   make up
   ```

4. **Access the Application**:
  - **Web UI**: [http://localhost:8081/app/](http://localhost:8081/app/)
  - **Grafana Dashboard**: [http://localhost:8081/grafana/](http://localhost:8081/grafana/)

5. **Run Automated Tests**:
   ```bash
   make run-tests
   ```

---

## Development Workflow

The project is pre-configured to use `Air` inside the Docker container, which automatically monitors file changes and rebuilds the Go application. This ensures efficient development without requiring additional installation on the host machine.

---

## Architecture Overview

1. **Backend**:
  - Written in Go.
  - Uses `templ` for HTML templates generation.

2. **Frontend**:
  - Dynamically updates the UI using `htmx`.
  - Styled with SCSS.

3. **Instrumentation**:
  - OpenTelemetry is configured to enable logs, traces, and metrics collection.

4. **Monitoring Components**:
  - Logs are stored in Loki.
  - Distributed tracing is managed with Tempo.
  - Metrics data is gathered by Prometheus.
  - Grafana visualizes all the collected data for real-time monitoring.

---

## Contributing

Contributions to improve this Proof of Concept are welcome! Feel free to submit issues or pull requests to help enhance the project.

---

## License

This project is open-sourced under an appropriate LICENSE (update this section with a license type).

---

## Acknowledgements

Special thanks to:

- The OpenTelemetry community for enabling effortless observability.
- The Go ecosystem for an excellent development experience.
- Grafana, Loki, Tempo, and Prometheus for their robust monitoring tools.

---

## Contact

If you have any questions or feedback, feel free to open an issue or reach out!

---

Happy Monitoring!
