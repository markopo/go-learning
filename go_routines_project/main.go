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
	slowFunc()

	fmt.Println("All finished! BINGO!")
}
