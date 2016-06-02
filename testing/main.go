package main

import "fmt"

func Hello() string {
	return "hello"
}

func main() {
	var response = Hello()
	fmt.Println(response)
}
