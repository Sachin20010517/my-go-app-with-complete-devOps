package dao

import (
	"context"
	"errors"
	"rest-api/dbConfig"

	"go.mongodb.org/mongo-driver/bson"
)

func DB_DeleteStudent(studentId string) error {
	result, err := dbConfig.DATABASE.Collection("Students").UpdateOne(context.Background(), bson.M{"studentid": studentId}, bson.M{"$set": bson.M{"deleted": true}})

	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 {
		return errors.New("Specified Id not found!")
	}
	return nil
}