package resources

import (
	"github.com/NghiaTranUIT/artify-core/models"
)

func (r *Resource) GetLatestFeaturePhoto() (*models.Photo, error) {

	// Get feature row for TODAY
	feature, err := r.GetFeatureOnDashboard()
	if err != nil {
		return nil, err
	}

	// Get author
	photo := models.Photo{}
	err = r.PostgreSQL.Eager("Author").Where("id = ?", feature.Photo.ID).First(&photo)
	if err != nil {
		return nil, err
	}

	// Has feature for today
	return &photo, nil
}
