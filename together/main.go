package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"artweb/handlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}
	http.Handle("/stats/", http.StripPrefix("/stats/", http.FileServer(http.Dir("stats"))))
	// Handle all requests to the root path
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)
	port := ":8000"
	// Start the server on port 8080
	log.Printf("Server started on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
