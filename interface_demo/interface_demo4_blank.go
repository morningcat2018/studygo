package interface_demo

import (
	"fmt"
	"time"
)

// 空接口
type blank interface{} // 所有类型都实现了空接口 即所有类型都相当于这种接口类型的变量

// 空接口 作为入参
// interface{} 空接口类型
func BlankInterfaceDemo1(param interface{}) {
	fmt.Println(param)
}

// 空接口列表
func BlankInterfaceDemo2(params ...interface{}) {
	for i, v := range params {
		fmt.Println(i, v)
	}
}

// 空接口 作为 map 的 value
func BlankInterfaceDemo3() {
	var interMap map[string]interface{} = map[string]interface{}{}
	interMap["hello"] = 514
	interMap["world"] = "world"
	interMap["mc"] = time.Now()
	interMap["xx"] = 's'
	fmt.Println(interMap)
}

//////// 类型断言

/*
x.(T)
x : 类型为 interface{} 的变量
T : 表示断言 x 可能的类型
*/
func interface_demo4(param interface{}) {
	value, ok := param.(string)
	if !ok {
		fmt.Println("param not is string")
		return
	}
	fmt.Printf("param is string:%s\n", value)
}

func BlankInterfaceDemo4() {
	interface_demo4("hello")
	interface_demo4(1024)
}

// 不建议使用
func interface_demo5(param interface{}) {
	switch t := param.(type) {
	case int:
		fmt.Println("is int -> ", t)
	case string:
		fmt.Println("is string -> ", t)
	case bool:
		fmt.Println("is string -> ", t)
		// ...
	}
}
