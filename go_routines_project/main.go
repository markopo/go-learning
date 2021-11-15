package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var c = make(chan string)

func slowFunc() {
	time.Sleep(time.Second * 2)
	c <- "slow func finished.."
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
	go slowFunc()

	msg := <-c
	fmt.Println("MESSAGE RECEIVED: ", msg)



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
