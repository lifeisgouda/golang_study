package main

import "fmt"

func main() {
	i := 100

	// if - else if example(1)
	if i >= 120 {
		fmt.Println("over 120")
	} else if i >= 100 && i < 120 {
		fmt.Println("over 100 under 120")
	} else if i < 100 && i >= 50 {
		fmt.Println("over 50 under 100")
	} else {
		fmt.Println("under 50")
	}
}
