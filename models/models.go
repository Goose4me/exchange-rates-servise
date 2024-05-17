package models

import "gorm.io/gorm"

type Subscriber struct {
	gorm.Model
	Email string `json:"email" gorm:"unique;not null"`
}
