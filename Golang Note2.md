# Golang 노트 2. Section4~6

# Section 4. Go 패키지 기초

## 패키지 관련 정리

- 패키지 : 코드 구조화(모듈화) 및 재사용
- **응집도는 높게, 결합도는 느슨하게 -> 클린코드**
- **패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작성할 것을 권고**
- Go는 패키지로 구성되어 있음
- 서로 다른 패키지간에 서로 import 후 사용
- **패키지 이름 = 디렉터리 이름**
- 같은 패키지 내 -> 소스파일 같은 디렉터리 위치
- 네이밍 규칙 : 소문자, 패키지명 = 소스파일명

### package1.go

```go
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
	var name string

	fmt.Print("이름은? :  ")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi %s\n", name)

}
```



## 접근제어 및 Alias



## 초기 메서드

