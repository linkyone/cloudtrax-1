package main

import (
	"fmt"
	"os"

	"github.com/ryanhatfield/cloudtrax"
	"github.com/ryanhatfield/cloudtrax/data/models"
)

func main() {
	fmt.Println("Cloudtrax AP Server Starting")

	ct := cloudtrax.NewCloudtrax(models.NewEnvironment())

	err := ct.ListenAndServe()
	fmt.Println("Cloudtrax AP Server Shutting Down")
	if err != nil {
		fmt.Printf("Error while exiting:\n%s", err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
