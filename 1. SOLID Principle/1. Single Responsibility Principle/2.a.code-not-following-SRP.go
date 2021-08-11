package main

import (
	"database/sql"
	"fmt"
)

type Student struct {
	id      string
	age     int
	address string
}

func (st Student) save() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "your-password"
		dbname   = "srp"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := fmt.Sprintf(`INSERT INTO students (id, age, address)
                     VALUES ($1, $2, $3) 
                     RETURNING id`)
	id := 0
	err = db.QueryRow(sqlStatement, 100, 200, "India").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

// This code is having tight coupling between the Student and DB. So in future if we want change the DB then it will lead to change in Student functionality
