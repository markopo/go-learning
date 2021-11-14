package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper finished..")
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
	//go slowFunc()
	//
	//fmt.Println("All finished! BINGO!")
	//
	//time.Sleep(time.Second * 3)

	urls := make([]string, 5)
	urls[0] = "https://sr.se"
	urls[1] = "https://svt.se"
	urls[2] = "https://dn.se"
	urls[3] = "https://www.usa.gov"
	urls[4] = "https://expressen.se"

	for _, u := range urls {
		go responseTime(u)
	}

	time.Sleep(time.Second * 5)
}
