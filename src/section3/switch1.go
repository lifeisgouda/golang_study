package main

import "fmt"

func main() {
	// ex1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "은/는 음수")
	case a == 0:
		fmt.Println(a, "은/는 0")
	case a > 0:
		fmt.Println(a, "은/는 양수")
	}

	// ex2: Golang에서 더 선호하는 스타일
	// b의 범위가 switch 안으로 제한 -> clean code
	switch b := 27; {
	case b < 0:
		fmt.Println(a, "은/는 음수")
	case b == 0:
		fmt.Println(a, "은/는 0")
	case b > 0:
		fmt.Println(a, "은/는 양수")
	}

	// ex3
	switch c := "go"; c {
	case "go":
		fmt.Println("Go")
	case "Java":
		fmt.Println("Java")
	default:
		fmt.Println("일치하는 값 없음.")
	}

	// ex4
	switch c := "go"; c + "lang" { // 연산자로 계산 가능
	case "golang":
		fmt.Println("Go")
	case "java":
		fmt.Println("Java")
	default:
		fmt.Println("일치하는 값 없음.")
	}

	// ex 5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i < j")
	case i == j:
		fmt.Println("i == j")
	case i > j:
		fmt.Println("i > j")
	}
}
