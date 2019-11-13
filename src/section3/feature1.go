// GO 특징 (1)
package main

import "fmt"

func main() {
	sum, i := 0, 0

	for i <= 100 {
		// sum += i++ : 반환값 없고, 예외 발생
		sum += i
		i++
		// ++i: 예외 발생 (전위 증감)
	}
	fmt.Println("ex1": sum)
}