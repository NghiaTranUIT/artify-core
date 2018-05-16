package resources

import (
	"github.com/NghiaTranUIT/artify-core/model"
	"github.com/NghiaTranUIT/artify-core/utils"
	"github.com/labstack/echo"
	"net/http"
)

func (r *Resource) CreateSampleAuthorAndPhoto(c echo.Context) error {
	photo := model.Photo{
		Name:     "The Starry Night",
		Height:   3959,
		Width:    5000,
		ImageUrl: "https://uploads3.wikiart.org/00142/images/vincent-van-gogh/the-starry-night.jpg",
	}
	author := model.Author{
		Born:        "30 March 1853; Zundert, Netherlands",
		Died:        "29 July 1890; Auvers-sur-Oise, France",
		Name:        "Vincent van Gogh",
		Nationality: "Dutch",
		Photos:      []model.Photo{photo},
		Wikipedia:   "en.wikipedia.org/wiki/Vincent_van_Gogh",
	}
	r.PostgreSQL.Create(&author)
	response := utils.ResponseSuccess(author)
	return c.JSON(http.StatusOK, response)
}
