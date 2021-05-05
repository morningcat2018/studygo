package interface_demo

import "fmt"

// 	定义接口
type speaker interface {
	speak() // 只要实现了 speak 方法的变量都是 speaker 类型
}

type dog struct{}
type cat struct{}

func (d dog) speak() {
	fmt.Println("wang wang wang...")
}
func (c cat) speak() {
	fmt.Println("miao miao miao...")
}

// 使用接口
func call(sp speaker) {
	sp.speak()
}

// 调用接口方法
func InterfaceDemo1() {
	var c1 cat
	var d1 dog
	call(c1)
	call(d1)

	var sp speaker
	sp = c1
	fmt.Println(sp)
}
