package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	restHandler "github.com/ngavinsir/golangtraining/internal/rest"
)

func main() {
	m := http.NewServeMux()

	restHandler.InitHelloHandler(m)

	port := ":5050"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = fmt.Sprintf(":%s", envPort)
	}

	s := &http.Server{
		Addr:    port,
		Handler: m,
	}

	log.Printf("Server started on %s", port)
	log.Fatal(s.ListenAndServe())
}
