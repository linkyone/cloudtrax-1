package models

//Accounting holds details about the referenced session
type Accounting struct {
	MacAddress  string `gorm:"primary_key:true"`
	NodeAddress string
	Download    uint
	Upload      uint
	Seconds     uint
	IPv4Address string
	Session     string
}
