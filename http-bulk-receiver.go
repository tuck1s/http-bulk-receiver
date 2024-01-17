package main

import (
	"encoding/json"
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

	log.SetOutput(os.Stdout)

	http.HandleFunc(*path, func(w http.ResponseWriter, r *http.Request) {
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
		err = storeRequestBody(body)
		if err != nil {
			http.Error(w, "Error logging request", http.StatusInternalServerError)
			return
		}

		var requestBody interface{}
		jsonInfo := ""
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			jsonInfo = err.Error()
		} else {
			// Check if the unmarshaled value is an array
			if arr, ok := requestBody.([]interface{}); ok {
				// It's an array, print its length
				jsonInfo = fmt.Sprintf("array length: %d", len(arr))
			} else {
				// It's not an array
				fmt.Println("not an array.")
			}
		}
		log.Printf("Received %s request from %s with body size %d bytes, JSON %s", r.Method, r.RemoteAddr, len(body), jsonInfo)

		// Sleep for the specified duration before responding with 200 OK
		time.Sleep(time.Duration(*sleepDuration) * time.Millisecond)

		fmt.Fprint(w, "OK")
	})

	fmt.Printf("Server is listening on port %s, path: %s, sleep duration: %d milliseconds...\n", *port, *path, *sleepDuration)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func storeRequestBody(body []byte) error {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("request_%s.json", timestamp) // incoming requests to have JSON payload

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(body)
	return err
}
