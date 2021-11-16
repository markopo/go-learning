package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)



func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func responseTime(url string)  {
	start := time.Now()

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	elapsed := time.Since(start).Seconds()

	fmt.Println("‰s took %v seconds.", url, elapsed)
}

func main() {
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"

	close(messages)

	fmt.Println("Push 2 messages onto Channel with no receivers.")

	time.Sleep(time.Second * 1)

	receiver(messages)

	//urls := make([]string, 5)
	//urls[0] = "https://sr.se"
	//urls[1] = "https://svt.se"
	//urls[2] = "https://dn.se"
	//urls[3] = "https://www.usa.gov"
	//urls[4] = "https://expressen.se"
	//
	//for _, u := range urls {
	//	go responseTime(u)
	//}
	//
	//time.Sleep(time.Second * 5)
}
