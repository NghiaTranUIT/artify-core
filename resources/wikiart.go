package resources

import "github.com/NghiaTranUIT/artify-core/models"

func (r *Resource) CreatePhotoAuthorFromWikiart(param models.WikiartParam) (*models.Photo, error) {

	author, err := r.CreateAuthorByWikiart(param.Author)
	if err != nil {
		return nil, err
	}
	photo, err := r.CreatePhotoByWikiart(param.Photo, *author)
	if err != nil {
		return nil, err
	}
	photo.Author = *author
	return photo, nil
}
