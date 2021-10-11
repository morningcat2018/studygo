package interface_demo

import (
	"fmt"
	"testing"
)

// 	定义接口
type speaker interface {
	speak() // 只要实现了 speak 方法的变量都是 speaker 类型
}

type dog struct{ name string }
type cat struct{ age int }

func (d dog) speak() {
	fmt.Println(d.name, ":汪汪汪")
}
func (c cat) speak() {
	fmt.Println("喵喵喵", c.age)
}

// 使用接口
func call(sp speaker) {
	sp.speak()
}

// 调用接口方法
func TestInterfaceDemo1(t *testing.T) {
	var c1 cat
	c1.age = 12
	call(c1)

	var d1 dog
	d1.name = "dico"
	call(d1)

	var sp speaker
	sp = c1
	sp.speak()
}
