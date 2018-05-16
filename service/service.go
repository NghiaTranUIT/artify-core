package service

import (
	"github.com/NghiaTranUIT/artify-core/constant"
	"github.com/NghiaTranUIT/artify-core/model"
	"github.com/NghiaTranUIT/artify-core/resources"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type Service struct {
	R    *resources.Resource
	echo *echo.Echo
}

func (s *Service) Start() {

	app := echo.New()
	s.echo = app

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Gzip())

	// Apply Routes
	s.applyPhotoRoute()
	s.applyAuthorRoute()

	// Default route
	app.GET("/", home)

	// Config
	config := &http.Server{
		Addr: constant.AppHost + ":" + constant.AppPort,
	}

	// Start
	app.Logger.Fatal(app.StartServer(config))
}

func home(c echo.Context) error {
	res := model.Response{
		Code:    http.StatusOK,
		Data:    map[string]string{"Hello": "It's Artify Core"},
		Message: "Success",
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) applyPhotoRoute() {
	g := s.echo.Group("/feature")
	g.GET("/today", s.R.GetFeatureToday)
}

func (s *Service) applyAuthorRoute() {
	g := s.echo.Group("/author")
	g.GET("/create_sample", s.R.CreateSampleAuthorAndPhoto)
}
