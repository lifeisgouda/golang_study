//패키지 접근제어(1)
package main

import (
	"fmt"
	"section4/lib"
)

func main() {
	//패키지 접근제어
	//변수,상수,함수,메서드,구조체 등 식별자
	//대문자 : 패키지 외부 접근 가능
	//소문자 : 패키지 외부 접근 불가(패키지 내에서만 접근 가능)

	fmt.Print("100 보다 큰 수? :  ", lib.CheckNum2(101))
	fmt.Print("1000 보다 큰 수? :  ", lib.checkNum3(1001)) //예외 발생 (접근 불가)

}
