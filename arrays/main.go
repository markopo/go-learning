package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)

	cheeses[0] = "Cheddar"
	cheeses[1] = "Västerbotten"

	for i, s := range cheeses {
		fmt.Println(i, s)
	}

}
