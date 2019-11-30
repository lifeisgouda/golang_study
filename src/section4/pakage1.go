//패키지(1)
package main

//선언 방법1
/*
import "fmt"
import "os"
*/

//선언 방법2
import (
	"fmt"
	"os"
)

func main() {
	//패키지 : 코드 구조화 및 재사용
	//Go는 패키지로 구성되어 있음
	//서로 다른 패키지간에 서로 import 후 사용
	//패키지 이름 = 디렉터리 이름
	//같은 패키지 내 -> 소스파일 같은 디렉터리 위치
	//네이밍 규칙 : 소문자, 패키지명 = 소스파일명

	var name string

	fmt.Print("이름은? :  ")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi %s\n", name)

}
