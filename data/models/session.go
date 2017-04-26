package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Session holds information about a WiFi user's session
type Session struct {
	gorm.Model
	Session   string
	Site      string
	Node      string
	IPv4      string
	Device    string
	ExpiresAt time.Time
}
