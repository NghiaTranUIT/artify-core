package router

import (
	"net/http"

	"github.com/NghiaTranUIT/artify-core/models"
	"github.com/labstack/echo"
)

func (r *Router) ApplyPhotoRoute(e *echo.Echo) {
	g := e.Group("/feature")
	g.GET("/today", r.getFeatureToday)
}

func (r *Router) getFeatureToday(c echo.Context) error {

	// Latest from DB
	photo, err := r.R.GetLatestFeaturePhoto()

	// Not found
	if photo == nil {
		return c.JSON(http.StatusOK, models.NewErrorResponse(err))
	}

	// Success
	return c.JSON(http.StatusOK, models.NewSuccessReponse(photo))
}
