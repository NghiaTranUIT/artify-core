package main

import (
	"fmt"
	"github.com/artify/constant"
	"github.com/artify/service"
)

func main() {

	// Default configuration
	config := constant.Config{
		PostgreSQL:       true,
		PostgreSQLLogger: true,
	}

	fmt.Printf("Hello, world.\n")

	// App
	app := service.GetEngine(config)
	app.Run(constant.AppHost + ":" + constant.AppPort)
}
