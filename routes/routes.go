package routes

import (
	"fmt"
	"go-server/api"
	"log"
	"net/http"
)

// RouteHandler - Handle routes
func RouteHandler() {
	http.HandleFunc("/", api.Homepage)
	http.HandleFunc("/articles", api.AllArticles)
	fmt.Println("Server is running on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
