package models

import "gorm.io/gorm"

type Source struct {
	gorm.Model
	Name     string    `json:"name"`
	Endpoint string    `json:"endpoint"`
	ImageUrl string    `json:"imageUrl"`
	RegionID uint      `json:"regionId"`
	Article  []Article `gorm:"-"`
}
