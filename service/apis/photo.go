package apis

import (
	"github.com/NghiaTranUIT/artify-core/utils"
	"github.com/labstack/echo"
	"net/http"
)

func ApplyPhotoRoute(e *echo.Echo) {
	g := e.Group("/feature")
	g.GET("/today", getFeatureToday)
}

func getFeatureToday(c echo.Context) error {
	data := map[string]string{"Hello": "It's Photo feature for today"}
	response := utils.ResponseSuccess(data)
	return c.JSON(http.StatusOK, response)
}
