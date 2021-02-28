package models

type Search struct {
	Title string `json:"Title"`
	Year string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type string `json:"Type"`
	Poster string `json:"Poster"`
}

type Result struct {
	Search []Search
	TotalResults string `json:"totalResults"`
	Response string `json:"response"`
}