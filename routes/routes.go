package routes

import (
	"fmt"
	"go-server/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RouteHandler - Handle routes
func RouteHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/", api.Homepage)
	router.HandleFunc("/person", api.CreatePerson).Methods("POST")
	router.HandleFunc("/persons", api.AllPersons).Methods("GET")
	router.HandleFunc("/articles", api.AllArticles)
	fmt.Println("Server is running on port 9090 ...")
	log.Fatal(http.ListenAndServe(":9090", router))
}
