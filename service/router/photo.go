package router

import (
	"github.com/NghiaTranUIT/artify-core/utils"
	"github.com/labstack/echo"
	"net/http"
)

func (r *Router) ApplyPhotoRoute(e *echo.Echo) {
	g := e.Group("/feature")
	g.GET("/today", r.GetFeatureToday)
}

func (r *Router) GetFeatureToday(c echo.Context) error {

	// Latest from DB
	photo, err := r.R.GetLatestFeaturePhoto()

	// Not found
	if photo == nil {
		response := utils.ResponseError(err)
		return c.JSON(http.StatusOK, response)
	}

	// Success
	response := utils.ResponseSuccess(photo)
	return c.JSON(http.StatusOK, response)
}
