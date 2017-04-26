package data

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	//This is required for the postgres driver within gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ryanhatfield/cloudtrax/data/models"
)

//Data holds information and methods about the datastorage
type Data interface {
	//FindSession returns a Session object from the provided site and session IDs
	FindSession(session, site, device string) (*models.Session, error)
	//UpdateSession attempts to insert or update the Session as needed
	UpdateSession(session models.Session) error
	//SaveAPRequest saves an ap request of type STATUS or ACCT
	SaveAPRequest(req models.APRequest, site string) error
}

type data struct {
	Env *models.Environment
	db  *gorm.DB
}

func (d data) FindSession(ses, sit, dev string) (*models.Session, error) {
	if ses == "" {
		return nil, fmt.Errorf("session id invalid:%s", ses)
	}
	session := &models.Session{}

	d.db.First(session, models.Session{Session: ses, Site: sit, Device: dev})

	if session.Session != ses {
		return nil, nil
	}

	return session, nil
}

func (d data) SaveAPRequest(req models.APRequest, site string) error {
	session := &models.Session{}

	d.db.First(session, models.Session{Session: req.Session, Site: site, Device: req.MacAddress})

	newSession := session.Session != req.Session
	if newSession {
		//session not found, save a new instance
		session.Session = req.Session
		session.Device = req.MacAddress
		session.Node = req.NodeAddress
		session.Site = site
	}

	//update info
	log.Println("Saving session")
	session.IPv4 = req.IPV4Address
	if req.RequestType == models.AccountingRequest {
		getint := func(s string) uint {
			u, _ := strconv.ParseUint(req.Download, 0, 32)
			return uint(u)
		}
		session.Download = getint(req.Download)
		session.Upload = getint(req.Upload)
		session.Seconds = getint(req.Seconds)
	}

	if newSession {
		d.db.Save(&session)
	} else {
		d.db.Model(&session).Updates(models.Session{
			Download: session.Download,
			Upload:   session.Upload,
			IPv4:     session.IPv4,
			Seconds:  session.Seconds,
		})
	}

	return nil
}

func (d data) UpdateSession(session models.Session) error {
	if session.Session == "" {
		return fmt.Errorf("session can't be nil")
	}

	dbSession := &models.Session{}
	d.db.FirstOrCreate(dbSession, models.Session{Session: session.Session, Site: session.Site, Device: session.Device})

	return nil
}

//NewData generates a new Data object from
func NewData(env *models.Environment) (Data, error) {

	db, err := gorm.Open("postgres", env.DatabaseURI)
	if err != nil {
		return nil, err
	}

	log.Printf("Max DB connections: %v", env.MaxDatabaseConnections)

	db.DB().SetMaxOpenConns(env.MaxDatabaseConnections)
	db.DB().SetMaxIdleConns(env.MaxDatabaseConnections)
	db.AutoMigrate(&models.Session{})

	d := data{Env: env, db: db}
	if env.Debug {
		d.db = db.Debug()
	}
	return d, nil
}
