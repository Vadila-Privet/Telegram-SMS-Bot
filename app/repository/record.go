package repository

import (
	"fmt"

	"../models"

	"github.com/globalsign/mgo/bson"
)

const (
	collectionName = "records"

	databaseName = "info"
)

//GetAllRecords returns all records from db
func GetAllRecords() ([]*models.Record, error) {

	r := models.Connection.Session.DB(databaseName).C(collectionName)

	var results []*models.Record

	err := r.Find(nil).Sort("-timestamp").All(&results)
	if err != nil {
		fmt.Println("Error is: ", err)
		return nil, err
	}

	return results, nil
}

//GetRecordsByName returns records from certain user
func GetRecordsByName(name string) ([]*models.Record, error) {

	r := models.Connection.Session.DB(databaseName).C(collectionName)

	var results []*models.Record

	err := r.Find(bson.M{"name": name}).Sort("-timestamp").All(&results)
	if err != nil {
		fmt.Println("Error is: ", err)
		return nil, err
	}

	return results, nil
}

//GetRecordsByPhone returns records that has been sended to certain number
func GetRecordsByPhone(phone string) ([]*models.Record, error) {

	r := models.Connection.Session.DB(databaseName).C(collectionName)

	var results []*models.Record

	err := r.Find(bson.M{"phone": phone}).Sort("-timestamp").All(&results)
	if err != nil {
		fmt.Println("Error is: ", err)
		return nil, err
	}

	return results, nil
}

