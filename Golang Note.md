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



