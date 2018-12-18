package models

import "encoding/json"

type User struct {
	Id int `gorm:"column:Id"`
	Age int `gorm:"column:Age"`
	Name string `gorm:"column:Name"`
	Description string `gorm:"column:Description"`
	SiblingInfo json.RawMessage `gorm:"column:SiblingInfo"`
}

type SiblingInfo struct {
	Sibling string `json:"sibling"`
	Name string `json:"name"`
}
