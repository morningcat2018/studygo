package func_demo

import "fmt"

// defer 类似 java 中的 finally 语句
// 在 return 之前执行
func DeferDemo() {
	defer fmt.Println("defer 语句")
	fmt.Println("Hello world")
}

// 多个 defer 语句，会按照先进后出顺序执行【类似栈的操作，入栈 出栈】
func DeferDemo2() {
	defer f1(1)
	fmt.Println("Hello defer 1")
	defer f1(2)
	fmt.Println("Hello defer 2")
	defer f1(3)
	fmt.Println("Hello defer 3")
}

func f1(x int) {
	fmt.Println(x)
}

/**
defer 多用于函数结束之前释放资源 （文件句柄 数据库连接 socket连接...）
**/

// defer 语句中已经无法对函数内的变量进行改动了
func DeferDemo3() {
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
}

func f2() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

/**
匿名函数
func() {
	x++
}

调用匿名函数
func() {
	x++
}()

*/

// 返回6 （注意：此处返回值已命名）
func f3() (x int) {
	defer func() {
		x++ // 第二步 x=x+1
	}()
	return 5 // 	第一步 x=5
} // 第三步 执行 RET 指令

func f4() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f5() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func DeferDemo4() {}
