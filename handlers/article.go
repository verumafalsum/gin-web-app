package handlers

import (
	"gin-web-app/helpers"
	"gin-web-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ShowIndexPage rendering index.html with title "Home Page" and payload of all articles
func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	helpers.Render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

// ShowArticlePage rendering article.html with article title and payload of article.
// The article to be rendered depends on param "article_id" in context
func ShowArticlePage(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := models.GetArticleByID(articleID); err == nil {
			helpers.Render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
