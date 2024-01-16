# Simple Go Web Service for Logging POST Requests

This is a simple Go web service that listens for POST requests on a specified port, logs each request to a separate file with a timestamp, and prints brief information about each incoming request to the standard output.

## Usage

Clone the repository and go to the project directory. Run the server:

```bash
go run http-bulk-rx.go -port 8080
```

Replace `8080` with your desired port number.

Make a POST request using curl:

```bash
curl -X POST -d "Hello, this is the request body" http://localhost:8080/log
```

Adjust the URL and request body as needed.

## Features

- Listens for POST requests on a specified port.
- Logs each request to a separate file with a timestamp in the filename.
- Prints brief information about each incoming request to the standard output.

## Command-Line Arguments

- `-port`: Specifies the port number for the server. Default is `8080`.

## Example

```bash
go run http-bulk-rx.go -port 8080
```

