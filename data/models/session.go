package models

type Session struct {
	SessionID string `gorm:"primary_key:true"`
	UserID    string
}
