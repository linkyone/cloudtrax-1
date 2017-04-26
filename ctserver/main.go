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
	ct := cloudtrax.NewCloudtrax(env)

	func() {
		//spin up the data object, and verify database connection
		//do this in function to clean up var namespace after
		//TODO: remove this or move it to proper method
		data, err := data.NewData(ct.Env)
		if err != nil {
			fmt.Printf("Error while starting db:\n%s", err.Error())
			os.Exit(1)
		}
		err = data.Ping()
		if err != nil {
			fmt.Printf("Error while connecting to db:\n%s", err.Error())
		}
	}()

	err := ct.ListenAndServe()
	fmt.Println("Cloudtrax AP Server Shutting Down")
	if err != nil {
		fmt.Printf("Error while exiting:\n%s", err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
