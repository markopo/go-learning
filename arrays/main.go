package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)

	cheeses[0] = "Cheddar"
	cheeses[1] = "Västerbotten"

	cheeses = append(cheeses, "Camembert", "Hushållsost")

	// delete at index 2
	// cheeses = append(cheeses[:2], cheeses[2+1:]...)

	// delete at index 0
	cheeses = append(cheeses[:0], cheeses[0+1:]...)

	for i, s := range cheeses {
		fmt.Println(i, s)
	}

}
