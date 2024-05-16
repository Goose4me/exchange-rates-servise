package main

import (
	"gses2.app/api"
)

const (
	port = "8080"
)

func main() {

	err := api.StartServer(port)
	if err != nil {
		panic(err)
	}
}
