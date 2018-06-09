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

func (r *Resource) CreatePhotoByWikiart(param models.PhotoParam, author models.Author) (*models.Photo, error) {

	// Search first
	photo := models.Photo{}

	// Search photo by name firstly
	err := r.PostgreSQL.Where("name = ?", param.Name).First(&photo)

	// Create new photo
	if err != nil {
		newPhoto := models.Photo{
			Name:           param.Name,
			ImageUrl:       param.Url,
			Location:       param.Location,
			Dimensions:     "",
			Media:          param.Media,
			Style:          param.Style,
			Date:           param.Date,
			Info:           param.Info,
			Width:          uint(param.Width),
			Height:         uint(param.Height),
			AuthorID:       author.ID,
			OriginalSource: param.OriginalSource,
		}
		err = r.PostgreSQL.Eager("Author").Create(&newPhoto)
		if err != nil {
			return nil, err
		}
		return &newPhoto, nil
	}

	// Skip
	return &photo, nil
}
