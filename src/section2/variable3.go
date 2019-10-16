package main

import "fmt"

func main() {
	// 짧은 선언(Go에만 있음)
	// 반드시 함수 안에서만 사용(전역으로는 사용 불가)
	// 선언 후 재할당 하면 에러 발생
	// 특정 메서드 안에 1회성으로 사용하는 것
	// 주로 제한된 범위의 함수 내에서 사용할 경우 코드 가독성을 높일 수 있다.

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	// shortVar1 := 10 // 에러 발생

	fmt.Println("shortVar1: ", shortVar1, "shortVar2: ", shortVar2, "shortVar3: ", shortVar3)

	// Example
	if i := 10; i < 11 {
		fmt.Println("short Variable Test Success!")
	}
}
