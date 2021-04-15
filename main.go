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
    http.HandleFunc("/hello-world", helloWorld)
    http.HandleFunc("/health", health)
    err := http.ListenAndServe(":5050", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}