package main

import (
	"github.com/NghiaTranUIT/artify-core/constant"
	"github.com/NghiaTranUIT/artify-core/resources"
	"github.com/NghiaTranUIT/artify-core/service"
	"github.com/NghiaTranUIT/artify-core/utils"
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

	// Migration
	r.MigrateDatabase()

	// App
	app := service.New(r)
	app.Start()
}
