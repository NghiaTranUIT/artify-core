package router

import (
	"github.com/NghiaTranUIT/artify-core/model"
	"github.com/labstack/echo"
	"net/http"
)

func (r *Router) ApplyHomeRoute(e *echo.Echo) {
	e.GET("/", r.getDefaultHome)
}

func (r *Router) getDefaultHome(c echo.Context) error {
	payload := map[string]string{"Hello": "It's Artify Core"}
	return c.JSON(http.StatusOK, model.NewSuccessReponse(payload))
}
