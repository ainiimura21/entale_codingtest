package domain

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

func InsertArticleFromAPI(db *sql.DB) error {
	// Connect to API
	response, err := http.Get("https://gist.githubusercontent.com/gotokatsuya/cc78c04d3af15ebe43afe5ad970bc334/raw/dc39bacb834105c81497ba08940be5432ed69848/articles.json")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Unmarshal JSON
	var articles []Article
	if err := json.Unmarshal(responseData, &articles); err != nil {
		return err
	}
	// Add data to Articles and Medias table
	for i := 0; i < len(articles); i++ {
		if err := InsertArticle(db, articles[i]); err != nil {
			return err
		}
		for j := 0; j < len(articles[i].Medias); j++ {
			if err := InsertMedia(db, articles[i].Medias[j], articles[i].ID); err != nil {
				return err
			}

		}
	}
	return err
}
