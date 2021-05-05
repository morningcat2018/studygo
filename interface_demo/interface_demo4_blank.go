package interface_demo

import "fmt"

// 空接口
type blank interface{}

// 空接口 作为入参
func BlankInterfaceDemo1(i interface{}) {
	fmt.Println(i)
}
