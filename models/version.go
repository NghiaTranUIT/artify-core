package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Version struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	BuildVersion string    `json:"build_version" db:"build_version"`
	BuildNumber  uint      `json:"build_number" db:"build_number"`
	Url          string    `json:"url" db:"url"`
	Notes        string    `json:"notes" db:"notes"`
}

// String is not required by pop and may be deleted
func (v Version) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Versions is not required by pop and may be deleted
type Versions []Version

// String is not required by pop and may be deleted
func (v Versions) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (v *Version) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (v *Version) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
