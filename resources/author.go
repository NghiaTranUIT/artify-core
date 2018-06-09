package resources

import (
	"github.com/NghiaTranUIT/artify-core/models"
)

func (r *Resource) CreateAuthorByWikiart(param models.AuthorParam) (*models.Author, error) {
	author := models.Author{}

	// Search by name first
	err := r.PostgreSQL.Where("name = ?", param.Name).First(&author)

	// Create new author
	if err != nil {
		newAuthor := models.Author{
			Name:        param.Name,
			Born:        param.Born,
			Died:        param.Died,
			Nationality: param.Nationality,
			Wikipedia:   param.Wikipedia,
		}
		err = r.PostgreSQL.Create(&newAuthor)
		if err != nil {
			return nil, err
		}
		return &newAuthor, nil
	}

	return &author, nil
}
