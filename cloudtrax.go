package cloudtrax

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	//This is required for the postgres driver within gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ryanhatfield/cloudtrax/data/models"
)

//Cloudtrax holds information about connecting with cloudtrax APs
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
		return fmt.Errorf("DatabaseURI is required, and is not set.\n"+
			"CLOUDTRAX_SERVER_DATABASEURI: %v", ct.DatabaseURI)
	}

	db, err := gorm.Open("postgres", ct.DatabaseURI)
	if err != nil {
		return fmt.Errorf("Unable to connect to database.\nError:\n%s", err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&models.User{})

	return nil
}

func (ct *Cloudtrax) logAccounting(req models.APRequest) {

}

// ListenAndServe sets up and starts the service
func (ct *Cloudtrax) ListenAndServe() error {
	err := ct.initializeDB()
	if err != nil {
		return err
	}

	return http.ListenAndServe(ct.Address, ct)
}

func (ct *Cloudtrax) handleAPRequest(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ct.DatabaseURI)
	if err != nil {
		log.Println("error connecting with database")
		return
	}
	defer db.Close()

	err = r.ParseForm()
	if err != nil {
		log.Println("error parsing request form")
		return
	}

	request := models.ParseRequest(&r.Form)
	response := models.APResponse{
		Request: &request,
	}

	switch request.RequestType {
	case models.AccountingRequest:
		response.ResponseCode = models.OKCode
		//start with a goroutine, so you don't hold up the response
		go ct.logAccounting(request)
	case models.StatusRequest:
		response.ResponseCode = models.RejectCode
		response.BlockedMessage = "Your session has expired."
	case models.LoginRequest:
		//TODO: Check login credentials here
		response.ResponseCode = models.AcceptCode
		response.Seconds = 3600
		response.Download = 2000
		response.Upload = 800
	default:
		log.Printf("Error: %v, URL: %v", "incorrect request type", r.URL)
		return
	}

	err = response.Execute(&w)
	if err != nil {
		log.Printf("error while handling Accounting Request response: %s\n", err.Error())
	}
}

func (ct *Cloudtrax) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/auth") {
		log.Printf("Handling auth call: %s", r.URL.Path)
		ct.handleAPRequest(w, r)
		return
	} else if strings.HasPrefix(r.URL.Path, "/rest") {
		log.Printf("Handling rest call: %s", r.URL.Path)
		//ct.handleRestRequest(w, r)
		return
	} else if strings.HasPrefix(r.URL.Path, "/favicon.ico") {
		//Let's just eat this one, it's annoying
		return
	}
	//else
	log.Printf("Unknown endpoint: %s", r.URL.Path)
}
