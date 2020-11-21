package main

import "gin-web-app/handlers"

// initializeRoutes initializing all available application routes
func initializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)

	router.GET("/article/view/:article_id", handlers.ShowArticlePage)
}
