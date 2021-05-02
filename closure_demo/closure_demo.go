package closure_demo

import "fmt"

func TestClosure() {
	addFunc1 := addN(10)
	fmt.Printf("%d\n", addFunc1(3))
	fmt.Printf("%d\n", addFunc1(7))

	addFunc, subFunc := addOrSubN(10)
	fmt.Printf("%d\n", addFunc(30))
	fmt.Printf("%d\n", subFunc(30))
}

func addN(n int) func(x int) int {
	return func(x int) int {
		return x + n
	}
}

func addOrSubN(n int) (func(y int) int, func(y int) int) {
	addFunc := func(x int) int {
		return x + n
	}
	subFunc := func(x int) int {
		return x - n
	}
	return addFunc, subFunc
}
