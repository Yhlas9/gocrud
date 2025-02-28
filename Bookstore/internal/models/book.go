package models

type Book struct {   //Book struktura
	ID        uint    `json:"id" gorm:"primary_key"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Year      int     `json:"year"`
	Pages     int     `json:"pages"`
	Language  string  `json:"language"`
	Publisher string  `json:"publisher"`
	Rating    float64 `json:"rating"`
}
// Dusnukli bulara
// Book struct should have the following fields:
// - id (auto-incremented, primary key)
// - title (string)
// - author (string)
// - year (int)
// - pages (int)
// - language (string)
// - publisher (string)
// - rating (float64)