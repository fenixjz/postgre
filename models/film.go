package models

// Film represents the structure of a film entity.
type Film struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Actors      string `json:"actors"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
}
