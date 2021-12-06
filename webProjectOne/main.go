package main

import (
	"fmt"
	"net/http"
	"strings"
)

const portNumber = ":6969"

// Home Page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// About Page
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("About Page, 2 + 2 = %d", sum))
}

// Addition
func addValues(x, y int) int {
	return x + y
}


func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the app at port%s", portNumber))

	if strings.Contains(portNumber, "69")  {
		fmt.Println("It's the dirty port!")
	}

	_ = http.ListenAndServe(portNumber, nil)
}
