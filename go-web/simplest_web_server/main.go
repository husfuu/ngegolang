package main

import (
	"net/http"
)

func main() {
	// jalan di port 8080
	// masih not found karena ngak ada handler juga
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	server.ListenAndServe()
}
