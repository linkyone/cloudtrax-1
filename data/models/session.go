package models

//Session holds information about a WiFi user's session
type Session struct {
	SessionID       uint `gorm:"primary_key:true;AUTO_INCREMENT"`
	Session         string
	RequestingNode  string
	LatestNode      string
	FirstIPAddress  string
	LatestIPAddress string
	TotalDownload   uint
	TotalUpload     uint
	Authorizations  []Authorization `gorm:"ForeignKey:AuthorizationID"`
}
