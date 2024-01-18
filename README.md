# Simple Go Web Service for Logging POST Requests

This is a simple Go web service that listens for POST requests on a specified port, logs each request to a separate file with a timestamp, and prints brief information about each incoming request to the standard output.

## Running locally

Clone the repository and go to the project directory. Run the server:

```bash
go run http-bulk-receiver.go -port 8080
```

Replace `8080` with your desired port number.

In a separate terminal session, make a POST request using `curl`:

```bash
curl -X POST -d "Hello, this is the request body" http://localhost:8080/log
```

Adjust the URL and request body as needed.

## Features

- Listens for POST requests on a specified port.
- Logs each request to a separate file with a timestamp in the filename.
- Checks received payload is JSON and reports on contents.
- Can simulate server round-trip time (slow responses).
- Prints brief information about each incoming request to the standard output.

## Command-Line Arguments

View full usage with `./http-bulk-receiver -help`.

```
Usage of ./http-bulk-receiver:
  -path string
        Path for logging requests (default "/log")
  -port string
        Port number for the server (default "8080")
  -sleep int
        Sleep duration in milliseconds before responding with 200 OK
  -store
        Store files
```
The `-sleep` setting is useful for testing throughput when the response-time or round-trip time to a remote service is significant.

# Docker Desktop

You can run this under Docker Desktop. This is most easily done using VS Code with the Dev Containers extension.
Once the container has started, open a new terminal session. Build the binary then run it:

```bash
go build
./http-bulk-receiver -port 8080
```

To send messages from another container to this one, you can discover your container's IP address using `ifconfig`. This is usually in the range `172.17.x.x`.