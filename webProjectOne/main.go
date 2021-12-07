package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

const portNumber = ":6969"

// Home Page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About Page
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
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
