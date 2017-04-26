package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//This is required for the postgres driver within gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ryanhatfield/cloudtrax/data/models"
)

//Data holds information and methods about the datastorage
type Data interface {
	//Ping verifies that the database is available
	Ping() error
	//FindSession returns a Session object from the provided site and session IDs
	FindSession(session, site, device string) (*models.Session, error)
	//UpdateSession attempts to insert or update the Session as needed
	UpdateSession(session models.Session) error
}

type data struct {
	Env *models.Environment
}

func (d data) openDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", d.Env.DatabaseURI)
	if err != nil {
		return nil, err
	}
	if d.Env.Debug {
		return db.Debug(), nil
	}
	return db, nil
}

func (d data) initializeDB() error {
	db, err := d.openDB()
	if err != nil {
		return fmt.Errorf("Unable to connect to database.\nError:\n%s", err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&models.Session{})

	return nil
}

func (d data) Ping() error {
	db, err := d.openDB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (d data) FindSession(ses, sit, dev string) (*models.Session, error) {
	if ses == "" {
		return nil, fmt.Errorf("session id invalid:%s", ses)
	}
	session := &models.Session{}

	db, err := d.openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.First(session, models.Session{Session: ses, Site: sit, Device: dev})
	if err != nil {
		return nil, err
	}

	if session.Session != ses {
		return nil, nil
	}

	return session, nil
}

func (d data) UpdateSession(session models.Session) error {
	if session.Session == "" {
		return fmt.Errorf("session can't be nil")
	}

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
