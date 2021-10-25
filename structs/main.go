package main

import(
	"fmt"
)

type Movie struct {
	Name string
	Rating float32
}

func main() {
	movie := Movie {
		Name: "Blade Runner",
		Rating: 10,
	}

	fmt.Println(movie.Name, movie.Rating)
}
