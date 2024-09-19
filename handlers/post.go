package handlers

import (
	"log"
	"net/http"

	"github.com/crownbackend/golang-api-blog/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	query := "SELECT * FROM post"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}