package main

import (
	"greeter_api/client"
	"greeter_api/router"
	"log"
)

func main() {
	// Remember to call the Init() function to initialize the go-micro client service
	client.Init()

	// Start Gin Router at port 3000
	r := router.NewRouter()
	if err := r.Run("0.0.0.0:3000"); err != nil {
		log.Print(err.Error())
	}
}