package domain

import (
	"database/sql"
)

// Create Tables
func CreateArticleTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS articles (
		id INT PRIMARY KEY,
		title VARCHAR(100),
		body TEXT,
		date VARCHAR(50)
	);`

	_, err := db.Exec(query)
	return err
}

func CreateMediaTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS medias (
		id INT PRIMARY KEY,
		content_url VARCHAR(100),
		content_type VARCHAR(50),
		article_id INT,
		FOREIGN KEY (article_id) REFERENCES articles(id)
	);`
	_, err := db.Exec(query)
	return err
}

// Insert Data Into Table

func InsertArticle(db *sql.DB, article Article) error {
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

	return nil
}

func InsertMedia(db *sql.DB, media Media, articleID int) error {
	query := `INSERT INTO medias (id, content_url, content_type, article_id) VALUES (?, ?, ?, ?)`
	insert, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	_, err = insert.Exec(media.ID, media.ContentURL, media.ContentType, articleID)
	if err != nil {
		return err
	}
	return nil
}

func GetArticles(db *sql.DB) ([]Article, error) {
	// Query articles and media from database
	rows, err := db.Query("SELECT articles.id, articles.title, articles.body, articles.date, medias.id, medias.content_url, medias.content_type FROM articles LEFT JOIN medias ON articles.id = medias.article_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Map to store articles by ID
	articlesMap := make(map[int]Article)

	for rows.Next() {
		var articleID int
		var article Article
		var media Media
		err := rows.Scan(&articleID, &article.Title, &article.Body, &article.Date, &media.ID, &media.ContentURL, &media.ContentType)
		if err != nil {
			return nil, err
		}

		if _, ok := articlesMap[articleID]; !ok {
			article.ID = articleID
			article.Medias = append(article.Medias, media)
			articlesMap[articleID] = article
		} else {
			updatedArticle := articlesMap[articleID]
			updatedArticle.Medias = append(updatedArticle.Medias, media)
			articlesMap[articleID] = updatedArticle
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var result []Article
	for _, article := range articlesMap {
		result = append(result, article)
	}

	return result, nil
}
