package main

import (
	"fmt"
	"os"

	"github.com/ryanhatfield/cloudtrax"
)

func main() {
	fmt.Println("Cloudtrax AP Server Starting")

	env := func(n, d string) string {
		//use the name and the default value to return an environment variable
		v := os.Getenv(n)
		if v == "" {
			return d
		}
		return v
	}

	ct := cloudtrax.Cloudtrax{
		Secret:      env("CLOUDTRAX_SERVER_SECRET", "default"),
		Address:     env("CLOUDTRAX_SERVER_ADDRESS", ":8080"),
		DatabaseURI: env("CLOUDTRAX_SERVER_DATABASEURI", ""),
	}

	err := ct.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Cloudtrax AP Server Shutting Down")
}
