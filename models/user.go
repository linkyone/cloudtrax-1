package models

import "github.com/jinzhu/gorm"

type Status struct {
	gorm.Model
	UserID uint
}

type User struct {
	gorm.Model
	Email    string
	Statuses []Status
}
