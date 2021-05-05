package interface_demo

import (
	"fmt"
	"time"
)

// 空接口
type blank interface{} // 所有类型都实现了空接口 即所有类型都相当于这种接口类型的变量

// 空接口 作为入参
// interface{} 空接口类型
func BlankInterfaceDemo1(i interface{}) {
	fmt.Println(i)
}

// 空接口列表
func BlankInterfaceDemo2(is ...interface{}) {
	for i, v := range is {
		fmt.Println(i, v)
	}
}

func BlankInterfaceDemo3() {
	var interMap map[string]interface{} = map[string]interface{}{}
	interMap["hello"] = 514
	interMap["world"] = "world"
	interMap["mc"] = time.Now()
	interMap["xx"] = 's'
	fmt.Println(interMap)
}
