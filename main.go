package main

import (
	"github.com/gin-gonic/gin"
	"url-shortener/handler"
)

func main() {
	r := gin.Default()

	r.POST("/shorten", handler.ShortenURL)
	r.POST("/:shortID", handler.ResolveURL)

	r.Run(":8080") // Port 8080
}