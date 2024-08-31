package dao

import (
	"context"
	"rest-api/dbConfig"
	"rest-api/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindStudentByStudentId(studentId string) (*dto.Student, error){
	var object dto.Student

	err := dbConfig.DATABASE.Collection("Students").FindOne(context.Background(), bson.M{"studentid":studentId, "deleted":false}).Decode(&object)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else{
			return nil, err
		}
	}
	return &object, nil
}