package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Dashboard struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Type      string    `json:"type" db:"type"`
	Photo     Photo     `belongs_to:"photo"`
	PhotoID   uuid.UUID `json:"photo_id" db:"photo_id"`
}

// String is not required by pop and may be deleted
func (d Dashboard) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Dashboards is not required by pop and may be deleted
type Dashboards []Dashboard

// String is not required by pop and may be deleted
func (d Dashboards) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Dashboard) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Dashboard) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
