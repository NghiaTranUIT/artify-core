package router

import (
	"net/http"

	"github.com/NghiaTranUIT/artify-core/models"
	"github.com/labstack/echo"
)

func (r *Router) ApplyHomeRoute(e *echo.Echo) {
	e.GET("/", r.getDefaultHome)
}

func (r *Router) getDefaultHome(c echo.Context) error {
	payload := map[string]string{"Hello": "It's Artify Core"}
	return c.JSON(http.StatusOK, models.NewSuccessReponse(payload))
}
