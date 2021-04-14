package main

import (
	"log"
	"net/http"
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
    http.Handle("/hello-world", http.HandlerFunc(helloWorld))
    http.Handle("/health", http.HandlerFunc(health))
    err := http.ListenAndServe(":5050", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}