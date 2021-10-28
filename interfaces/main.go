package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name string
	Rating float64
}

func (m *Movie) summary() string  {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + ", " + r
}


func main() {
	fmt.Println("METHODS!")

	m := Movie{
		Name: "Rambo",
		Rating: 5.138,
	}

	fmt.Println(m.summary())

	fmt.Println("***** INTERFACES!!! *****")
}
