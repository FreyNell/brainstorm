package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func getUsers(db *sql.DB) ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT id,name,username FROM users")
	if err != nil {
		return nil, fmt.Errorf("users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.ID, &usr.Name, &usr.Username); err != nil {
			return nil, fmt.Errorf("users: %v", err)
		}
		users = append(users, usr)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("users:", err)
	}

	return users, nil
}

func main() {
	var db *sql.DB

	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "db:3306",
		DBName: "iam",
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
	fmt.Println("Connected to db!")

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Users found: %v\n", users)

}
