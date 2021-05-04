package pointer_demo

import "fmt"

// x = *&x
func PointerDemo1() {
	var x = 123
	fmt.Println(&x)
	var xPointer *int = &x
	fmt.Println(xPointer)

	var y = *xPointer
	fmt.Println(y)
}

func PointerDemo2() {

}

func PointerDemo3() {

}

func PointerDemo4() {

}
