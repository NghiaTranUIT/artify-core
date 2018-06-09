package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Author struct {
	ID             uuid.UUID `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Name           string    `json:"name" db:"name"`
	Born           string    `json:"born" db:"born"`
	Died           string    `json:"died" db:"died"`
	Nationality    string    `json:"nationality" db:"nationality"`
	Wikipedia      string    `json:"wikipedia" db:"wikipedia"`
	Photos         Photos    `json:"photos" has_many:"books"`
	OriginalSource string    `json:"original_source" db:"original_source"`
}

// String is not required by pop and may be deleted
func (a Author) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Authors is not required by pop and may be deleted
type Authors []Author

// String is not required by pop and may be deleted
func (a Authors) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Author) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Author) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
