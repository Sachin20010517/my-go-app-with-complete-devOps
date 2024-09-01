package dao

import (
	"context"
	"errors"
	"rest-api/dbConfig"
	"rest-api/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_GetAllStudent () (*[]dto.Student, error) {
	var objects []dto.Student
	results, err := dbConfig.DATABASE.Collection("Students").Find(context.Background(), bson.M{"deleted":false})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else{
		return nil, nil
		}
	}

	for results.Next(context.Background()) {
		var object dto.Student
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Deals")
		}
		objects = append(objects, object)
	}
	return &objects, nil

}
