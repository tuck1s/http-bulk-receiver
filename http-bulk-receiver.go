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
	path := flag.String("path", "/log", "Path for logging requests")
	sleepDuration := flag.Int("sleep", 0, "Sleep duration in milliseconds before responding with 200 OK")
	flag.Parse()

	http.HandleFunc(*path, func(w http.ResponseWriter, r *http.Request) {
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

		// Sleep for the specified duration before responding with 200 OK
		time.Sleep(time.Duration(*sleepDuration) * time.Millisecond)

		fmt.Fprint(w, "Request logged successfully")
	})

	fmt.Printf("Server is listening on port %s, path: %s, sleep duration: %d milliseconds...\n", *port, *path, *sleepDuration)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
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
