package main

import (
	"gses2.app/api"
	cronjobs "gses2.app/cronJobs"
	"gses2.app/database"
	"gses2.app/mail"
)

const (
	port = "8080"
)

func main() {
	database.ConnectDb()
	mail.SetupMailService()
	cronjobs.SetupCronJobs()

	err := api.StartServer(port)
	if err != nil {
		panic(err)
	}
}
