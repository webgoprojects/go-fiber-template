![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white) ![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB) ![macOS](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge&logo=macos&logoColor=F0F0F0)

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