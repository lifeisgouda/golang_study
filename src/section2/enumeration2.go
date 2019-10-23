package main

import "fmt"

func main() {
	const (
		Jan = iota + 1
		Feb
		March
		April
		May
		June
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(March)
	fmt.Println(April)
	fmt.Println(May)
	fmt.Println(June)
}
