package models

type PhotoParam struct {
	Name           string `json:"name"`
	Date           string `json:"date"`
	Style          string `json:"style"`
	Genre          string `json:"genre"`
	Media          string `json:"media"`
	Tag            string `json:"tag"`
	Location       string `json:"location"`
	Info           string `json:"info"`
	Url            string `json:"url"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	OriginalSource string `json:"original_source"`
}

type AuthorParam struct {
	Name           string `json:"name"`
	Born           string `json:"born"`
	Died           string `json:"died"`
	Nationality    string `json:"nationality"`
	ArtMovement    string `json:"art movement"`
	Wikipedia      string `json:"wikipedia"`
	OriginalSource string `json:"original_source"`
}

type WikiartParam struct {
	Photo  PhotoParam  `json:"photo"`
	Author AuthorParam `json:"author"`
}
