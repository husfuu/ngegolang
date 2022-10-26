package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

// ingat ServeHTTP adalah handler
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	// cuman ngeresponse "Hello World" doang di setiap route
	// karena handler tersebut ngehandle semua request
	server.ListenAndServe()
}
