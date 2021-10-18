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

	fmt.Println("MAPS ****")

	// players score
	var players = make(map[string]int)

	players["marko"] = 23
	players["alex"] = 12
	players["alice"] = 34
	players["linda"] = 11

	// fmt.Println(players)

	for i, s := range players {
		fmt.Println("key: ", i)
		fmt.Println("value: ", s)
	}

	delete(players, "marko")

	fmt.Println(len(players))

	

}
