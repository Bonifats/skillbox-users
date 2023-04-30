package main

import (
	"fmt"
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

	fmt.Println("Run port:", httpPort)

	a.Run(httpPort)
}
