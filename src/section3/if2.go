package main

import "fmt"

func main() {
	var a int = 50
	b := 70 // 함수 안에서만 사용할 짧은 선언

	// ex1
	if a >= 65 {
		fmt.Println("over 65")
	} else {
		fmt.Println("under 65")
	}

	// ex2
	if b >= 70 {
		fmt.Println("over 70")
	} else {
		fmt.Println("under 70")
	}
}
