package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Define the Structs

type Article struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	Date   string  `json:"publishedAt"`
	Medias []Media `json:"medias"`
}

type Media struct {
	ID          int    `json:"id"`
	ContentURL  string `json:"contentUrl"`
	ContentType string `json:"contentType"`
}

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
	if err := createArticleTable(db); err != nil {
		log.Fatal(err)
	}
	if err := createMediaTable(db); err != nil {
		log.Fatal(err)
	}

	// Connect to API
	response, err := http.Get("https://gist.githubusercontent.com/gotokatsuya/cc78c04d3af15ebe43afe5ad970bc334/raw/dc39bacb834105c81497ba08940be5432ed69848/articles.json")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal JSON
	var articles []Article
	if err := json.Unmarshal(responseData, &articles); err != nil {
		log.Fatal(err)
	}
	// Add data to Articles and Medias table
	for i := 0; i < len(articles); i++ {
		if err := insertArticle(db, articles[i]); err != nil {
			log.Println(err)
		}
		for j := 0; j < len(articles[i].Medias); j++ {
			if err := insertMedia(db, articles[i].Medias[j], articles[i].ID); err != nil {
				log.Println(err)
			}

		}
	}
	// Get Articles and Media from database
	articles, err = fetchArticles(db)
	if err != nil {
		log.Fatal(err)
	}

	// Print database
	printArticlesAsJSON(articles)

}
