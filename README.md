# devops_agent

[![Go](https://img.shields.io/badge/language-Go-00ADD8)](https://golang.org)
[![Python](https://img.shields.io/badge/language-Python-3776AB)](https://www.python.org)
[![License](https://img.shields.io/badge/license-MIT-blue)](LICENSE)

DevOps Agent is a lightweight, extensible agent for running automation, collectors, and orchestration tasks on remote servers. The core is implemented in Go for performance and portability, with Python used for plugins and helper scripts and Shell for small operational utilities.


Language composition
--------------------
- Go: 77.2%
- Python: 19.9%
- Shell: 2.9%

What devops_agent does
----------------------
- Runs as a small background agent on servers to execute tasks, collect metrics, and run user-defined plugins.
- Exposes a local API/CLI for scheduling jobs and retrieving results.
- Supports long-running workers and asynchronous job queues.
- Extensible plugin system: write plugins in Python for quick prototyping or Go for high-performance needs.

High-level architecture
-----------------------
                       +-----------------+
                       |   Management    |
                       |  (CLI / API)    |
                       +--------+--------+
                                |
                                | HTTP / gRPC / Local socket
                                v
   +-----------+    +-----------+-----------+    +-----------+
   |  Storage  |<-->|   Agent Core (Go)    |<-->|  Workers  |
   | (SQLite)  |    | - API, Auth, Queue   |    | (Python/Go)|
   +-----------+    +----------------------+    +-----------+
                                |
                                v
                         +--------------+
                         |  Monitoring  |
                         |  Logging/Apm |
                         +--------------+

How it works
------------
- The Agent Core (Go) boots and loads configuration from `config.yaml`.
- Plugins (Python or Go) are discovered in `plugins/` and registered.
- The agent exposes a local CLI and optionally a management API for remote control.
- Jobs are scheduled or triggered; long-running jobs are executed in worker processes.
- Results, logs, and metrics are stored in local storage and optionally forwarded to a remote server.

Core components
---------------
- cmd/agent: main agent process (Go).
- internal/*: core libraries (queue, scheduler, plugin loader).
- plugins/: Python plugin examples and helper scripts.
- scripts/: shell utilities for installation and demo.
- docs/: architecture and operational docs.

Quickstart — build and run
--------------------------
Prerequisites:
- Go 1.20+ (for building the agent)
- Python 3.8+ (for Python plugins)
- Make or GNU tools (optional)
- Docker (optional: run agent in container)

Clone and build from source:

```bash
git clone https://github.com/starcnobucks/devops_agent.git
cd devops_agent
# Build the agent (Go)
go build -o bin/devops_agent ./cmd/agent
# Create a Python venv for plugins (optional)
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

Run the agent locally with the example configuration:

```bash
cp config.example.yaml config.yaml
./bin/devops_agent --config config.yaml
```

Run in Docker (example):

```bash
docker build -t devops_agent:latest .
docker run -v $(pwd)/config.yaml:/app/config.yaml -v $(pwd)/plugins:/app/plugins devops_agent:latest
```

Configuration
-------------
Example minimal `config.yaml`:

```yaml
agent:
  listen_addr: "127.0.0.1:8080"
  data_dir: "/var/lib/devops_agent"
  log_level: "info"

plugins:
  path: "./plugins"
  python:
    enabled: true

storage:
  type: "sqlite"
  path: "./data/agent.db"
```

Plugins
-------
- Python plugins: place a Python package/module in `plugins/python/`. Each plugin must implement a `register()` function that the agent loader will call.
- Go plugins: compile as part of the agent or use the plugin interface defined in `internal/plugin`.

Examples
--------
Schedule a job using the local CLI:

```bash
./bin/devops_agent job run --name "collect-disk" --command "plugins/python/disk_collector.py"
```

Call the local HTTP API to list jobs:

```bash
curl -s http://127.0.0.1:8080/v1/jobs | jq
```

Development and testing
-----------------------
Go tests:

```bash
go test ./... -v
```

Python tests (if any):

```bash
source .venv/bin/activate
pytest -q
```

Linting:

```bash
golangci-lint run
flake8 plugins
```

Logging and monitoring
----------------------
- The agent logs structured JSON to stdout by default; configure log level in `config.yaml`.
- Metrics can be exported via Prometheus if enabled in config.
- Integrations for external logging/APM services can be added via plugins.

Security considerations
-----------------------
- The agent runs user-specified code as configured; run with least privilege and use containerization for untrusted plugins.
- Secure the management API with TLS and authentication in production.
- Use OS-level process isolation for untrusted workloads.

Packaging and deployment
------------------------
- Systemd unit file example located in `contrib/systemd/devops_agent.service`.
- Container image: Dockerfile builds the static Go binary and copies Python plugin code.
- For fleet deployment, use your configuration management (Ansible/Chef/Puppet) or a container orchestrator.

Release and versioning
----------------------
- Follow semantic versioning: MAJOR.MINOR.PATCH.
- Tag releases in Git and build release artifacts for each supported platform using `goreleaser` or similar.

Contributing
------------
- Fork, create a feature branch, run tests, and open a PR.
- Add tests for new features and update documentation.
- For large changes, open an issue first to discuss design.

Troubleshooting
---------------
- Agent fails to start: check `config.yaml` path and permissions.
- Plugins not discovered: ensure plugin files are executable and in the configured `plugins.path`.
- Long job durations: configure worker timeouts and resource limits.

Acknowledgements
----------------
This project uses a small Go runtime for the core agent and leverages Python for rapid plugin development and operational scripts.

License
-------
MIT License. See LICENSE file for details.
