package main

import (
	"net/http"
	"os"
	"todo-app/pkg/backend"
)

func main() {
	server := backend.CreateServer()
	server.RegisterRoutes()
	err := http.ListenAndServe(":8080", server.Router)
	if err != nil {
		os.Exit(1)
	}
}
