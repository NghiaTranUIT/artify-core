package router

import (
	"net/http"

	"github.com/NghiaTranUIT/artify-core/models"
	"github.com/labstack/echo"
)

func (r *Router) ApplyGeneratorRoute(echo *echo.Echo) {
	g := echo.Group("/generate")
	g.GET("/feature", r.generateFeaturePhoto)
}

func (r *Router) generateFeaturePhoto(c echo.Context) error {

	// Generate Feature if need
	feature, err := r.R.GetLatestFeatureOnDashboard()
	if err != nil {
		feature, err = r.R.GenerateSampleFeaturePhoto()
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.NewErrorResponse(err))
		}
	}

	// Success
	return c.JSON(http.StatusOK, models.NewSuccessReponse(feature))
}
