// GO 특징 (2)
package main

import "fmt"

func main() {
	// ex 1
	for i := 0; i <= 10; i++ {
		// fmt.Print("ex1: ", i); fmt.Println(i) // 두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론 사용(권장하지 않음)
		fmt.Print("ex1: ", i)
		fmt.Println("ex1: ", i)

	}

	// ex 2
	for j := 10; j >= 0; j-- { // 괄호 위치 중요
		fmt.Println("ex2 : ", j)
	}

}
