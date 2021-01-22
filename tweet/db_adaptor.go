package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password" // change for specific postgres
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	insertFirstUser(db)

	fmt.Println("Successfully inserted!")
}

func insertFirstUser(db *sql.DB) {
	sqlStatement := `
    INSERT INTO public."Users" (user_id, email, date_created, date_modified, username, blurb, birthdate, first_name, last_name)
    VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err := db.Exec(sqlStatement, 1, "michael.wang.co@gmail.com", time.Now().Format("2006-01-02"), nil, "mwang", "Michael is one of the first users", time.Date(1992, 01, 13, 12, 30, 0, 0, time.UTC), "Michael", "Wang")
	if err != nil {
		panic(err)
	}
}
