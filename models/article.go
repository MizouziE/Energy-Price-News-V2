package models

type Article struct {
	ID       uint    `json:"id" gorm:"unique"`
	Title    string  `json:"title"`
	Url      string  `json:"url" gorm:"unique"`
	ImageUrl string  `json:"imageUrl"`
	Source   *Source `json:"source"`
	Region   *Region `json:"region"`
}
