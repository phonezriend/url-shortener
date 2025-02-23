package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/teris-io/shortid"
)

var ctx = context.Background()

// Connect Redis
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// POST /shorten
func ShortenURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	// Generate Short ID
	shortID, _ := shortid.Generate()

	// Save in Redis
	err := rdb.Set(ctx, shortID, req.URL, 24*time.Hour).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
		return
	}

	shortURL := "http://localhost:8080/" + shortID
	c.JSON(http.StatusOK, ShortenResponse{ShortURL: shortURL})
}

// GET /:shortID
func ResolveURL(c *gin.Context) {
	shortID := c.Param("shortID")
	originalURL, err := rdb.Get(ctx, shortID).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
	} else {
		c.Redirect(http.StatusMovedPermanently, originalURL)
	}
}