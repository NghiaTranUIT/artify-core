package router

import (
	"net/http"

	"github.com/NghiaTranUIT/artify-core/models"
	"github.com/labstack/echo"
)

func (r *Router) ApplyGeneratorRoute(echo *echo.Echo) {
	g := echo.Group("/generate")
	g.GET("/author_photo", r.generateAuthorPhoto)
	g.GET("/version", r.generateSampleVersion)
	g.GET("/feature", r.generateFeaturePhoto)
}

func (r *Router) generateAuthorPhoto(c echo.Context) error {

	// Latest from DB
	author, err := r.R.CreateNewSampleAuthorAndPhoto()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse(err))
	}

	// Success
	return c.JSON(http.StatusOK, models.NewSuccessReponse(author))
}

func (r *Router) generateSampleVersion(c echo.Context) error {
	// Generate
	version, err := r.R.GenerateSampleVersion()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse(err))
	}

	// Success
	return c.JSON(http.StatusOK, models.NewSuccessReponse(version))
}

func (r *Router) generateFeaturePhoto(c echo.Context) error {
	feature, err := r.R.GenerateSampleFeaturePhoto()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse(err))
	}
	// Success
	return c.JSON(http.StatusOK, models.NewSuccessReponse(feature))
}
