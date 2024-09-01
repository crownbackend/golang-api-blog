package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/crownbackend/golang-api-blog/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func InitializeDatabase(database *sql.DB) {
	db = database
}

// send this c *gin.Context to fun use gin
func GetUsers(c *gin.Context) {
	var query string = "SELECT id, email, first_name, last_name, created_at FROM user"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := emailExists(user.Email)

	if err != nil {
		log.Println("Erreur lors de la vérification de l'email :", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la vérification de l'email"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "L'email est déjà utilisé"})

		return
	}

	// Hacher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Erreur lors du hachage du mot de passe :", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du hachage du mot de passe"})
		return
	}
	user.Password = string(hashedPassword)

	// Préparer la requête d'insertion
	query := `INSERT INTO user (email, password, first_name, last_name, created_at) VALUES (?, ?, ?, ?, ?)`
	result, err := db.Exec(query, user.Email, user.Password, user.FirstName, user.LastName, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println("Erreur lors de l'insertion :", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'insertion"})
		return
	}

	// Obtenir l'ID du dernier enregistrement inséré
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Erreur lors de l'obtention du dernier ID :", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'obtention du dernier ID"})
		return
	}

	// Mettre à jour l'ID de l'utilisateur
	user.Id = int(id)

	c.JSON(http.StatusCreated, user)
}

func emailExists(email string) (bool, error) {
	query := `SELECT COUNT(*) FROM user WHERE email = ?`
	var count int
	err := db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
