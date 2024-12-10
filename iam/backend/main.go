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
	defer db.Close()

	name := c.Query("name")
	if name != "" {
		users, err = getUserByName(db, name)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		users, err = getUsers(db)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Users found: %v\n", users)
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getAllUsersHandler)

	router.Run("0.0.0.0:" + os.Getenv("SERVERPORT"))
}
