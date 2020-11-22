package main

import (
	"gin-web-app/config"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	config.SetConfig()

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run()
}
