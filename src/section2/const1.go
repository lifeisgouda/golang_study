package main

import "fmt"

func main() {
	/*
		- 상수
			const 사용 초기화, 한번 선언 후 값 변경 금지, 고정된 값 관리용
	*/

	const a string = "Test1" // 선언과 동시에 초기화 해야 함
	const b = "Test2"
	const c int32 = 10 * 10

	/* const d = getHeight()
	에러 발생. getHeight()가 항상 같은 값을 받는다는 보장이 없으므로 */

	const e = 35.6
	const f = false

	fmt.Println("a: ", a, "b: ", b, "c: ", c, "e: ", e, "f: ", f)

}
