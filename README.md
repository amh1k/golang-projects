# ♠️ Poker Player Server

[![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A robust, test-driven backend application built in Go for managing poker games, tracking player scores, and maintaining a competitive league.

This project implements a flexible architecture with a core **domain logic**, multiple **user interfaces** (REST API and CLI), and interchangeable **data stores**. It serves as an excellent reference implementation of Test-Driven Development (TDD), Dependency Injection, and SOLID principles in Go.

---

## ✨ Features

- **Multi-Interface Support**: Interact with the system via a RESTful HTTP Server or an interactive Command Line Interface.
- **RESTful API**: Endpoints to record wins, retrieve individual player scores, and fetch the full sorted league leaderboard.
- **Persistent Storage**: Uses a JSON-backed File System Store (`game.db.json`) for persistence, ensuring no data is lost between restarts.
- **Interchangeable Storage**: Defines a `PlayerStore` interface, making it trivial to swap the JSON store with an in-memory store or an external database (e.g., PostgreSQL).
- **Test-Driven Design**: The entire codebase is thoroughly covered by unit and integration tests, utilizing `httptest`, mock stores, and dependency injection.

---

## 🏗️ Project Structure

```text
├── CLI.go                     # Command-line interface logic
├── server.go                  # HTTP Server routing and handlers
├── file_ststem_store.go       # Persistent JSON file storage implementation
├── in_memory_player_store.go  # In-memory transient storage mapping
├── league.go                  # League domain models and logic
├── tape.go                    # Helper for safe file truncation and rewriting
├── testing.go                 # Test stubs and assertion helpers
├── cmd/
│   ├── cli/
│   │   └── main.go            # Entry point for the interactive CLI application
│   └── webserver/
│       └── main.go            # Entry point for the HTTP Web Server
└── *test.go                   # Comprehensive unit and integration test coverage
```

---

## 🚀 Getting Started

### Prerequisites
- [Go](https://go.dev/dl/) 1.22 or higher installed.

### 1. Web Server

Start the REST API server:
```bash
go run ./cmd/webserver/main.go
```
The server will start listening on port `:5000` and automatically create/use `game.db.json` in the project root.

### 2. Command Line Interface (CLI)

Run the CLI tool to track games interactively:
```bash
go run ./cmd/cli/main.go
```
Once started, simply input `{PlayerName} wins` to tally up the scores! Example:
```text
Let's play poker
Type {Name} wins to record a win
Chris wins
Cleo wins
```

---

## 📡 API Endpoints

When running the web server on `http://localhost:5000`:

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/players/{name}` | Returns the total wins (score) for the specified player. Returns `404` if the player does not exist. |
| `POST`| `/players/{name}` | Records a win for the specified player. Returns `202 Accepted`. |
| `GET` | `/league` | Returns a JSON array of all players sorted by their scores in descending order. |

### API Examples

**Record a Win**
```bash
curl -X POST http://localhost:5000/players/Chris
```

**Get Player Score**
```bash
curl http://localhost:5000/players/Chris
# Response: 1
```

**Get League Leaderboard**
```bash
curl http://localhost:5000/league
# Response: [{"Name":"Chris","Wins":1}]
```

---

## 🧪 Testing

The library heavily leverages Go's standard `testing` package along with parallel executions and table-driven tests.

To run the full test suite:
```bash
go test -v ./...
```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!
1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
