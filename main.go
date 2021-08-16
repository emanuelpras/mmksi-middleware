package main

import (
	"middleware-mmksi/server"
	"os"
)

func main() {
	route := os.Getenv("MIDDLEWARE_SERVER")
	server.NewServer(route)
}
