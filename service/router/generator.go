package router

import (
	"net/http"

	"github.com/NghiaTranUIT/artify-core/models"
	"github.com/labstack/echo"
)

func (r *Router) ApplyGeneratorRoute(echo *echo.Echo) {
	g := echo.Group("/generate")
	g.GET("/author_photo", r.generateAuthorPhoto)
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
