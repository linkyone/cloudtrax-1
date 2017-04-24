package models

import (
	"os"
	"strconv"
)

//Environment holds information about the current environment
type Environment struct {
	Address     string //CLOUDTRAX_SERVER_ADDRESS
	DatabaseURI string //CLOUDTRAX_SERVER_DATABASEURI
	Secret      string //CLOUDTRAX_SERVER_SECRET
	Debug       bool   //CLOUDTRAX_SERVER_DEBUG
}

//Parse gets environment variables from the system
func (env *Environment) Parse() {
	getEnv := func(n, d string) string {
		//use the name and the default value to return an environment variable
		v := os.Getenv(n)
		if v == "" {
			return d
		}
		return v
	}
	env.Address = getEnv("CLOUDTRAX_SERVER_ADDRESS", ":8080")
	env.DatabaseURI = getEnv("CLOUDTRAX_SERVER_DATABASEURI", "")
	env.Secret = getEnv("CLOUDTRAX_SERVER_SECRET", "default")
	env.Debug, _ = strconv.ParseBool(getEnv("CLOUDTRAX_SERVER_DEBUG", "false"))
}

//NewEnvironment initializes and returns a new environment object
func NewEnvironment() *Environment {
	env := &Environment{}
	env.Parse()
	return env
}
