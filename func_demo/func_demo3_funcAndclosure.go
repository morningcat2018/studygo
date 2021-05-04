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

//////  函数作为出参 闭包

func ClosureDemo4() {

}

// 面向函数编程
