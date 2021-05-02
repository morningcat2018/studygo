package base

import "fmt"

/**
基础数据类型
定义在函数体外部为全局变量
*/
var url string = "https://tour.go-zh.org/basics/1"
var isEat bool = true
var id1 int = 1 // int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽
// 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
var my3 = 9 // 类型推断
var id2 int8 = 1
var id3 int16 = -1
var id4 int32 = 1
var id5 int64 = 1
var id6 uint = 1
var id7 uint8 = 1
var id8 uint16 = 1
var id9 uint32 = 1
var id0 uint64 = 1
var id11 uintptr = 45
var di1 byte = 2 // uint8 的别名
var di2 rune = 8 // int32 的别名
var mo0 float32 = 0.1
var mo1 float64 = 0.31
var mi1 complex64
var mi2 complex128

func BaseType0() {
	fmt.Println(url)
}

func BaseType() {
	// 定义在函数内的为局部变量
	// 定义在函数内的变量，声明后必须使用，否则通过不了编译
	var name string = "morningcat"
	fmt.Println(name)
	var name2 = "2018" // 类型推断
	fmt.Println(name2)
	name3 := "meng" // 简短声明
	fmt.Println(name3)

}

func BaseType2() {
	var i int = 5
	fmt.Println(i, url, mo1)

	// 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明
	// 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用
	j := 6
	fmt.Println(j)

	/**
	  没有明确初始值的变量声明会被赋予它们的 [零值]。

	  零值是：
	  数值类型为 0，
	  布尔类型为 false，
	  字符串为 ""（空字符串）。
	*/
	var in_v int
	var f float64
	var bo_v bool
	var s string
	fmt.Printf("%v %v %v %q\n", in_v, f, bo_v, s)

	// 类型转换
	k := 42
	fl := float64(k)
	ux := uint(fl)
	fmt.Println(k, fl, ux)

	// 类型推导
	ii := 42           // int
	ff := 3.142        // float64
	gg := 0.867 + 0.5i // complex128
	fmt.Println(ii, ff, gg)
	fmt.Printf("type is : %T %T %T\n", ii, ff, gg)

	// 常量的声明与变量类似，只不过是使用 const 关键字。
	// 常量不能用 := 语法声明。
	const World = "世界"
	fmt.Println("Hello", World)

	const Truth = false
	fmt.Println("Go rules?", Truth)
}
