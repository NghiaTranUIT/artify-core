package resources

import (
	"github.com/NghiaTranUIT/artify-core/model"
)

func (r *Resource) CreateNewSampleAuthorAndPhoto() (*model.Author, error) {
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
	err := r.PostgreSQL.Create(&author).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}
