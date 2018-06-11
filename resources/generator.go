package resources

import (
	"github.com/NghiaTranUIT/artify-core/models"
)

func (r *Resource) GenerateSampleFeaturePhoto() (*models.Dashboard, error) {

	// Get first image
	firstPhoto := models.Photo{}
	err := r.PostgreSQL.First(&firstPhoto)
	if err != nil {
		return nil, err
	}

	feature := models.Dashboard{
		Type:    "feature",
		PhotoID: firstPhoto.ID,
	}
	err = r.PostgreSQL.Create(&feature)
	if err != nil {
		return nil, err
	}
	return &feature, nil
}
