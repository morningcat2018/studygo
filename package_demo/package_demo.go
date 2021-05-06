package package_demo

import (
	. "fmt"
	_ "studygo/book_manage"
	ind "studygo/interface_demo"
)

var x = 100

func init() {
	Println("package_demo 包的初始化方法")
	Printf("%d\n", x)
}

func PackageDemo0() {
	Println("Hello PackageDemo0")

	ind.BlankInterfaceDemo4()

}
