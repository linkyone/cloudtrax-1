package cloudtrax

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ryanhatfield/cloudtrax/data"
	"github.com/ryanhatfield/cloudtrax/data/models"
)

//Cloudtrax holds information about connecting with cloudtrax APs
type Cloudtrax struct {
	Env  *models.Environment
	Data data.Data
}

func logInterface(i interface{}) {
	b, e := json.MarshalIndent(i, "", " ")
	if e != nil {
		log.Printf("%+v\n", i)
	}
	log.Printf("\n%s", string(b))
}

// ListenAndServe sets up and starts the service
func (ct *Cloudtrax) ListenAndServe() error {

	router := httprouter.New()
	router.GET("/:site/sessions/:session", ct.getSession)
	router.GET("/:site/sessions/:session/:device", ct.getSession)
	router.GET("/:site/sessions/:session/:device/authorize", ct.authorizeSession)
	router.GET("/:site/auth.html", ct.handleAPRequest)
	return http.ListenAndServe(":"+ct.Env.Port, router)
}

func (ct *Cloudtrax) getSession(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ses := p.ByName("session")
	sit := p.ByName("site")
	dev := p.ByName("device")
	if ses == "" || sit == "" || dev == "" {
		http.Error(w, fmt.Errorf("request invalid").Error(), http.StatusBadRequest)
	}

	session, err := ct.Data.FindSession(ses, sit, dev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session == nil {
		http.Error(w, fmt.Errorf("session '%s' was not found", ses).Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ct *Cloudtrax) authorizeSession(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := p.ByName("session")
	device := p.ByName("device")
	site := p.ByName("site")
	if session != "" && device != "" && site != "" {
		go func(ses, dev, sit string) {
			//TODO: do something fun here.
			log.Printf("Session: %s, Device: %s, Site: %s", ses, dev, sit)
		}(session, device, site)
	} else {
		http.Error(w, fmt.Errorf("request invalid").Error(), http.StatusBadRequest)
	}
}

func (ct *Cloudtrax) handleAPRequest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Println("error parsing request form")
		return
	}

	request := models.NewAPRequest(&r.Form)
	response := models.NewAPResponse(request)

	data, derr := data.NewData(ct.Env)

	if derr != nil {
		log.Println("error creating data object")
	} else {
		derr = data.SaveAPRequest(*request, p.ByName("site"))
		if derr != nil {
			log.Printf("error saving to db:\n%s", derr.Error())
		}
	}

	//Get the new response authorization
	response.ResponseAuthorization, err = models.GenerateRA(response.ResponseCode,
		request.RequestAuthorization,
		ct.Env.Secret)

	switch request.RequestType {
	case models.StatusRequest:
		break

	case models.LoginRequest:
		break

	case models.AccountingRequest:
		break

	default:
		log.Printf("unknown request type")
		return
	}

	if err != nil {
		//nothing will work after this, should we do something here?
		log.Printf("error occured while generating the response authenticator:\n%s", err.Error())
	}

	go logInterface(response)

	err = response.Execute(&w)
	if err != nil {
		log.Printf("error while handling Accounting Request response: %s\n", err.Error())
	}
}

//NewCloudtrax initializes and returns a new cloudtrax object
func NewCloudtrax(env *models.Environment, data *data.Data) *Cloudtrax {
	return &Cloudtrax{Env: env, Data: *data}
}
