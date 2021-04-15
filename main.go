package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)


func helloWorld(w http.ResponseWriter, req *http.Request) {	
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world"}`))
}

func health(w http.ResponseWriter, req *http.Request) {	
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"healthy"}`))
}

func main() {
    http.HandleFunc("/hello-world", helloWorld)
    http.HandleFunc("/health", health)
    
	port := ":5050"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = fmt.Sprintf(":%s", envPort)
	}

	log.Printf("Server started on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}