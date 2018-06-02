package resources

import (
	"github.com/NghiaTranUIT/artify-core/models"
	"time"
)

func (r *Resource) GetFeatureOnDashboard() (*models.Dashboard, error) {
	feature := models.Dashboard{}

	// Get first one
	todayStr := time.Now().Format("2006-01-02")
	err := r.PostgreSQL.Eager("Photo").
		Where("DATE(updated_at) = ? AND type = ?", todayStr, "feature").
		First(&feature)

	// Don't have feature for today
	if err != nil {
		err = r.PostgreSQL.Eager("Photo").
			Where("type = ?", "feature").
			First(&feature)

		// Dont' have any feature row
		if err != nil {
			newFeature, _ := r.initializeFeatureOnDashboard()
			return newFeature, nil
		}

		// Old feature (yesterday)
		// Get next photo
		lastPhotoCreatedAt := feature.Photo.CreatedAt.Format("2006-01-02 15:04:05.000000000")
		nextPhoto := models.Photo{}
		err = r.PostgreSQL.
			Where("created_at > ?", lastPhotoCreatedAt).
			Order("created_at ASC").
			Limit(1).
			First(&nextPhoto)

		// Run out of photo
		if err != nil {
			firstPhoto := models.Photo{}
			err = r.PostgreSQL.First(&firstPhoto)

			// Update the first again
			feature.Photo = firstPhoto
			feature.PhotoID = firstPhoto.ID
			r.PostgreSQL.Eager("Photo").Update(&feature)
			return &feature, nil
		}

		// Update
		feature.Photo = nextPhoto
		feature.PhotoID = nextPhoto.ID
		r.PostgreSQL.Eager("Photo").Update(&feature)
		return &feature, nil
	}

	// Has feature photo for TODAY
	return &feature, nil
}

func (r *Resource) initializeFeatureOnDashboard() (*models.Dashboard, error) {

	// Get first photo
	firstPhoto := models.Photo{}
	err := r.PostgreSQL.First(&firstPhoto)

	// There is no photos
	if err != nil {
		return nil, err
	}

	// Create first Dashboard
	dashboard := models.Dashboard{
		Type:  "feature",
		Photo: firstPhoto,
	}
	err = r.PostgreSQL.Eager("Photo").Create(&dashboard)
	if err != nil {
		return nil, err
	}
	return &dashboard, nil
}
