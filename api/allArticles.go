package api

import (
	"encoding/json"
	"go-server/structs"
	"net/http"
)

// AllArticles - Get all articles
func AllArticles(w http.ResponseWriter, r *http.Request) {
	articles := structs.Articles{
		structs.Article{
			Title:       "Article Title 1",
			Description: "Article Description 1",
			Content:     "Article Content 1",
		},
		structs.Article{
			Title:       "Article Title 2",
			Description: "Article Description 2",
			Content:     "Article Content 2",
		},
	}

	json.NewEncoder(w).Encode(articles)
}
