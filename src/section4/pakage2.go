//패키지(2)
package main

import (
	"fmt"
	"section4/lib"
)

func main() {
	//패키지 종류
	//1.메인 프로그램(main)
	//2.다른 패키지에서 호출 가능한 라이브러리

	fmt.Print("10 보다 큰 수? :  ", lib.CheckNum(15))

}
