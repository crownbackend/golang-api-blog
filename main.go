package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/crownbackend/golang-api-blog/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// connect database
	connectDb()

	// config gin
	r := gin.Default()

	// Configuration CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                            // Changez ces valeurs selon vos besoins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}, // Méthodes HTTP autorisées
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},          // En-têtes autorisés
		ExposeHeaders:    []string{"Content-Length"},                                   // En-têtes exposés
		AllowCredentials: true,                                                         // Permet l'envoi des cookies
		MaxAge:           12 * time.Hour,                                               // Durée de validité des pré-vols CORS
	}))

	// list of routes
	r.GET("/", home)
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)

	// run server
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
