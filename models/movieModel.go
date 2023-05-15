package models

type Movies struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	Release_date string `json:"release_date"`
	Tests	[]Test
}

type Test struct {
	Id int 

}
