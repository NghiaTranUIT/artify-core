package router

import (
	"github.com/NghiaTranUIT/artify-core/model"
	"github.com/NghiaTranUIT/artify-core/resources"
	"github.com/labstack/echo"
	"net/http"
)

type Router struct {
	R *resources.Resource
}

func (r *Router) ApplyRoutes(e *echo.Echo) {
	r.ApplyDefaultRoute(e)
	r.ApplyPhotoRoute(e)
	r.ApplyAuthorRoute(e)
	r.ApplyGeneratorRoute(e)
}

func (r *Router) ApplyDefaultRoute(e *echo.Echo) {
	e.GET("/", r.Home)
}

func (r *Router) Home(c echo.Context) error {
	res := model.Response{
		Code:    http.StatusOK,
		Data:    map[string]string{"Hello": "It's Artify Core"},
		Message: "Success",
	}
	return c.JSON(http.StatusOK, res)
}
