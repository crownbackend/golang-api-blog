package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/crownbackend/golang-api-blog/handlers"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	connectDb()
	r := gin.Default()
	r.GET("/", home)
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)
	r.Run("localhost:8000")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "testo"})
}

func connectDb() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3308",
		DBName: "blog",
	}

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	handlers.InitializeDatabase(db)
	fmt.Println("Connected!")
}
