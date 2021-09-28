
## Go Setting

GOPATH -> ~/go/

GOROOT -> /usr/local/go/

## go version

> go version

go version go1.16.2 darwin/amd64

## go run

> go run main.go 

## go build

> go build -o myRunExec

> ./myRunExec

## go install


## go clean

> go clean -i -n

## go get

> go env -w GOPROXY=https://goproxy.cn

// go env -w GOPROXY=https://goproxy.io

// export GOPROXY=https://goproxy.io

---

> go get github.com/go-sql-driver/mysql

// cd $GOPATH/pkg/mod

---

GoLand

Preference -> Go -> Go Modules(vgo) 勾选启用 -> Proxy 选择 direct

## mysql

> which mysql
>
> cd /usr/local/Cellar/mysql/8.0.23/bin
>
> mysql.server start


