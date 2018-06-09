package router

import (
	"github.com/NghiaTranUIT/artify-core/models"
	"github.com/labstack/echo"
	"net/http"
)

func (r *Router) ApplySpiderRoute(e *echo.Echo) {
	g := e.Group("/spider")
	g.POST("/wikiart", r.createPhotoAuthorFromWikiArt)
}

func (r *Router) createPhotoAuthorFromWikiArt(c echo.Context) error {
	param := models.WikiartParam{}

	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusNoContent, models.NewErrorResponse(err))
	}

	// Create photo and author if need
	photo, err := r.R.CreatePhotoAuthorFromWikiart(param)
	if err != nil {
		return c.JSON(http.StatusNoContent, models.NewErrorResponse(err))
	}

	// Success
	return c.JSON(http.StatusOK, models.NewSuccessReponse(photo))
}
