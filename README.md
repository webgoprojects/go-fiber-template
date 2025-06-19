# go-fiber-template

A simple web server template using [Go Fiber](https://gofiber.io/).

## Getting Started

### Prerequisites
- Go 1.22+

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/linesmerrill/go-fiber-template.git
   cd go-fiber-template
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

### Running the Server

Start the server with:
```sh
go run main.go
```

The server will start on [http://localhost:3000](http://localhost:3000).

### Available Endpoints

- `GET /` — Returns `Hello, World!`
- `GET /health` — Returns `{ "status": "ok" }` (health check)