package main

import (
	"log"
	"os"

	"github.com/ryanhatfield/cloudtrax/data"
	"github.com/ryanhatfield/cloudtrax/data/models"
)

func main() {
	d, err := data.NewData(&models.Environment{
		DatabaseURI: os.Getenv("CLOUDTRAX_SERVER_DATABASEURI"),
		Debug:       true,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	d.Ping()

	d.UpdateSession(&models.Session{Session: "12345678"})

}
