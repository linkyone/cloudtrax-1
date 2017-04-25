package cloudtrax

import (
	"log"
	"net/http"
	"strings"

	"github.com/ryanhatfield/cloudtrax/data"
	"github.com/ryanhatfield/cloudtrax/data/models"
)

//Cloudtrax holds information about connecting with cloudtrax APs
type Cloudtrax struct {
	Address        string
	Secret         string
	SessionSeconds int
	DownloadLimit  int
	UploadLimit    int
	env            *models.Environment
}

func (ct *Cloudtrax) logAccounting(req models.APRequest) {

}

// ListenAndServe sets up and starts the service
func (ct *Cloudtrax) ListenAndServe() error {
	return http.ListenAndServe(ct.Address, ct)
}

func (ct *Cloudtrax) handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/sessions") {
		if strings.HasPrefix(r.URL.Path, "/api/sessions/authorize") {
			a, err := models.NewAuthorization(r.Form)
			if err != nil {
				log.Printf("error occured while handling a session authorization request:\n%s", err.Error())
			}
			data, derr := data.NewData(ct.env)
			if derr != nil {
				log.Printf("error occured while handling a session authorization request:\n%s", derr.Error())
			}
			data.AuthorizeSession("test", *a)
		}
	}
}

func (ct *Cloudtrax) handleAPRequest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("error parsing request form")
		return
	}

	request := models.NewAPRequest(&r.Form)
	response := models.NewAPResponse(request)

	go func(req models.APRequest, res models.APResponse) {
		//log the AP request and responses
		data, derr := data.NewData(ct.env)
		if derr != nil {
			log.Printf("error occured while logging the request:\n%s", derr.Error())
		}
		derr = data.SaveRequest(request)
		if derr != nil {
			log.Printf("error occured while logging the request:\n%s", derr.Error())
		}
	}(*request, *response)

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
	} else if strings.HasPrefix(r.URL.Path, "/api") {
		log.Printf("Handling api call: %s", r.URL.Path)
		ct.handleAPIRequest(w, r)
		return
	} else if strings.HasPrefix(r.URL.Path, "/favicon.ico") {
		//Let's just eat this one, it's annoying
		return
	}
	//else
	log.Printf("Unknown endpoint: %s", r.URL.Path)
}

//NewCloudtrax initializes and returns a new cloudtrax object
func NewCloudtrax(env *models.Environment) *Cloudtrax {
	return &Cloudtrax{env: new(models.Environment)}
}
