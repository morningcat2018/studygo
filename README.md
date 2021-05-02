


> cd github.com/morningcat2018/studygo/base

> go build

go: go.mod file not found in current directory or any parent directory; see 'go help modules'

---

> go mod edit -module=github.com/morningcat2018/studygo

---

> go help modules

> go env -w GOPROXY=https://goproxy.cn

> go get github.com/go-sql-driver/mysql
// cd $GOPATH/pkg/mod

---

#go
export GOROOT=/usr/local/go
export GOARCH=amd64
export GOOS=darwin
export GOPATH=/Users/morningcat/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin


---

