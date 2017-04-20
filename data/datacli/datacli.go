package main

import "github.com/ryanhatfield/cloudtrax/data"

func main() {
	d, err := data.Factory("")
	if err != nil {
		return err
	}
	d.Ping()
}
