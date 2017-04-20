package models

import "github.com/jinzhu/gorm"

//User holds information about a User, their devices, and their sessions
type User struct {
	gorm.Model
	UserID string `gorm:"primary_key:true"`
	Email  string
}
