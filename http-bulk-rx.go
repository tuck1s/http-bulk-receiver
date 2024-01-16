package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := flag.String("port", "8080", "Port number for the server")
	flag.Parse()
	path := "/log"
	http.HandleFunc(path, logHandler)
	fmt.Printf("Server is listening on port %s... Send POST requests to path %s\n", *port, path)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request from %s\n", r.Method, r.RemoteAddr)

	if r.Method != http.MethodPost {
		fmt.Printf("Invalid request method: %s\n", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Log the request to a file with a timestamp in the filename
	err = logRequest(body)
	if err != nil {
		http.Error(w, "Error logging request", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "OK") // send "200 OK"
}

func logRequest(body []byte) error {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("request_%s.log", timestamp)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(string(body))

	return nil
}
