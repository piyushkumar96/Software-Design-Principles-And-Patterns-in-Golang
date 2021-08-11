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

type DBConnection struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func (dbc DBConnection) Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dbc.host, dbc.port, dbc.user, dbc.password, dbc.dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}

type DBOperation struct {
}

func (db DBOperation) AddRecord() {
	dbc := &DBConnection{
		host:     "localhost",
		port:     5432,
		user:     "postgres",
		password: "your-password",
		dbname:   "srp_demo",
	}
	dbConnection := dbc.Connect()

	sqlStatement := fmt.Sprintf(`INSERT INTO students (id, age, address)
                     VALUES ($1, $2, $3) 
                     RETURNING id`)
	id := 100
	err := dbConnection.QueryRow(sqlStatement, id, 200, "India")
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

// This code is having loose coupling between the Student and DB. So in future if we want to change the DB then it will not lead to change in Student functionality
