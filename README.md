# HTTP/1.1 (WebSocket) Echo Server
---
A simple echo server for educational purposes.

The design focuses on delivering a reliable WebSocket echo mechanism without introducing unnecessary complexity.

## Behavior
---
Any messages sent from a WebSocket client are echoed back as a WebSocket message.

## Configuration
---
- The `PORT` environment variable sets the server port, which defaults to 8080.

# Running the server
---
## Instalattion
`git clone https://github.com/xddbom/echo-server.git`

## Build the application (if necessary)
`go build -o echo-server ./cmd/echo-server`

## Starting Server Using WebSocket CLI (e.g., `wscat`)
There is no Docker start option, so you'll need to run it locally.
`wscat -c ws://localhost:8080/ws`







 
