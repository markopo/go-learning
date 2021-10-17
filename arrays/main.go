package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)

	cheeses[0] = "Cheddar"
	cheeses[1] = "Västerbotten"

	cheeses = append(cheeses, "Camembert", "Hushållsost")

	var copyCheeses = make([]string, 4)

	copy(copyCheeses, cheeses)

	// delete at index 2
	// cheeses = append(cheeses[:2], cheeses[2+1:]...)

	// delete at index 0
	cheeses = append(cheeses[:0], cheeses[0+1:]...)

	fmt.Println("CHEESES ****")
	for i, s := range cheeses {
		fmt.Println(i, s)
	}

	fmt.Println("COPY CHEESES ****")
	for i, s := range copyCheeses {
		fmt.Println(i, s)
	}



}
