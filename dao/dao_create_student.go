package dao

import (
	"context"
	"rest-api/dbConfig"
	"rest-api/dto"
)

func DB_CreateStudent (object *dto.Student) error {
	_, err := dbConfig.DATABASE.Collection("Students").InsertOne(context.Background(), object)
	if err != nil{
		return err
	}
	return nil
}