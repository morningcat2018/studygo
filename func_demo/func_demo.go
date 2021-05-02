package func_demo

import "fmt"

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func FuncDemo() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(3))

	// 数值常量是高精度的 值。
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// 与 add(x int, y int) 等效
func add(x, y int) int {
	return x + y
}

// 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量
func split(sum int) (x, y int) {
	x = sum * 4 / 9.0
	y = sum - x
	// 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
	return
}

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
