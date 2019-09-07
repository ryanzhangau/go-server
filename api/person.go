package api

import (
	"context"
	"encoding/json"
	"fmt"
	"go-server/db"
	"go-server/model"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var client, err = db.Connection()
var collection = client.Database("goserver").Collection("persons")

func getDbConnectErr() {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// CreatePerson - create person end point
func CreatePerson(response http.ResponseWriter, request *http.Request) {
	getDbConnectErr()
	response.Header().Add("content-type", "application/json")
	var person model.Person
	json.NewDecoder(request.Body).Decode(&person)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

// AllPersons - get all persons
func AllPersons(response http.ResponseWriter, request *http.Request) {
	getDbConnectErr()
	response.Header().Add("content-type", "application/json")
	var persons model.Persons
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "}` + err.Error() + `"}`))
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var person model.Person
		cursor.Decode(&person)
		persons = append(persons, person)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "}` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(persons)
}
