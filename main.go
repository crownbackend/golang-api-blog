package main

import (
	"fmt"

	"github.com/crownbackend/golang-api-blog/models"
)

func main() {
	user := models.User{
		Id:       1,
		Email:    "john@doe.fr",
		Password: "1234",
	}
	fmt.Printf("User: %+v \n", user)
}
