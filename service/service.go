package service

import (
	"github.com/artify/constant"
	"github.com/artify/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func GetEngine(config constant.Config) *echo.Echo {

	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Gzip())

	// Apply Routes

	// Default route
	app.GET("/", home)

	// Return
	return app
}

func home(e echo.Context) error {
	res := model.Response{
		Code:    http.StatusOK,
		Data:    map[string]string{"Hello": "It's Artify Core"},
		Message: "Success",
	}
	return e.JSON(http.StatusOK, res)
}
