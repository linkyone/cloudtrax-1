package models

import "net/url"

//LoginRequest string representation
const LoginRequest string = "login"

//StatusRequest string representation
const StatusRequest string = "status"

//AccountingRequest string representation
const AccountingRequest string = "acct"

//APRequest has all variables expected in requests from an AP
type APRequest struct {
	RequestType          string
	RequestAuthorization string
	MacAddress           string
	Username             string
	Password             string
	NodeAddress          string //Mac address of the AP it's connected to
	IPV4Address          string
	Session              string
	Download             string
	Upload               string
}

func ParseRequest(v *url.Values) APRequest {
	return APRequest{
		RequestType:          v.Get("type"),
		RequestAuthorization: v.Get("ra"),
		MacAddress:           v.Get("mac"),
		Username:             v.Get("username"),
		Password:             v.Get("password"),
		NodeAddress:          v.Get("node"),
		IPV4Address:          v.Get("ipv4"),
		Session:              v.Get("session"),
		Download:             v.Get("download"),
		Upload:               v.Get("upload"),
	}
}
