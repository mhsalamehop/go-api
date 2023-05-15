package models

type SavedMovies struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	Overview  string  `json:"overview"`
	SavedInfo []Infos `json:"saved_info"`
}

type Infos struct {
	SavedBy string    `json:"saved_by"`
	SavedAt string `json:"saved_at"`
}
