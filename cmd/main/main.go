package main

import (
	"fmt"
	"net/http"
	"os"
	"todo-app/pkg/backend"
)

func main() {
	server := backend.CreateServer()
	server.NoteStore.CreateNote("Introduction Note", []string{})
	server.RegisterRoutes()
	fmt.Printf("Beginning server\n")
	err := http.ListenAndServe(":8080", server.Router)
	if err != nil {
		os.Exit(1)
	}
}
