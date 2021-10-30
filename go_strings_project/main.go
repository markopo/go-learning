package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	var buffer bytes.Buffer
	
	for i := 0; i < 500; i++ {
		buffer.WriteString("go! ")
	}

	fmt.Println(buffer.String())
}
