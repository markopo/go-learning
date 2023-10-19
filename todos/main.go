package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
)

var connStr = "host=localhost port=5432 dbname=todo user=markopoikkimaki password="

func main() {

	//conn, err := sql.Open("pgx", connStr)
	//
	//if err != nil {
	//	log.Fatalln(fmt.Sprintf("Unable to connect %v\n", err))
	//}
	//defer conn.Close()
	//
	//log.Println("Connected to database!")

	var gb GoogleBooksStruct

	resp, err := http.Get("https://www.googleapis.com/books/v1/volumes?q=Golang")

	if err != nil {
		log.Fatalln(fmt.Sprintf("Unable to request %v\n", err))
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&gb)

	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range gb.Items {
		fmt.Println("*****")
		fmt.Println(item.VolumeInfo.Title)
		fmt.Println("*****")
		fmt.Println(item.VolumeInfo.PageCount)
		fmt.Println("*****")
		fmt.Println(item.VolumeInfo.Description)
		fmt.Println("*****")
	}

	//fmt.Printf("%+v\n\n", gb)

}
