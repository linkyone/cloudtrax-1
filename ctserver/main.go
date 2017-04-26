package main

import (
	"fmt"
	"os"

	"github.com/ryanhatfield/cloudtrax"
	"github.com/ryanhatfield/cloudtrax/data"
	"github.com/ryanhatfield/cloudtrax/data/models"
)

func main() {
	fmt.Println("Cloudtrax AP Server Starting")
	env := models.NewEnvironment()

	data, err := data.NewData(&env)
	if err != nil {
		fmt.Printf("Error while starting db:\n%s", err.Error())
		os.Exit(1)
	}
	ct := cloudtrax.NewCloudtrax(&env, &data)

	err = ct.ListenAndServe()
	fmt.Println("Cloudtrax AP Server Shutting Down")
	if err != nil {
		fmt.Printf("Error while exiting:\n%s", err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
