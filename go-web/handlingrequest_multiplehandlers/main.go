package main

import (
	"fmt"
	"net/http"
)

// handler 1
type HelloHandler struct {
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

// handler 2
type WorldHandler struct {
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "World!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}
