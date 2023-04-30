package main

import (
	"log"
	"os"
	"students/pkg/app"
)

func main() {
	a := app.App{}
	a.Initialize()

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8081"
	}

	log.Println("Run port:", httpPort)

	a.Run(httpPort)
}
