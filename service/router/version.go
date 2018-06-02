package router

import (
	"net/http"

	"github.com/NghiaTranUIT/artify-core/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func (r *Router) ApplyVersionRoute(e *echo.Echo) {
	g := e.Group("/version")
	g.GET("/latest", r.getLatestAppVersion)
	g.GET("/update", r.getUpdateVersion)
}

func (r *Router) getLatestAppVersion(c echo.Context) error {

	// Latest from DB
	version, err := r.R.GetLatestAppVersion()

	// Not found
	if version == nil {
		return c.JSON(http.StatusNoContent, models.NewErrorResponse(err))
	}

	// Success
	return c.JSON(http.StatusOK, models.NewSuccessReponse(version))
}

func (r *Router) getUpdateVersion(c echo.Context) error {
	buildVersion := c.QueryParam("build_version")

	// Latest from DB
	version, _ := r.R.GetUpdateAppVersion(buildVersion)

	// Not found
	if version == nil {
		return c.JSON(http.StatusNoContent, models.NewErrorResponse(errors.New("There is no latest build")))
	}

	// Success
	return c.JSON(http.StatusOK, models.NewSuccessSquirrelReponse(version))
}
