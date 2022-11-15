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

	err = insertUser(conn, "Musse", "Pigg")
	if err != nil {
		log.Fatal(err)
	}

	err = insertUser(conn, "Kalle", "Anka")
	if err != nil {
		log.Fatal(err)
	}

	err = insertUser(conn, "Arne", "Anka")
	if err != nil {
		log.Fatal(err)
	}

	err = insertUser(conn, "Mimmi", "Anka")
	if err != nil {
		log.Fatal(err)
	}

	err = getUsers(conn)
	if err != nil {
		log.Fatal(err)
	}

}

func insertUser(conn *sql.DB, fName string, lName string) error {

	row := conn.QueryRow("select count(id) as antal from users where first_name = $1 and last_name = $2", fName, lName)

	var antal int
	row.Scan(&antal)

	if antal > 0 {
		fmt.Printf("User %s %s already exists!!\n", fName, lName)
		return nil
	}

	fmt.Println("antal: ", antal)

	sql := "insert into users (first_name, last_name) values($1, $2)"
	_, err := conn.Exec(sql, fName, lName)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Inserted a row!")

	return nil
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
