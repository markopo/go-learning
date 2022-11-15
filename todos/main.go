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

	err = getUsers(conn)

	if err != nil {
		log.Fatal(err)
	}

}

func getUsers(conn *sql.DB) error {
	sql := "select id, first_name, last_name from users"
	rows, err := conn.Query(sql)

	if err != nil {
		log.Println(err)
		return err
	}

	var first_name, last_name string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &first_name, &last_name)

		if err != nil {
			log.Println(err)
			return err
		}

		fmt.Println("User: ", id, first_name, last_name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows ", err)
	}

	return nil
}
