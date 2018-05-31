package resources

import (
	"github.com/NghiaTranUIT/artify-core/models"
)

func (r *Resource) CreateNewSampleAuthorAndPhoto() (*models.Author, error) {
	photo := models.Photo{
		Name:     "The Starry Night",
		Height:   3959,
		Width:    5000,
		ImageUrl: "https://uploads3.wikiart.org/00142/images/vincent-van-gogh/the-starry-night.jpg",
	}
	author := models.Author{
		Born:        "30 March 1853; Zundert, Netherlands",
		Died:        "29 July 1890; Auvers-sur-Oise, France",
		Name:        "Vincent van Gogh",
		Nationality: "Dutch",
		Photos:      []models.Photo{photo},
		Wikipedia:   "en.wikipedia.org/wiki/Vincent_van_Gogh",
	}
	err := r.PostgreSQL.Create(&author).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}
