package main

import (
	"fmt"
	"time"
)

func ping1(c chan string) {
	time.Sleep(time.Millisecond * 200)
	c <- "ping on channel 1"
}

func ping2(c chan string) {
	time.Sleep(time.Millisecond * 100)
	c <- "ping on channel 2"
}

func ping3(c chan string) {
	time.Sleep(time.Millisecond * 300)
	c <- "ping on channel 3"
}

func main() {
	quit := false
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	for i := 0; i < 10; i++ {
		if i % 2 == 1 {
			go ping1(channel1)
		}

		if i % 2 == 0 {
			go ping2(channel2)
		}

		if i % 3 == 0 {
			go ping3(channel3)
		}

	}


	for {
		select {
			case msg1 := <-channel1:
				fmt.Println("#1: received", msg1)
			case msg2 := <-channel2:
				fmt.Println("#2: received", msg2)
			case msg3 := <-channel3:
				fmt.Println("#3: received", msg3)
			case <-time.After(5 * time.Second):
				fmt.Println("no messages, giving up.")
				quit = true
			}

		if quit == true {
			break
		}
	}

}
