package database

import (
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	connection *sqlx.DB
}

func (db *Database) Connect() error {
	var err error
	db.connection, err = sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.WithField("databaseIdentify", ":memory:").Fatal(err)
	}
	if err = db.connection.Ping(); err != nil {
		log.WithField("databaseIdentify", ":memory:").Fatal(err)
	}
	log.WithField("databaseIdentify", ":memory:").Debug("Database connected")
	return nil
}

func (db *Database) Migrate(schemaFile string) error {
	schemaQuery, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		log.Fatal("Cannot read schema file")
	}
	db.connection.MustExec(string(schemaQuery))
	log.WithField("schemaFile", schemaFile).Debug("Migrate successful")
	return nil
}

func (db *Database) Select(model []interface{}, query string) error {
	err := db.connection.Select(model, query)
	if err != nil {
		log.WithField("query", query).Error(err)
		return err
	}
	log.WithField("query", query).Debug("Query success")
	return nil
}

func (db *Database) SelectOne(model interface{}, query string) error {
	err := db.connection.Get(model, query)
	if err != nil {
		log.WithField("query", query).Error(err)
		return err
	}
	log.WithField("query", query).Debug("Query success")
	return nil
}

func (db *Database) Insert(query string) (int64, error) {
	result, err := db.connection.Exec(query)
	if err != nil {
		log.WithField("query", query).Error(err)
		return 0, err
	}
	log.WithField("query", query).Debug("Query success")
	return result.LastInsertId()
}

func (db *Database) UpdateDelete(query string) (int64, error) {
	result, err := db.connection.Exec(query)
	if err != nil {
		log.WithField("query", query).Error(err)
		return 0, err
	}
	log.WithField("query", query).Debug("Query success")
	return result.RowsAffected()
}
