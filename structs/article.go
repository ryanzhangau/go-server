package structs

//Article - article struct
type Article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

// Articles - array of articles
type Articles []Article