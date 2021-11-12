package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper finished..")
}

func main() {
	go slowFunc()

	fmt.Println("All finished! BINGO!")

	time.Sleep(time.Second * 3)
}
