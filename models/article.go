package models

import "errors"

// Article is the db table where
// Title is article title
// Content is article content
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ArticleList is hardcoded list of Article
// used to getting blog articles
var ArticleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// GetAllArticles returns all articles
func GetAllArticles() []Article {
	return ArticleList
}

// GetArticleByID returns article by it's id
func GetArticleByID(id int) (*Article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
