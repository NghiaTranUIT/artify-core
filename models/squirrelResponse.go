package models

import (
	"strconv"
	"time"
)

type SquirrelResponse struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	Notes   string `json:"notes"`
	PubDate string `json:"pub_date"`
}

func NewSuccessSquirrelReponse(version *Version) SquirrelResponse {
	name := "Artify " + version.BuildVersion + "(" + strconv.FormatUint(uint64(version.BuildNumber), 10) + ")"
	return SquirrelResponse{
		Name:    name,
		Notes:   "Latest Artify build",
		PubDate: time.Now().UTC().Format(time.RFC3339),
		Url:     version.Url,
	}
}
