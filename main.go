package main

import (
	"fmt"
	"github.com/artify/constant"
	"github.com/artify/resources"
	"github.com/artify/service"
	"net/http"
)

func main() {

	// Default configuration
	config := constant.Config{
		IsEnablePostgreSQL:       true,
		IsEnablePostgreSQLLogger: true,
	}

	// Resource
	r, err := resources.Init(config)
	if err != nil {
		fmt.Println("[ERROR] Initialed Resource ...", err)
	}
	defer r.Close()

	// Hello
	fmt.Printf("Hello, world.\n")

	// App
	app := service.GetEngine(config)
	s := &http.Server{
		Addr: constant.AppHost + ":" + constant.AppPort,
	}

	// Start
	app.Logger.Fatal(app.StartServer(s))
}
