package router

import (
	"net/http"

	"github.com/NghiaTranUIT/artify-core/models"
	"github.com/NghiaTranUIT/artify-core/utils"
	"github.com/labstack/echo"
)

func (r *Router) ApplyGeneratorRoute(echo *echo.Echo) {
	g := echo.Group("/generate")
	g.GET("/all", r.generateAll)
	g.GET("/author_photo", r.generateAuthorPhoto)
	g.GET("/version", r.generateSampleVersion)
	g.GET("/feature", r.generateFeaturePhoto)
}

func (r *Router) generateAll(c echo.Context) error {
	// Author photo
	author, err := r.R.CreateNewSampleAuthorAndPhoto()

	// Generate Version if need
	_, err = r.R.GetLatestAppVersion()
	if err != nil {
		_, err = r.R.GenerateSampleVersion()
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.NewErrorResponse(err))
		}
	}

	// Generate Feature if need
	feature, err := r.R.GetLatestFeatureOnDashboard()
	if err != nil {
		utils.LogWarning("Feature = nil => Create new same feature")
		feature, err = r.R.GenerateSampleFeaturePhoto()
		utils.LogWarning("New feature = ", feature)
		utils.LogWarning("Error = ", err)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.NewErrorResponse(err))
		}
	}

	// Success
	return c.JSON(http.StatusOK, models.NewSuccessReponse(author))
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
