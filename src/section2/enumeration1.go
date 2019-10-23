package main

import "fmt"

func main() {
	// 열거형
	// 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가 시키는 묶음
	const (
		Jan   = 1
		Feb   = 2
		March = 3
		April = 4
		May   = 5
		June  = 6
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(March)
	fmt.Println(April)
	fmt.Println(May)
	fmt.Println(June)
}
