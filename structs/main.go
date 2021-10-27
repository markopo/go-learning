package main

import(
	"fmt"
	"reflect"
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

	fmt.Println(reflect.TypeOf(movie))

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

	not_real_batman := batman

	not_real_batman.Name = "False Batman"

	fmt.Println(not_real_batman.Name, batman.Name)
	
	// pointer
	copy_batman := &batman

	copy_batman.Name = "Batman (Bruce Wayne)"

	fmt.Println(copy_batman.Name, batman.Name)

	fmt.Println(reflect.TypeOf(batman))

	fmt.Println("%+v\n", batman)
}
