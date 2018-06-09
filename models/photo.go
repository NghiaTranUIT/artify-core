package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Photo struct {
	ID             uuid.UUID `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Name           string    `json:"name" db:"name"`
	ImageUrl       string    `json:"image_url" db:"image_url"`
	Author         Author    `json:"author" belongs_to:"author"`
	AuthorID       uuid.UUID `json:"-" db:"author_id"`
	Width          uint      `json:"width" db:"width"`
	Height         uint      `json:"height" db:"height"`
	Info           string    `json:"info" db:"info"`
	Date           string    `json:"date" db:"date"`
	Style          string    `json:"style" db:"style"`
	Location       string    `json:"location" db:"location"`
	Dimensions     string    `json:"dimensions" db:"dimensions"`
	Media          string    `json:"media" db:"media"`
	OriginalSource string    `json:"original_source" db:"original_source"`
}

// String is not required by pop and may be deleted
func (p Photo) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Photos is not required by pop and may be deleted
type Photos []Photo

// String is not required by pop and may be deleted
func (p Photos) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Photo) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Photo) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
