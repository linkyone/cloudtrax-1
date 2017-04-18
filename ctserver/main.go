package main

import (
	"cloudtrax"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	ct := cloudtrax.Cloudtrax{}

	if &ct != nil {
		return
	}
}
