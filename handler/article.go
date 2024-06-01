package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"entale_codingtest/domain"
)

func GetArticles(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Get Articles and Media from database
	articles, err := domain.GetArticles(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Marshal articles to JSON
	articlesJSON, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(articlesJSON)
}
