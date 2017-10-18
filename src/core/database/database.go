package database

//
//import (
//	"io/ioutil"
//
//	"github.com/jmoiron/sqlx"
//	log "github.com/sirupsen/logrus"
//)
//
//type Database struct {
//	connection *sqlx.DB
//}
//
//func (db *Database) Connect() error {
//	var err error
//	DbConnection, err = sqlx.Connect("sqlite3", dbInfo.SQLite.DatabaseFileName)
//	if err != nil {
//		log.WithField("databaseIdentify", dbInfo.SQLite.DatabaseFileName).Fatal(err)
//	}
//	if err = DbConnection.Ping(); err != nil {
//		log.WithField("databaseIdentify", dbInfo.SQLite.DatabaseFileName).Fatal(err)
//	}
//	log.WithField("databaseIdentify", dbInfo.SQLite.DatabaseFileName).Debug("Database connected")
//	return nil
//}
//
//func (db *Database) Migrate() error {
//	schemaQuery, err := ioutil.ReadFile(dbInfo.SQLite.SchemaFile)
//	_, err = DbConnection.Exec(string(schemaQuery))
//	if err != nil {
//		log.WithField("query", schemaQuery).Fatal(err)
//	}
//	log.WithField("schemaFile", dbInfo.SQLite.SchemaFile).Debug("Migrate successful")
//	return nil
//}
//
//func (db *Database) Select(model []interface{}, query string) error {
//	err := DbConnection.Select(model, query)
//	if err != nil {
//		log.WithField("query", query).Error(err)
//		return err
//	}
//	log.WithField("query", query).Debug("Query success")
//	return nil
//}
//
//func (db *Database) SelectOne(model interface{}, query string) error {
//	err := DbConnection.Get(model, query)
//	if err != nil {
//		log.WithField("query", query).Error(err)
//		return err
//	}
//	log.WithField("query", query).Debug("Query success")
//	return nil
//}
//
//func Insert(query string) (int64, error) {
//	result, err := DbConnection.Exec(query)
//	if err != nil {
//		log.WithField("query", query).Error(err)
//		return 0, err
//	}
//	log.WithField("query", query).Debug("Query success")
//	return result.LastInsertId()
//}
//
//func Delete(query string) (int64, error) {
//	result, err := DbConnection.Exec(query)
//	if err != nil {
//		log.WithField("query", query).Error(err)
//		return 0, err
//	}
//	log.WithField("query", query).Debug("Query success")
//	return result.RowsAffected()
//}
