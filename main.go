package main

import (
	"log"
	"middleware-mmksi/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	route := os.Getenv("MIDDLEWARE_SERVER")
	server.NewServer(route)
}
