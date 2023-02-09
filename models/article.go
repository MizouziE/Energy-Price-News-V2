package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string `json:"title"`
	Url      string `json:"url" gorm:"unique"`
	ImageUrl string `json:"imageUrl" default:"unassigned"`
	// SourceID uint   `json:"sourceId" default:"unassigned"`
	// Source   Source
}
