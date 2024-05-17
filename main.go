package main

import (
	"gses2.app/api"
	"gses2.app/database"
)

const (
	port = "8080"
)

func main() {
	database.ConnectDb()

	err := api.StartServer(port)
	if err != nil {
		panic(err)
	}
}
