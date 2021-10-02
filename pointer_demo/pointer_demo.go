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
	var a *int   // nil
	a = new(int) // 使用 new 分配一块 int 类型大小的内存
	// 返回一个 int 类型指针

	*a = 101
	fmt.Println(*a)
}

/*
make也是用于内存分配的，区别于new
它只用于slice、map 以及 chan 的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型
因为这三种类型就是引用类型，所以就没有必要返回他们的指针了
https://www.liwenzhou.com/posts/Go/07_pointer/
*/
