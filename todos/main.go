package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func main() {
	connStr := "host=localhost port=5432 dbname=todo user=markopoikkimaki password="

	conn, err := sql.Open("pgx", connStr)

	if err != nil {
		log.Fatalln(fmt.Sprintf("Unable to connect %v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to database!")
	//conn.Query()

	err = conn.Ping()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Cannot ping database %v\n", err))
	}

	log.Println("PING DB BINGO!!")

}
