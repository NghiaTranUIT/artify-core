package main

import (
	"fmt"
	"github.com/artify/constant"
	"github.com/artify/service"
	"net/http"
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
	s := &http.Server{
		Addr: constant.AppHost + ":" + constant.AppPort,
	}

	// Start
	app.Logger.Fatal(app.StartServer(s))
}
