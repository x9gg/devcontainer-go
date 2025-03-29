package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {
	// Set the port to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Define routes
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/version", handleVersion)

	// Start the server
	currentTime := time.Now().Format("15:04:05")
	fmt.Printf("[%s] Server starting on port %s...\n", currentTime, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// handleRoot is a simple handler that returns a Hello World message
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! Welcome to this Go web application.")
}

// handleVersion returns information about the Go version running
func handleVersion(w http.ResponseWriter, r *http.Request) {
	goVersion := runtime.Version()
	goOS := runtime.GOOS
	goArch := runtime.GOARCH

	response := fmt.Sprintf("Go Version: %s\nOperating System: %s\nArchitecture: %s",
		goVersion, goOS, goArch)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, response)
}
