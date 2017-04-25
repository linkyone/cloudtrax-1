package models

import (
	"fmt"
	"net/url"
	"time"
)

//Authorization holds information about an authorized device and session
type Authorization struct {
	AuthorizationID string
	SessionID       string
	Device          string
	ExpirationTime  time.Time
}

//NewAuthorization returns a new uathorization, from URL parameters
func NewAuthorization(v url.Values) (*Authorization, error) {
	var errs []error
	get := func(key string, defaultValue string, required bool) string {
		value := v.Get(key)
		if required && value == "" {
			errs = append(errs, fmt.Errorf("The following parameter is required and was not found: %s", key))
		} else if value == "" {
			value = defaultValue
		}
		return value
	}

	var err error

	apr := &Authorization{
		Device: get("device", "", true),
	}
	apr.ExpirationTime, err = time.Parse("Mon Jan 2 15:04:05 -0700 CST 2006", get("expirationTime", "", true))
	if err != nil {
		errs = append(errs, fmt.Errorf("the following error occured: %s", err.Error()))
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("the following errors occured:\n %v", errs)
	}
	return apr, nil
}
