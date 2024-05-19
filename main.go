package main

import (
	"gses2.app/api"
	"gses2.app/database"
	"gses2.app/mail"
)

const (
	port = "8080"
)

func main() {
	database.ConnectDb()
	mail.SetupMailService()

	err := api.StartServer(port)
	if err != nil {
		panic(err)
	}
}
