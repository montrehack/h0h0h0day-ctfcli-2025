# Xmaze Gift

Small Go web app: a 10×10 grid with a red cube representing a player controlled by WASD keys.

This README explains how to build and run the project locally and with Docker Compose, and how to pass runtime flags.

## Prerequisites

- Go 1.25+ (for local build)
- Docker and Docker Compose (for containerized run)

## Quick local run

1. From the repository root:
    ```bash
    go mod tidy
    go run main.go
    ```
2. Open: http://localhost:8080

To run a built binary:
```bash
go build -o xmaze-gift main.go
./xmaze-gift
```

The CTF flags are in the `./src/flag.go`

## Controls

- W - move up
- A - move left
- S - move down
- D - move right
- Q - interact
- E - cancel/confirm

## Challenge rules

This repository is part of a web challenge. Do not perform brute-force requests against the API and do not rely on steganography. Respect rate limits and responsible testing procedures.

## Docker Compose (recommended)

Add the Dockerfile and docker-compose.yml below to the project root (examples included). Then build and run:

```bash
docker compose up --build -d
docker compose logs -f
docker compose down
```

Open: http://localhost:25032
