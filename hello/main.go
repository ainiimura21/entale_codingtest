package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

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
	pswd := os.Getenv("MYSQL_ROOT_PASSWORD")
	db, err := sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/entale")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := createArticleTable(db); err != nil {
		log.Fatal(err)
	}

	response, err := http.Get("https://gist.githubusercontent.com/gotokatsuya/cc78c04d3af15ebe43afe5ad970bc334/raw/dc39bacb834105c81497ba08940be5432ed69848/articles.json")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var articles []Article
	if err := json.Unmarshal(responseData, &articles); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(articles)

	for i := 0; i < len(articles); i++ {
		fmt.Println(len(articles[i].Medias))
		for j := 0; j < len(articles[i].Medias); j++ {

			fmt.Println(articles[i].Medias[j].ContentURL)
		}

		fmt.Println(articles[i].Title)
	}

	// for _, article := range articles {
	// 	if err := insertArticle(db, article); err != nil {
	// 		log.Println(err)
	// 	}
	// }
}

func createArticleTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS articles (
		id INT PRIMARY KEY,
		title VARCHAR(100),
		body TEXT,
		date VARCHAR(50)
	);`

	_, err := db.Exec(query)
	return err
}

func insertArticle(db *sql.DB, article Article) error {
	query := `INSERT INTO articles (id, title, body, date) VALUES (?, ?, ?, ?)`
	insert, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	_, err = insert.Exec(article.ID, article.Title, article.Body, article.Date)
	if err != nil {
		return err
	}

	fmt.Println("Article inserted successfully")
	return nil
}
