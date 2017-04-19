package cloudtrax

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/ryanhatfield/cloudtrax/models"
	"github.com/ryanhatfield/cloudtrax/models/accesspoints"
)

type Cloudtrax struct {
	Address        string
	Secret         string
	SessionSeconds int
	DownloadLimit  int
	UploadLimit    int
	DatabaseURI    string
	db             *sql.DB
}

func (ct *Cloudtrax) initializeDB() error {
	if ct.DatabaseURI == "" {
		return fmt.Errorf("DatabaseURI is required, and is not set.\nCLOUDTRAX_SERVER_DATABASEURI: %v", ct.DatabaseURI)
	}

	db, err := gorm.Open("postgres", ct.DatabaseURI)
	if err != nil {
		return fmt.Errorf("Unable to connect to database.\nError:\n%+v", err)
	}
	defer db.Close()

	db.AutoMigrate(&models.User{})

	return nil
}

// ListenAndServe sets up and starts the service
func (ct *Cloudtrax) ListenAndServe() error {
	err := ct.initializeDB()
	if err != nil {
		return err
	}
	return http.ListenAndServe(ct.Address, ct)
}

func (ct *Cloudtrax) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println("error parsing request form")
		return
	}

	request := accesspoints.ParseRequest(&r.Form)
	response := accesspoints.APResponse{
		Request: &request,
	}

	switch request.RequestType {
	case accesspoints.AccountingRequest:
		response.ResponseCode = accesspoints.OKCode
	case accesspoints.StatusRequest:
		response.ResponseCode = accesspoints.RejectCode
		response.BlockedMessage = "Your session has expired."
	case accesspoints.LoginRequest:
		//TODO: Check login credentials here
		response.ResponseCode = accesspoints.AcceptCode
		response.Seconds = 3600
		response.Download = 2000
		response.Upload = 800
	default:
		log.Printf("Error: %v, URL: %v", "incorrect request type", r.URL)
		return
	}

	err = response.Execute(&w)
	if err != nil {
		log.Printf("error while handling Accounting Request response: %v\n", err)
	}
}
