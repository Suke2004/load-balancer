# Smart Adaptive Load Balancer

A high-performance, intelligent load balancer built in **Go (Golang)**. This project is designed to efficiently distribute incoming network traffic across multiple backend servers, ensuring high availability and reliability. It features an adaptive algorithm that monitors server health and load to make smart routing decisions, coupled with a real-time dashboard for monitoring.

## 🚀 Features

* **⚡ High Performance**: Built with Go's lightweight concurrency model (Goroutines) to handle high throughput with low latency.
* **🧠 Adaptive Balancing**: Routes traffic based on real-time server metrics (response time, active connections) rather than just static round-robin.
* **💓 Active Health Checks**: Periodically probes backend servers to ensure they are alive. Unhealthy servers are automatically removed from the pool; healthy ones are reintegrated.
* **📊 Real-time Dashboard**: Includes a web-based dashboard (located in `dashboard/`) to visualize cluster status, request rates, and server health.
* **🛡️ Fault Tolerance**: Graceful error handling and retry mechanisms.

## 📂 Project Structure

```bash
.
├── cmd/
│   └── balancer/    # Application entry point (main.go)
├── dashboard/       # Frontend code for the monitoring dashboard
├── internal/        # Core application logic (algorithms, health checks, stats)
├── go.mod           # Go module definitions
└── go.sum           # Checksums for dependencies

```

## 🛠️ Prerequisites

Before running the load balancer, ensure you have the following installed:

* **[Go](https://golang.org/dl/)** (Version 1.20 or higher)
* **Git**

## 📥 Installation

1. **Clone the repository:**
```bash
git clone https://github.com/Suke2004/load-balancer.git
cd load-balancer

```


2. **Install dependencies:**
```bash
go mod download

```



## ⚙️ Configuration

*Note: If the project uses a configuration file (e.g., `config.json` or `config.yaml`), ensure it is created in the root directory or inside `cmd/balancer`. A typical configuration might look like this:*

```json
{
  "port": 8000,
  "backends": [
    "http://localhost:8081",
    "http://localhost:8082",
    "http://localhost:8083"
  ],
  "health_check_interval": "10s",
  "strategy": "adaptive"
}

```

## 🏃 Usage

To start the load balancer:

```bash
go run cmd/balancer/main.go

```

Once running, the load balancer will start listening on the configured port (default is usually `8080` or `8000`).

### Accessing the Dashboard

The dashboard frontend is located in the `dashboard/` folder. Depending on the implementation, it may be served automatically by the Go binary or require a separate start.

If it requires a separate build (e.g., if it's a React/Vue app), navigate to the folder and follow the frontend instructions. Otherwise, visit:

`http://localhost:<PORT>/dashboard` (or the specific URL printed in your console).

## 🧪 Testing

To run the internal unit tests:

```bash
go test ./internal/...

```

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

---

<p align="center">
Built with ❤️ by <a href="[https://github.com/Suke2004](https://www.google.com/search?q=https://github.com/Suke2004)">Suke2004</a>
</p>
