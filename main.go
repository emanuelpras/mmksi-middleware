package main

import (
	"middleware-mmksi/server"
	"middleware-mmksi/util"
	"os"
)

func main() {
	route := os.Getenv("MIDDLEWARE_SERVER")
	config := util.LoadConfig()
	dbInit, err := util.MySQL(config)

	if err != nil {
		panic(err)
	}

	server := server.NewServer(dbInit, route)
	server.ListenAndServe("8080")
}
