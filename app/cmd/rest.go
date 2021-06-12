package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	restHandler "github.com/ngavinsir/golangtraining/internal/rest"
)

func Execute() {
	r := httprouter.New()

	restHandler.InitHelloHandler(r)
	restHandler.InitPaymentCodesHandler(r, paymentCodesService)

	port := ":5050"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = fmt.Sprintf(":%s", envPort)
	}

	log.Printf("Server started on %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
