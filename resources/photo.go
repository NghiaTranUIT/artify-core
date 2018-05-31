package resources

import (
	"github.com/NghiaTranUIT/artify-core/models"
)

func (r *Resource) GetLatestFeaturePhoto() (*models.Photo, error) {
	lastPhoto := models.Photo{}
	err := r.PostgreSQL.Eager("Author").First(&lastPhoto)

	if err != nil {
		return nil, err
	}

	return &lastPhoto, nil
}
