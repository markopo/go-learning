package main

import(
	"fmt"
)

type Movie struct {
	Name string
	Rating float32
}

type Superhero struct {
	Name string
	Age int
	Address
}

type Address struct {
	Number int
	Street string
	City string
}

func main() {
	movie := Movie {
		Name: "Blade Runner",
		Rating: 10,
	}

	fmt.Println("%+v\n", movie)

	fmt.Println(movie.Name, movie.Rating)

	batman := Superhero{
		Name: "Batman",
		Age: 50,
		Address: Address{
			Number: 12,
			Street: "Bat Cave",
			City: "Gotham City",
		},
	}

	fmt.Println("%+v\n", batman)
}
