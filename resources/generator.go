package resources

import (
	"github.com/NghiaTranUIT/artify-core/models"
	"strconv"
)

func (r *Resource) CreateNewSampleAuthorAndPhoto() (*models.Photo, error) {
	fileName := "data.csv"
	csvPhotos, err := ReadCSVFile(fileName)
	if err != nil {
		return nil, err
	}

	// Parse CSV Photo to Photo and Author
	for _, csvPhoto := range csvPhotos {

		author := models.Author{}
		err = r.PostgreSQL.Where("name = ?", csvPhoto.AuthorName).First(&author)

		width, _ := strconv.Atoi(csvPhoto.Width)
		height, _ := strconv.Atoi(csvPhoto.Height)

		photo := models.Photo{
			Name:       csvPhoto.Name,
			ImageUrl:   csvPhoto.ImageURL,
			Width:      uint(width),
			Height:     uint(height),
			Info:       csvPhoto.Info,
			Date:       csvPhoto.Date,
			Style:      csvPhoto.Style,
			Dimensions: csvPhoto.Dimensions,
			Media:      csvPhoto.Media,
			Location:   csvPhoto.Location,
		}

		// Don't have Author yet
		if err != nil {
			author = models.Author{
				Name:        csvPhoto.AuthorName,
				Wikipedia:   csvPhoto.Width,
				Born:        csvPhoto.Born,
				Died:        csvPhoto.Died,
				Nationality: csvPhoto.Nationality,
			}
			photo.Author = author
		} else {
			photo.AuthorID = author.ID
		}

		err = r.PostgreSQL.Eager().Create(&photo)
	}

	return &models.Photo{}, nil
}

func (r *Resource) GenerateSampleVersion() (*models.Version, error) {
	version := models.Version{
		BuildVersion: "0.3.0",
		BuildNumber:  300,
		Url:          "https://www.dropbox.com/s/x3ki4sjh2ky5ba1/Artify_0.3.0.zip?dl=1",
		Notes:        "+ Support Artify Status bar app with foundation logic\n+ Auto Update app\n+ Fetch image to artify-core\n+ Generate beautiful wallpaper\n+ Get feature photo",
	}
	err := r.PostgreSQL.Create(&version)
	if err != nil {
		return nil, err
	}
	return &version, nil
}

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
