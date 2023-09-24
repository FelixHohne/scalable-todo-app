package main

import "todo-app/pkg/backend"

func main() {
	server := backend.CreateServer()
	server.RegisterRoutes()
}
