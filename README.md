# Court Case Management System (CCMS)

A web application for digitizing court workflows with role-based access for **Administrator**, **Judge**, **Plaintiff**, and **Accused**. Core features include case registration/assignment, notifications, appeals, search, and reporting. ([GitHub][1])

## Table of Contents

* [Overview](#overview)
* [Architecture](#architecture)
* [Features](#features)
* [Tech Stack](#tech-stack)
* [Getting Started](#getting-started)

  * [Option A: Run with Docker (recommended)](#option-a-run-with-docker-recommended)
  * [Option B: Local Go build (advanced)](#option-b-local-go-build-advanced)
* [Project Layout](#project-layout)
* [Configuration](#configuration)
* [Development](#development)
* [Testing](#testing)
* [Roadmap / Ideas](#roadmap--ideas)
* [License](#license)

---

## Overview

The system replaces manual, paper-based court processes with a unified web app. After authentication, users are routed to role-appropriate dashboards:

* **Admin:** manage judges, register cases (plaintiff/accused), assign judges, post notifications, manage waiting lists, review reports.
* **Judge:** view notifications and assigned cases, set appointments, close cases with decisions, attach witness documents, view progress reports.
* **Plaintiff / Accused:** track case status, submit appeals, view notifications. ([GitHub][1])

---

## Architecture

**Monorepo** with a Go backend and a static JS/HTML UI:

* **Backend:** Go services (Go modules) that implement domain logic (cases, notifications, appeals, search, reports, sessions). The repository is a valid Go module (contains `go.mod`). ([Go Packages][2])
* **Frontend:** Static HTML/JS under `UI/` served by the backend (no separate Node build step observed). ([GitHub][1])
* **Containerization:** Multi-stage Docker build producing a tiny Alpine image that runs the compiled server and exposes port **8181**. ([GitHub][3])

> Runtime port: the Dockerfile exposes **8181**; map this to a host port when running. ([GitHub][3])

---

## Features

* Case lifecycle: register, assign judge, update status
* Role-based notifications
* Appeals workflow
* Search across cases
* Statistics & reports ([GitHub][1])

---

## Tech Stack

* **Language:** Go (backend), JavaScript + HTML (frontend)
* **Repo language mix:** \~**54.8% JS**, **30.9% Go**, **13.2% HTML** (GitHub analysis). ([GitHub][1])
* **Container:** Docker (multi-stage build; final image based on Alpine, `ENTRYPOINT ./app`, `EXPOSE 8181`). ([GitHub][3])

---

## Getting Started

### Prerequisites

* **Docker** 20+ (if you use the Docker path)
* or **Go** 1.18+ (module-aware build), if building locally without Docker (see advanced path) ([Go Packages][2])

> **Note:** The repo includes `docker-compose.yml` files at the root. If you prefer Compose, you can run via Compose as well. (Files are present in the root listing.) ([GitHub][1])

---

### Option A: Run with Docker (recommended)

1. **Clone**

```bash
git clone https://github.com/Surafeljava/Court-Case-Management-System.git
cd Court-Case-Management-System
```

2. **Build**

```bash
docker build -t ccms:latest .
```

> This uses the multi-stage Dockerfile to compile the Go app and copy the binary to a minimal Alpine image. ([GitHub][3])

3. **Run**

```bash
docker run --rm -p 8181:8181 ccms:latest
```

4. **Open the app**
   Visit: `http://localhost:8181`

> If you prefer Compose and your environment has Docker Compose v2+, you can try:

```bash
docker compose up --build
```

(Compose files exist in the repo root.) ([GitHub][1])

---

### Option B: Local Go build (advanced)

> This project uses Go modules; ensure you have a recent Go toolchain installed. ([Go Packages][2])

```bash
# from repo root
go mod download
go install ./...
```

That produces a binary (per the Dockerfile it ends up as `app` in `/go/bin` inside the builder image). Locally, you can run the compiled binary or `go run` the main package if you prefer. If the server uses port 8181 (as in Docker), open `http://localhost:8181` in your browser. ([GitHub][3])

---

## Project Layout

Root folders (top-level) you’ll interact with most: ([GitHub][1])

```
Entity/           # Domain entities (case, user, judge, plaintiff, accused, etc.)
UI/               # Static HTML/JS assets for the web UI
appealUse/        # Appeal use-cases / business logic
caseUse/          # Case registration/assignment/update logic
court/            # HTTP handlers / routing / server (and tests)
form/             # Form DTOs / request parsing helpers
notificationUse/  # Notification workflows
reportUse/        # Reporting workflows / stats
rtoken/           # Token helpers (e.g., reset/registration tokens, etc.)
searchUse/        # Search use-cases / indexing helpers
session/          # Session/auth helpers (login/logout, role checks)
Dockerfile
docker-compose.yml
docker-compose.debug.yml
README.md
```

> The names above reflect the folder list in the repo root; internal package boundaries follow the “use-case” naming convention (`*Use`). ([GitHub][1])

---

## Configuration

The app runs with sensible defaults in Docker (no env required in the Dockerfile). If you need configuration:

* **Port:** container exposes **8181** (map to the host with `-p 8181:8181`). ([GitHub][3])
* **Environment variables:** none are explicitly referenced in the Dockerfile; check the `court/` package for server settings (e.g., address/port) if you want to customize. ([GitHub][3])

> If your deployment requires persistence (DB/files), add the corresponding env vars/volumes to `docker run` or `docker-compose.yml`. (The repo lists Compose files; customize as needed.) ([GitHub][1])

---

## Development

* **Hot reload:** Not configured by default. Run with `go run` or rebuild the Docker image when you change server code.
* **Frontend:** The `UI/` directory contains static HTML/JS; edit and refresh your browser to see changes. ([GitHub][1])
* **Code organization:** Business logic is grouped by “use-case” directories (`caseUse`, `appealUse`, etc.), with domain structures in `Entity/`. HTTP handlers/routing live under `court/`. ([GitHub][1])

---

## Testing

The `court/` directory includes Go test files (e.g., handler tests). To run tests:

```bash
go test ./...
```

(There’s at least one test file under `court/` visible from the listing.) ([GitHub][4])

[1]: https://github.com/Surafeljava/Court-Case-Management-System "GitHub - Surafeljava/Court-Case-Management-System"
[2]: https://pkg.go.dev/github.com/surafeljava/court-case-management-system?utm_source=chatgpt.com "court-case-management-system module - github.com/surafeljava/court-case ..."
[3]: https://github.com/Surafeljava/Court-Case-Management-System/blob/master/Dockerfile?utm_source=chatgpt.com "Court-Case-Management-System/Dockerfile at master - GitHub"
[4]: https://github.com/Surafeljava/Court-Case-Management-System/blob/master/court/admin_case_handler_test.go?utm_source=chatgpt.com "Court-Case-Management-System/court/admin_case_handler_test.go at master ..."
[5]: https://github.com/Surafeljava/Court-Case-Management-System/releases?utm_source=chatgpt.com "Releases: Surafeljava/Court-Case-Management-System - GitHub"
