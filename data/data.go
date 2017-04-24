package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ryanhatfield/cloudtrax/data/models"
	//This is required for the postgres driver within gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Data holds information about the database
type Data interface {
	Ping() error
	AuthorizeSession(session string, a models.Authorization) error
	IsSessionAuthorized(sessionID string, deviceMac string) (bool, error)
	SaveRequest(req *models.APRequest) error
}

type data struct {
	Env *models.Environment
}

func (d data) open() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", d.Env.DatabaseURI)
	if err != nil {
		return nil, err
	}
	if d.Env.Debug {
		return db.Debug(), nil
	}
	return db, nil
}

func (d data) Ping() error {
	db, err := d.open()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (d data) AuthorizeSession(session string, a models.Authorization) error {
	return nil
}

func (d data) SaveRequest(req *models.APRequest) error {
	db, err := d.open()
	if err != nil {
		return err
	}
	defer db.Close()

	var s = &models.Session{}
	db.First(s, models.Session{Session: req.Session})
	if s == nil {

	}

	switch req.RequestType {
	case models.StatusRequest:

		break
	case models.LoginRequest:
		break
	case models.AccountingRequest:
		break
	}

	return nil
}

// //UpdateSession updates the session information in the DB
// func (d data) UpdateSession(s *models.Session) error {
// 	db, err := d.open()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()
// 	db.FirstOrCreate(s, models.Session{Session: s.Session})
// 	t := time.Now()
// 	s.Authorizations = append(s.Authorizations, models.Authorization{Device: "00:01:02:03:04", ExpirationTime: t})
// 	db.Save(s)
//
// 	log.Printf("Session Object from DB:\n%v", *s)
// 	return nil
// }

func (d data) AddAuthorization(sessionID string, a *models.Authorization) error {
	return nil
}

func (d data) IsSessionAuthorized(sessionID string, deviceMac string) (bool, error) {
	//Check the database for sessions already authorized
	db, err := d.open()
	if err != nil {
		return false, err
	}
	err = db.Close()
	if err != nil {
		return false, err
	}

	return false, nil
}

func (d data) initializeDB() error {
	db, err := d.open()
	if err != nil {
		return fmt.Errorf("Unable to connect to database.\nError:\n%s", err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.Authorization{})

	return nil
}

//NewData generates a new Data object from
func NewData(env *models.Environment) (Data, error) {
	d := &data{Env: env}
	err := d.Ping()
	if err != nil {
		return nil, err
	}
	err = d.initializeDB()
	if err != nil {
		return nil, err
	}
	return d, nil
}
