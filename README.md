# Simple Go Web Service for Logging POST Requests

This is a simple Go web service that listens for POST requests on a specified port, logs each request to a separate file with a timestamp, and prints brief information about each incoming request to the standard output.

## Usage

Clone the repository:
```bash
git clone https://github.com/yourusername/golang-web-service.git
```

Navigate to the project directory:

```bash
cd golang-web-service
```

Run the server:

```bash
go run main.go -port 8080
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
go run main.go -port 8080
```

