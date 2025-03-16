package models

type Book struct {   
	ID        uint    `json:"id" gorm:"primary_key"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Year      int     `json:"year"`
	Pages     int     `json:"pages"`
	Language  string  `json:"language"`
	Publisher string  `json:"publisher"`
	Rating    float64 `json:"rating"`
}
