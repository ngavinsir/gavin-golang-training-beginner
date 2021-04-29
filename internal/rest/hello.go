package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InitHelloHandler(r *httprouter.Router) {
	r.HandlerFunc("GET", "/hello-world", helloWorld)
	r.HandlerFunc("GET", "/health", health)
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world"}`))
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"healthy"}`))
}
