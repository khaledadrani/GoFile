package main

import (
	"fmt"
	"net/http"
	// "time"
)

func checkStatus(url string, c chan string) {
	result, err := http.Get(url)

	if err != nil {
		fmt.Println(url + " might be down!")
		c <- url
		return
	}

	fmt.Println(url + " OK!")
	c <- url

	defer result.Body.Close()
}

func main() {

	// websites := []string{
	// 	"http://google.com",
	// 	"http://facebook.com",
	// 	"http://stackoverflow.com",
	// 	"http://golang.com",
	// }

	// c := make(chan string)

	// for _, url := range websites {
	// 	go checkStatus(url, c)
	// }

	// for i := 0; i < len(websites); i++ {
	// 	go checkStatus(<-c, c)
	// }

	// for l := range c {
	// 	go func(l string) {
	// 		time.Sleep(6 * time.Second)
	// 		checkStatus(l, c)
	// 	}(l)
	// }

	c := make(chan string)

	for i := 0; i < 4; i++ {
		go printString("Hello there!", c)
	}

	for {
		fmt.Println(<-c)
	}
}

func printString(s string, c chan string) {
	fmt.Println(s)
	c <- "Done printing."
}
