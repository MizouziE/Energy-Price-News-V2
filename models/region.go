package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	Name     string   `json:"name"`
	Endpoint string   `json:"endpoint"`
	Source   []Source `gorm:"-"`
}
