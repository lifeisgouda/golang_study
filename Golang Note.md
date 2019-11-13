# Golang 노트

## Go 설치
### 다운로드
- https://golang.org/dl/



### Go 폴더 만들기
- `bin` / `pkg` / `scr` 폴더 생성
	- `bin` : insatll 하면 bin 폴더에 실행 가능한 파일들(build file)이 생성 
	- `pkg` : 패키지 파일 저장되는 곳
	- `src` : 실질적인 프로젝트 소스 파일이 저장되는 곳

- Go engine이 설치된 곳: `cd /usr/local/go` 
```powershell
jane@Janeui-MacBookPro  /usr/local/go > ls -l
total 352
-rw-r--r--    1 root  wheel  55389  9 26 03:52 AUTHORS
-rw-r--r--    1 root  wheel   1339  9 26 03:52 CONTRIBUTING.md
-rw-r--r--    1 root  wheel  84339  9 26 03:52 CONTRIBUTORS
-rw-r--r--    1 root  wheel   1479  9 26 03:52 LICENSE
-rw-r--r--    1 root  wheel   1303  9 26 03:52 PATENTS
-rw-r--r--    1 root  wheel   1607  9 26 03:52 README.md
-rw-r--r--    1 root  wheel    397  9 26 03:52 SECURITY.md
-rw-r--r--    1 root  wheel      8  9 26 03:52 VERSION
drwxr-xr-x   19 root  wheel    608  9 26 03:55 api
drwxr-xr-x    4 root  wheel    128  9 26 03:55 bin
drwxr-xr-x   50 root  wheel   1600  9 26 03:55 doc
-rw-r--r--    1 root  wheel   5686  9 26 03:52 favicon.ico
drwxr-xr-x    3 root  wheel     96  9 26 03:55 lib
drwxr-xr-x   16 root  wheel    512  9 26 03:55 misc
drwxr-xr-x    6 root  wheel    192  9 26 03:55 pkg
-rw-r--r--    1 root  wheel     26  9 26 03:52 robots.txt
drwxr-xr-x   71 root  wheel   2272  9 26 03:55 src
drwxr-xr-x  327 root  wheel  10464  9 26 03:55 test
```

- pkg 파일을 찾는데 작업 중인 폴더의 pkg에 없으면 엔진이 설치된 곳 pkg에 가서 찾는다.
```powershell
jane@Janeui-MacBookPro  /usr/local/go/pkg  ls
darwin_amd64      darwin_amd64_race include           tool
```

### `env` 설정
```powershell
$ go env
...
GOBIN=""
GOPATH="/Users/jane/dev/go"
GOROOT="/usr/local/go"
export PATH="GOPATH:$GOBIN"
...
```


- `go path` 에 진행할 프로젝트의 경로를 알려줘야 한다.
- `GOPATH` 는 프로젝트의 경로이다.
- `GOBIN` 은 프로젝트 경로 하위의 `bin` 폴더이다.
- `bin` 폴더는 실제 컴파일이 완료되어 실행 가능한 결과물이 저장되는 곳
```powershell
$ vim ~/.bash_profile
-----
# HOME
export HOME="/Users/jane/"

# GO ENV
export GOPATH=$HOME/dev/go
export GOBIN=$GOPATH/bin
-----

$ source ~/.bash_profile
```

- `.zsh` 설정
```powershell
$ vi ~/.zshrc
-----
export GOPATH=$HOME/dev/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH
```

https://blog.outsider.ne.kr/1349




## Test Code

```go
package main
 
import "fmt"
 
func main() {    // main함수: go의 시작점
    fmt.Println("Hello world!") 
}
```



- 실행

```powershell
$ go run helloworld.go
```



## Go Build, Run, Install 의 차이점

```powershell
$ go run helloworld.go    # 컴파일 없이 실행 (테스트 할때 실행)
$ go build helloworld.go    # 해당 파일 있는 폴더에 실행 가능한 바이너리 파일로 만들어줌
$ go install    # bin 폴더에 어플리케이션 생성 됨(프로젝트 완성되었을 때 씀)
```



### Go source code 💡

*https://golang.org/cmd/go/

```powershell
$ go <command> [arguments]

bug         start a bug report
build       compile packages and dependencies
clean       remove object files and cached files
doc         show documentation for package or symbol
env         print Go environment information
fix         update packages to use new APIs
fmt         gofmt (reformat) package sources
generate    generate Go files by processing source
get         add dependencies to current module and install them
install     compile and install packages and dependencies
list        list packages or modules
mod         module maintenance
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
vet         report likely mistakes in packages
```



# GO 장점 및 특징

- 간결한 문법 및 단순함
  - Go로 만들어진 코드를 보면 마치 한사람이 작성한 것 같은 느낌이 든다
    - -> GO FMT라는 라이브러리를 통해서 코드의 정리 등을 지원하기 때문
- 병행 프로그래밍 지원 *핵심
  - 여러 큰 기업에서 쓰는 이유 중 하나
  - 타 언어에서 쓰레드를 선언 해서 멀티 프로그래밍을 하는데, GO는 Go Rutine이라는 것을 통해서 내부적으로 처리. 
- 정적 타입 및 동적 실행
  - 컴파일을 해서 exe 파일 만들어서 실행
- 간편한 협업 지원
  - 필요 라이브러리 표준화 되어 있어서 github 통해서 다운로드 하여 사용 가능
- 컴파일 및 실행속도 빠름
- 제네릭 및 예외 처리 미지원
  - 자바 등의 제네릭, 예외처리 지원 안하지만 코드가 간결해짐.
  - 예외처리 없는 대신 try-catch, Finally, ERROR라는 타입 지원
- 컨벤션 통일
  - GO FMT 통해서 컨벤션이 통일 되어 있음



# 변수 및 상수

- 변수 선언 및 사용법: `var`
- 상수 선언 및 사용법: `const`

## 변수

### 변수 선언 (variable1.go)

```go
package main

import "fmt"

func main() {
	// 변수는 기본 초기화 과정 필요
	// 정수 타입: 0, 실수(소수점): 0.0, 문자열: "", Boolean: true, false

	// var a int32     // 선언만 하고 초기화 안하면 에러 발생
	// var b string
	// var c, d, e int32
	// var f, g, h int32 = 1, 2, 3
	// var i float32 = 11.4
	// var j string = "Hi, Golang!"
	// var m = true

	var k = 4.75

	fmt.Println("k: ", k)

}
```



### 변수 여러 개 선언 (variable2.go)

```go
package main

import "fmt"

func main() {
	// 여러 개 선언: 가독성이 좋아짐
	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)

	height = 250
	weight = 350.56
	isRunning = true

	fmt.Println("name: ", name, "height : ", height, "weight : ", weight, "isRunning : ", isRunning)
}
```



### 짧은 선언 (variable3.go)

```go
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
```



- 재할당 하면 아래와 같이 에러 발생

```go
package main

import "fmt"

func main() {
	shortVar1 := 3
	shortVar1 := 10

	fmt.Println("shortVar1: ", shortVar1)
}

/*
-----------------------ERROR MESSAGE-----------------------------------
# command-line-arguments
./variable3.go:7:12: no new variables on left side of :=
*/
```



## 상수

### 상수 선언 (const1.go)

```go
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
```



### 상수 여러개 선언 (const2.go)

```go
package main

import "fmt"

func main() {
	const a, b int = 10, 90
	const c, d, e = true, 0.84, "test"
	const (
		x, y int16 = 50, 90
		i, k       = "Data", 7776
	)

	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y, i, k)

}
```



# 열거형 Enumeration 🚩

## `Iota`

- java의 `enum class` 같은 것

- 상수를 연속되는 숫자로 나열하거나, 연속되는 규칙에 의해서 계산을 하여 선언 후 비즈니스 로직에서 변수를 계산, 표현 할때 사용

```go
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
```



```go
package main

import "fmt"

func main() {
	const (
		A = iota
		B
		C
	)

	fmt.Println(A, B, C)
}
----------
0, 1, 2
```



```go
package main

import "fmt"

func main() {
	const (
		A = iota * 10
		B
		C
	)

	fmt.Println(A, B, C)
}
-----------
0, 10, 20
```

```go
package main

import "fmt"

func main() {
	const (
		Jan   = iota + 1
		Feb
		March
		April
		May
		June 
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(March)
	fmt.Println(April)
	fmt.Println(May)
	fmt.Println(June)
}
-----------
1, 2, 3, 4, 5, 6
```



## Under bar

```go
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
-------------------
./enumeration3.go:21:66: undefined: GOLD
```



# Go 제어문 및 반복문

- IF / SWITCH / FOR



## IF

- 반드시 Boolean으로 명시적으로 정확하게 검사
- 1, 0 검사 불가. 자동 형변환 없음.
- 소괄호 사용하지 않음

```go
package main

import "fmt"

func main() {
	var a int = 20
	b := 20

	// ex1
	if a >= 15 {
		fmt.Println("over 15")
	}

	// ex2
	if b >= 25 {
		fmt.Println("over 25")
	}
}
```



### Error Case 1.

- byte code로 만들면서 세미콜론을 끝에 붙인다. 그러므로 괄호 위치 중요. 

```go
package main

import "fmt"

func main() {
	var a int = 20
	b := 20
  
	if b >= 25
	{	// 괄호 위치 때문에 에러 발생
		fmt.Println("over 25")
	}
}
```



### Error Case 2.

- 괄호 임의 생략하면 에러 발생

```go
package main

import "fmt"

func main() {
	var a int = 20
	b := 20
  
	if b >= 25
  	fmt.Println("over 25")
}
```



### Error Case 3.

- 0, 1로 검사 불가. Boolean으로만 할 수 있음

```go
package main

import "fmt"

func main() {
	var a int = 20
	b := 20
  
  if c :=1; c{
   fmt.Println("True") 
  }
}
```



### Error Case 4.

- if문 안에서 c가 사용 후 소멸되므로 if문 밖에서 `c += 20` 하게 되면 에러 발생.

```go
package main

import "fmt"

func main() {
	var a int = 20
	b := 20
  
  if c :=40; c >= 35 {
   fmt.Println("over 35") 
  }
  
  c += 20   // 에러 발생
}
```



## IF ... ELSE

```go
package main

import "fmt"

func main() {
	var a int = 50
	b := 70 // 함수 안에서만 사용할 짧은 선언

	// ex1
	if a >= 65 {
		fmt.Println("over 65")
	} else {
		fmt.Println("under 65")
	}

	// ex2
	if b >= 70 {
		fmt.Println("over 70")
	} else {
		fmt.Println("under 70")
	}
}
```



## IF ... ELSE IF

```go
package main

import "fmt"

func main() {
	i := 100

	// if - else if example(1)
	if i >= 120 {
		fmt.Println("over 120")
	} else if i >= 100 && i < 120 {
		fmt.Println("over 100 under 120")
	} else if i < 100 && i >= 50 {
		fmt.Println("over 50 under 100")
	} else {
		fmt.Println("under 50")
	}
}
```



## Switch

- golang만의 switch문 특징 익히는 것 중요
- `switch` 뒤 `expression` 생략 가능
- `case` 뒤 `expression` 생략 가능
- 자동 `break` 때문에 fallthrough 존재
- 값이 아닌 Type으로 분기 가능

### Example 1

```go
package main

import "fmt"

func main() {
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "은/는 음수")
	case a == 0:
		fmt.Println(a, "은/는 0")
	case a > 0:
		fmt.Println(a, "은/는 양수")
	}
}
```



### Example 2

-  Golang에서 더 선호하는 스타일
- clean code: b를 switch 안에서 선언하여 스코프가 생성 되어 b의 범위가 switch 안으로 제한
- 타 언어에 비하여  `go` 는  `switch` 의 비중 높다.

```go
package main

import "fmt"

func main() {
	// ex2: Golang에서 더 선호하는 스타일
	// b의 범위가 switch 안으로 제한 -> clean code
	switch b := 27; {
	case b < 0:
		fmt.Println(a, "은/는 음수")
	case b == 0:
		fmt.Println(a, "은/는 0")
	case b > 0:
		fmt.Println(a, "은/는 양수")
	}
}
```



### Example 3

- `case` 뒤 `expression` 생략 가능 문법

```go
package main

import "fmt"

func main() {
	switch c := "go"; c {		// ; 뒤에 c -> expression 생략 가능
	case "go":
		fmt.Println("Go")
	case "Java":
		fmt.Println("Java")
	default:
		fmt.Println("일치하는 값 없음.")
	}
}
```



### Example 4

- `c + "lang"` : 연산자로 계산 가능

```go
package main

import "fmt"

func main() {
	switch c := "go"; c + "lang" {    // 연산자로 계산 가능
	case "golang":
		fmt.Println("Go")
	case "java":
		fmt.Println("Java")
	default:
		fmt.Println("일치하는 값 없음.")
	}
}
```



### Example 5

- `c + "lang"` : 연산자로 계산 가능

```go
package main

import "fmt"

func main() {
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i < j")
	case i == j:
		fmt.Println("i == j")
	case i > j:
		fmt.Println("i > j")
	}
}
```



### Example 6

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ex1
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " 50이상 100미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " 25이상 50미만")
	default:
		fmt.Println("i -> ", i, " 기본값")
	}
}
```



### Example 7

```go
package main

import "fmt"

func main() {
	// ex1
	a := 30 / 15
	switch a {
	case 2, 4, 6: // i가 2|4|6 인 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "는 홀수")
	}
}
```



### Example 8

- Go는 자동 `break`. 
- 그렇기 때문에 fallthrough 존재
- fallthrough: 조건에 맞지 않아도 다음 case 문에 들어가서 실행

```go
package main

import "fmt"

func main() {
	switch e := "Go"; e {
	case "Java":
		fmt.Println("Java")
		fallthrough
	case "Go":
		fmt.Println("Go")
		fallthrough
	case "Python":
		fmt.Println("Python")
		fallthrough
	case "Ruby":
		fmt.Println("Ruby")
		fallthrough
	case "C":
		fmt.Println("C")
	}
}

--------------------------------------
Go
Python
Ruby
C
```





## for

- Go는 **반복문**이 for문 한개만 있다.
- Go에서 유일하게 제공되는 반복문
- 다양한 사용법 숙지

### for example1: for1.go

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("ex1: ", i)
	}
}
```



### 에러 발생 case

```go
package main

import "fmt"

func main() {
	// error 1
	for i := 0; i <5; i++ 
	{
		fmt.Println("ex1: ", i)
	}

	// error 2
	for i := 0; i < 5; i++
		fmt.Println("ex1: ", i)
}
```



### Infinite Loop

```go
package main

import "fmt"

func main() {
	for {
		fmt.Println("ex2: Hello, Go!")
		fmt.Println("ex2: Infinite loop!")
	}
}
```



### Range

```go
package main

import "fmt"

func main() {
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
}
```



### for example2: for2.go

```go
package main

import "fmt"

func main() {
	// ex1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println("ex1: ", sum1)

  // ex2: 가독성 향상
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
	}
	fmt.Println("ex2: ", sum2)
  
  // ex3: while문과 유사한 사용법
  sum3, i := 0, 0

	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3 : ", sum3)
  
  // ex4: 다양한 갯수의 변수 증감연산식 가능
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4: ", i, j)
	}
}
```



### error case

- error1: Go에서 증감연산은 반환 값 없음

```go
package main

import "fmt"

func main() {
	// error 1
	for i <= 100 {
		sum2 += i
		j := i++    // Go에서 증감연산은 반환 값 없음
	}
	fmt.Println("ex2: ", sum2)
  
  // error 2: i++, j += 10
  for i, j := 0, 0; i <= 10; i, i++, j += 10 {
	fmt.Println("ex4: ", i, j)
}
```



### for example3 : for3.go

- Loop label
- continue
- label을 활용한 continue (자주 쓰이는 패턴은 아님)

```go
package main

import "fmt"

func main() {
	// ex1: Loop label
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex1: ", i, j)
		}
	}

	// ex2: continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex2: ", i)
	}

	// ex3: label을 활용한 continue
Loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex3: ", i, j)
		}
	}
}
```



# Go 문법 특징 정리

- 주석
- 세미콜론
- 연산(증감, 전치)
- 코드 자동 서식

## Feature 1

- 모호하거나 애매한 문법 금지
- 후치(후위) 연산자 허용 (ex) i++
- 전치(전위) 연산자 비 허용 (ex) ++i
- 증감 연산은 반환 값 없음. sum := i++ (x)
- 포인터(Pointer) 허용. 단, 복잡한 연산은 비허용
- 주석: // , /**/

```go
// GO 특징 (1) : feature1.go
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
```



## Feature 2

- 문장 끝 세미콜론 주의
- 자동으로 끝에 세미콜론 삽입
- 두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론 사용(권장하지 않음)
- 반복문 및 제어문(조건문)(if, for) 사용할 경우에 주의

```go
// GO 특징 (2) 
package main

import "fmt"

func main() {
	// ex 1
	for i := 0; i <= 10; i++ {
		// fmt.Print("ex1: ", i); fmt.Println(i) // 두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론 사용(권장하지 않음)
		fmt.Print("ex1: ", i);
		fmt.Println("ex1: ", i);

	}

	// ex 2
	for j := 10; j >= 0; j-- {   // 괄호 위치 중요
		fmt.Println("ex2 : ", j)
	}
}
```



## Feature 3

- 코드 서식 지정
- `fmt` : 한 사람이 코딩한 것 같은 일관성 유지를 도와 줌
- 코드 스타일 유지
  - `gofmt -h` : 사용법 설명
  - `gofmt -w` : 원본 파일에 반영

```go
// Go 특징 (3)
package main

import "fmt"

func main() {
	// 코드 서식 지정: 한 사람이 코딩한 것 같은 일관성 유지

	// ex1
	for i := 0; i <= 100; i++ {
		fmt.Println("ex1 :", i)
	}
}
```

```shell
$ gofmt -w .\feature3.go
```

