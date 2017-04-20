package data

import "github.com/jinzhu/gorm"

//Data holds information about the database
type Data interface {
	Ping() error
}

type data struct {
	DatabaseURI string
}

func (d data) Ping() error {
	db, err := gorm.Open("postgres", d.DatabaseURI)
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}

//Factory generates a new Data object from
func Factory(databaseURI string) (Data, error) {
	d := &data{}
	err := d.Ping()
	if err != nil {
		return nil, err
	}
	return d, nil
}
