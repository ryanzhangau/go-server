package api

import (
	"fmt"
	"net/http"
)

// Homepage - homepage end point
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage End Point")
}
