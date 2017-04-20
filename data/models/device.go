package models

//Device holds information about a specific mac address
type Device struct {
	MacAddress  string       `gorm:"primary_key:true"`
	Accountings []Accounting `gorm:"ForeignKey:MacAddress;AssociationForeignKey:MacAddress"`
}
