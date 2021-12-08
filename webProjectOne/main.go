package main

import (
	"fmt"
	"net/http"
	"strings"
)

const portNumber = ":6969"





func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the app at port%s", portNumber))

	if strings.Contains(portNumber, "69")  {
		fmt.Println("It's the dirty port!")
	}

	_ = http.ListenAndServe(portNumber, nil)
}
