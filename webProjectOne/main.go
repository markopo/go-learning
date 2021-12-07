package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 0.0
	sum, err := divideValues(float32(x), float32(y))

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("Divide %f / %f = %f", x, y, sum))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Can't divide by zero!")
		return 0, err
	}

	result := x / y
	return result, nil
}


func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting the app at port%s", portNumber))

	if strings.Contains(portNumber, "69")  {
		fmt.Println("It's the dirty port!")
	}

	_ = http.ListenAndServe(portNumber, nil)
}
