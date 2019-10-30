package main

import "fmt"

func main() {
	// basic
	for i := 0; i < 5; i++ {
		fmt.Println("ex1: ", i)
	}

	/* error 1
	for i := 0; i <5; i++
	{
		fmt.Println("ex1: ", i)
	}

	//error 2
	for i := 0; i < 5; i++
		fmt.Println("ex1: ", i)

	*/

	/* basic 2: Infinite loop!
	for {
		fmt.Println("ex2: Hello, Go!")
		fmt.Println("ex2: Infinite loop!")
	}
	*/

	// basic 3: Range
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
}
