package main

import "fmt"

func main() {
	const (
		_ = iota // 밑줄로 시작하면 0
		A
		_ // 생략. skip 처리. 출력안됨. 규칙은 그대로 유지.
		C
	)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		_
		PLATINUM
	)

	fmt.Println("Default: ", DEFAULT, "Silver: ", SILVER, "Gold: ", GOLD, "PLATINUM: ", PLATINUM)
}
