package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"

	"entale_codingtest/domain"
	"entale_codingtest/handler"
)

func main() {
	// Connect to mysql database
	pswd := os.Getenv("MYSQL_ROOT_PASSWORD") // Hide password
	db, err := sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/entale")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Add Table to Database if does not exist
	if err := domain.CreateArticleTable(db); err != nil {
		log.Fatal(err)
	}
	if err := domain.CreateMediaTable(db); err != nil {
		log.Fatal(err)
	}
	// Insert article data to table
	if err := domain.InsertArticleFromAPI(db); err != nil {
		log.Fatal(err)
	}

	// Create new router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handler.GetArticles(db, w, r)
	})
	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
