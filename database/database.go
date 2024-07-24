package database

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/main/internal/handler"
	_ "github.com/lib/pq"
)

func Initdb() {

	const (
		host     = "ep-red-pine-a2fvgpqs.eu-central-1.pg.koyeb.app"
		port     = 5432
		user     = "koyeb-adm"
		password = "Txk1LDAUcZB9"
		dbname   = "koyebdb"
	)

	// Construct the connection string
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	// Attempt to connect to the database
	var err error
	handler.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = handler.DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Database connection established")

	// 	createTable := `CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	name TEXT NOT NULL,
	// 	orders TEXT[] NOT NULL
	// 	)`

	//		_, err = handler.DB.Exec(createTable)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Println("Table created successfully")
	//	}
}
