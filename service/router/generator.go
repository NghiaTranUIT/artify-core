package router

import (
	"github.com/NghiaTranUIT/artify-core/model"
	"github.com/labstack/echo"
	"net/http"
)

func (r *Router) ApplyGeneratorRoute(echo *echo.Echo) {
	g := echo.Group("/generate")
	g.GET("/author_photo", r.generateAuthorPhoto)
}

func (r *Router) generateAuthorPhoto(c echo.Context) error {

	// Latest from DB
	author, err := r.R.CreateNewSampleAuthorAndPhoto()

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(err))
	}

	// Success
	return c.JSON(http.StatusOK, model.NewSuccessReponse(author))
}
