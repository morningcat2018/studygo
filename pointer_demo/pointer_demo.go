package pointer_demo

import "fmt"

// x = *&x
func PointerDemo1() {
	var x = 123
	fmt.Println(&x)
	var xPointer *int = &x // int 类型指针
	fmt.Println(xPointer)

	var y = *xPointer
	fmt.Println(y)
}

func PointerDemo2() {
	var a *int
	*a = 101
	fmt.Println(*a)
}

func PointerDemo3() {

}

func PointerDemo4() {

}
