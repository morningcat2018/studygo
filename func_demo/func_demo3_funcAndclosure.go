package func_demo

import "fmt"

//////  函数赋值给变量

func fc0()             {}
func fc1() string      { return "golang" }
func fc2(x int) string { return "golang" }

func ClosureDemo0() {
	a := fc0
	fmt.Printf("%T\n", a)
	b := fc1
	fmt.Printf("%T\n", b)
	c := fc2
	fmt.Printf("%T\n", c)
}

//////  函数作为入参

func fc3(x func()) {
	fmt.Println("执行函数x")
	x()
}

func ClosureDemo1() {
	a := func() {
		fmt.Println("ClosureDemo1")
	}

	fc3(a)
}

func fc4(x func(int)) {
	fmt.Println("执行函数x")
	x(10)
}

func ClosureDemo2() {
	a := func(x int) {
		fmt.Println(x * 2)
	}
	fc4(a)
}

func fc5(x func(int) int) {
	fmt.Println(x(50))
}

func ClosureDemo3() {
	a := func(x int) (y int) {
		y = x * 2
		return
	}
	fc5(a)
}

//////  函数作为出参

func fc6(x func(int, int), m, n int) func() {
	ret := func() {
		x(m, n)
	}
	return ret
}

func ClosureDemo4() {
	var k1 = func(x, y int) {
		fmt.Println(x + y)
	}
	k2 := fc6(k1, 3, 4)
	k2()
}

//////  闭包
// 1. 函数作为出参
// 2. 函数包含外部作用域的变量

func addN(n int) func(int) int {
	return func(y int) int {
		return n + y
	}
}

func ClosureDemo5() {
	var k3 = addN(100) // 得到一个与100相加的函数
	var sum = k3(245)
	fmt.Println(sum)
}

//////  匿名函数

func ClosureDemo6() {
	// 函数内部无法声明有名称的函数
	// 可以使用匿名函数
	var f1 = func(x int, y int) {
		fmt.Println(x + y)
	}
	f1(2, 3)

	// 声明加调用
	func(x int, y int) {
		fmt.Println(x + y)
	}(4, 5)
}

//////  总结

// 面向函数编程

// 高阶函数 函数可以作为参数和返回值
