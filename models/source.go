package models

type Source struct {
	ID       uint    `json:"id" gorm:"unique"`
	Name     string  `json:"name"`
	Endpoint string  `json:"endpoint"`
	ImageUrl string  `json:"imageUrl"`
	Region   *Region `json:"region"`
}
