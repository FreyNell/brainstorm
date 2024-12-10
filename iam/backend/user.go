package main

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	password string `json:"password"`
}

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
		return nil, fmt.Errorf("users: %v", err)
	}

	return users, nil
}

func getUserByName(db *sql.DB, name string) ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT id,name,username FROM users WHERE name LIKE ?", "%"+name+"%")
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
		return nil, fmt.Errorf("users: %v", err)
	}

	return users, nil
}
