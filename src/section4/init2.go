//GO 초기화 함수(2)
package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init1 is being executed")
}

func init() {
	fmt.Println("Init2 is being executed")
}

func init() {
	fmt.Println("Init3 is being executed")
}

func main() {
	fmt.Println("Main Source Code!")
}
