package resources

import (
	"github.com/NghiaTranUIT/artify-core/model"
)

func (r *Resource) GetLatestFeaturePhoto() (*model.Photo, error) {
	lastPhoto := model.Photo{}
	err := r.PostgreSQL.Last(&lastPhoto).Related(&lastPhoto.Author).Error

	if lastPhoto.ID == 0 {
		return nil, err
	}

	return &lastPhoto, nil
}
