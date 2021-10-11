package interface_demo

import (
	"fmt"
	"testing"
)

//////// 类型断言

/*
x.(T)
x : 类型为 interface{} 的变量
T : 表示断言 x 可能的类型
*/
func interface_demo4(param interface{}) {
	value, ok := param.(string)
	if !ok {
		fmt.Println("param not is string:", param)
		return
	}
	fmt.Printf("param is string: %s\n", value)
}

// 不建议使用
func interface_demo5(param interface{}) {
	switch t := param.(type) {
	case int:
		fmt.Println("is int -> ", t)
	case string:
		fmt.Println("is string -> ", t)
	case bool:
		fmt.Println("is bool -> ", t)
	case float64:
		fmt.Println("is float -> ", t)
		// ...
	}
}

func TestBlankInterfaceDemo4(t *testing.T) {
	interface_demo4("hello")
	interface_demo4(1024)
	interface_demo5("world")
	interface_demo5(1024)
	interface_demo5(true)
	interface_demo5(3.1415926)
}
