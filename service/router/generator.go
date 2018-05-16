package router

import (
	"github.com/NghiaTranUIT/artify-core/utils"
	"github.com/labstack/echo"
	"net/http"
)

func (r *Router) ApplyGeneratorRoute(echo *echo.Echo) {
	g := echo.Group("/generate")
	g.GET("/author_photo", r.GenerateAuthorPhoto)
}

func (r *Router) GenerateAuthorPhoto(c echo.Context) error {

	// Latest from DB
	photo, err := r.R.GetLatestFeaturePhoto()

	// Not found
	if photo == nil {
		response := utils.ResponseError(err)
		return c.JSON(http.StatusBadRequest, response)
	}

	// Success
	response := utils.ResponseSuccess(photo)
	return c.JSON(http.StatusOK, response)
}
