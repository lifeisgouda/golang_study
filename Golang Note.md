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



### 상수 여러게 선언 (const2.go)

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

