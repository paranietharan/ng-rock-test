# ng-rock-test

A simple Go HTTP server with two endpoints:

## Endpoints

- `/` — Returns a JSON message: `{"message": "Hello, World!"}`
- `/welcome` — Returns a JSON message: `{"message": "Welcome to our API!"}`

## How to Run

1. Make sure you have Go 1.20 or newer installed.
2. Open a terminal in the project directory.
3. Run:
   ```powershell
   go run main.go
   ```
4. The server will start on port 80. Visit [http://localhost](http://localhost) or [http://localhost/welcome](http://localhost/welcome) in your browser.

## Files

- `main.go` — Main application file with HTTP handlers.
- `go.mod` — Go module definition.

