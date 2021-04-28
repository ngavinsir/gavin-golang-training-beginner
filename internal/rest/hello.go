package rest

import "net/http"

func InitHelloHandler(m *http.ServeMux) {
	m.HandleFunc("/hello-world", helloWorld)
	m.HandleFunc("/health", health)
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world"}`))
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"healthy"}`))
}
