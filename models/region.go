package models

type Region struct {
	ID       uint   `json:"id" gorm:"unique"`
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
}
