package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string `json:"title"`
	Url      string `json:"url" gorm:"unique"`
	ImageUrl string `json:"imageUrl"`
	SourceID uint   `json:"sourceId"`
	Source   Source
}
