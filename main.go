package main

import (
	"github.com/artify/constant"
	"github.com/artify/resources"
	"github.com/artify/service"
	"github.com/artify/utils"
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
		utils.LogError("Initialed Resource ...", err)
		return
	}
	defer r.Close()

	// Hello
	utils.LogInfo("Hello, world.")

	// App
	app := service.GetEngine(config)
	s := &http.Server{
		Addr: constant.AppHost + ":" + constant.AppPort,
	}

	// Start
	app.Logger.Fatal(app.StartServer(s))
}
