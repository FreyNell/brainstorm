package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getAllUsersHandler(c *gin.Context) {
	var users []User

	db, err := getDB()
	if err != nil {
		log.Fatal(err)
	}

	users, err = getUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Users found: %v\n", users)
	c.IndentedJSON(http.StatusOK, users)
}

func getUsersByNameHandler(c *gin.Context) {
	var users []User
	db, err := getDB()

	if err != nil {
		log.Fatal(err)
	}

	users, err = getUserByName(db, "Sarah")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User found: %v\n", users)
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getAllUsersHandler)

	router.Run("0.0.0.0:" + os.Getenv("SERVERPORT"))
}
