package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Hello")

	var buffer bytes.Buffer

	for i := 0; i < 500; i++ {
		buffer.WriteString("go! ")
	}

	str := buffer.String()

	fmt.Println(str)

	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.Contains(str, "go!"))
	fmt.Println(strings.Repeat("Hello! \n", 10))

	err := errors.New("Ooops!")

	if err != nil {
		fmt.Println(err)
	}

   txt, err := ioutil.ReadFile("hello.txt")

   if err != nil {
	   fmt.Println("Could not read file!")
   }

   if txt != nil {
	   fmt.Println(string(txt))
   }

}
