package main

import "fmt"

func main() {
	// ex1
	a := 30 / 15
	switch a {
	case 2, 4, 6: // i가 2|4|6 인 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "는 홀수")
	}
	// ex 2
	switch e := "Go"; e {
	case "Java":
		fmt.Println("Java")
		fallthrough
	case "Go":
		fmt.Println("Go")
		fallthrough
	case "Python":
		fmt.Println("Python")
		fallthrough
	case "Ruby":
		fmt.Println("Ruby")
		fallthrough
	case "C":
		fmt.Println("C")
	}
}
