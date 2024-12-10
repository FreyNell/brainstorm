package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var cfg mysql.Config = mysql.Config{
	User:   os.Getenv("DB_USER"),
	Passwd: os.Getenv("DB_PASSWORD"),
	Net:    "tcp",
	Addr:   "db:3306",
	DBName: "iam",
}

func getDB() (*sql.DB, error) {
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return nil, err
	}
	fmt.Println("Connected to db!")

	return db, nil
}
