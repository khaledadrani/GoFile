package main

import (
	"fmt"
	"io"

	// "net/http"
	"os"
)

type bot interface {
	getMessage() string
	speak(string) string
}

type EnglishBot struct {
	helloMessage string
	language     string
}

type FrenchBot struct {
	helloMessage string
	language     string
}

func (b EnglishBot) getMessage() string {
	b.helloMessage = "Hello"
	return b.helloMessage
}

func (b FrenchBot) getMessage() string {
	b.helloMessage = "Bonjour"
	return b.helloMessage
}

func printMessage(b bot) {
	fmt.Println(b.getMessage())
}

func (b FrenchBot) speak(message string) string {
	return b.getMessage() + message
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	surface := s.getArea()
	fmt.Println("Surface ", surface)

}

func main() {

	filePath := ""

	if len(os.Args) > 1 {
		fmt.Println(os.Args)
		filePath = os.Args[1]
	} else {
		fmt.Println("No file was provided")
		os.Exit(1)
	}

	fileData, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, fileData)

	defer fileData.Close()

	// s := square{
	// 	sideLength: 10.0,
	// }

	// t := triangle{
	// 	height: 10.0,
	// 	base:   5.5,
	// }

	// printArea(s)
	// printArea(t)

	// resp, err := http.Get("http://google.com")

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }

	// bs := make([]byte, 9999)

	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	// lw := logWriter{}

	// io.Copy(lw, resp.Body)

	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)

	// if err != nil {
	// 	fmt.Println("Error while getting body: ", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(body)

}
