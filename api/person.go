package api

import (
	"context"
	"encoding/json"
	"fmt"
	"go-server/db"
	"go-server/model"
	"net/http"
	"time"
)

// CreatePerson - create person end point
func CreatePerson(response http.ResponseWriter, request *http.Request) {
	client, err := db.Connection()
	fmt.Println("run")
	if err != nil {
		fmt.Println(err)
	}
	response.Header().Add("content-type", "application/json")
	var person model.Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("goserver").Collection("persons")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}
