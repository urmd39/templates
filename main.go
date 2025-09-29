package main

import (
	infrastructure "templates/infrastructure"
	"templates/server"
)

func main() {
	server.HandleHTTPServer(infrastructure.APIHostport)
}
