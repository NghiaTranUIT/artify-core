package resources

import (
	"github.com/NghiaTranUIT/artify-core/model"
	"github.com/NghiaTranUIT/artify-core/utils"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
)

func (r *Resource) GetFeatureToday(c echo.Context) error {
	lastPhoto := model.Photo{}
	r.PostgreSQL.Last(&lastPhoto)

	if lastPhoto.ID == 0 {
		response := utils.ResponseError(errors.New("There is no photo"))
		return c.JSON(http.StatusBadRequest, response)
	}

	response := utils.ResponseSuccess(lastPhoto)
	return c.JSON(http.StatusOK, response)
}
