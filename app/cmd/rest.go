package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	restHandler "github.com/ngavinsir/golangtraining/internal/rest"
)

func Execute() {
	m := http.NewServeMux()

	restHandler.InitHelloHandler(m)
	restHandler.InitPaymentCodesHandler(m, paymentCodesRepository)

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
