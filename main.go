package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", home)

	r.Run("localhost:8000")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "testo"})
}
