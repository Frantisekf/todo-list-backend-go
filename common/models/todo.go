package models

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Text string `json:"text"`
}
