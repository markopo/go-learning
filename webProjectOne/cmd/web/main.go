package main

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/markopo/go-learning/pkg/handlers"
)

const portNumber = ":6969"





func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the app at port%s", portNumber))

	if strings.Contains(portNumber, "69")  {
		fmt.Println("It's the dirty port!")
	}

	_ = http.ListenAndServe(portNumber, nil)
}
