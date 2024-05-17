package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	Medias []Media `json:"medias"`
}

type Media struct {
	ID          int    `json:"id"`
	ContentURL  string `json:"contentUrl"`
	ContentType string `json:"contentType"`
}

func main() {
	fmt.Println(os.Getenv("MYSQL_ROOT_PASSWORD"))
	pswd := os.Getenv("MYSQL_ROOT_PASSWORD")

	db, err := sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/entale")
	if err != nil {
		fmt.Println("Not connected to the database")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error in connection")
		panic(err.Error())
	}
	fmt.Println("Connected to the database")

	// Fetch articles from the URL
	response, err := http.Get("https://gist.githubusercontent.com/gotokatsuya/cc78c04d3af15ebe43afe5ad970bc334/raw/dc39bacb834105c81497ba08940be5432ed69848/articles.json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	fmt.Println("Connected to Database")

}
