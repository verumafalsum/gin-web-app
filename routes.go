package main

import "gin-web-app/handlers"

func initializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)

	router.GET("/article/view/:article_id", handlers.GetArticle)
}
