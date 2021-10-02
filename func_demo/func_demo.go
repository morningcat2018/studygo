package func_demo

// 与 add(x int, y int) 等效
func add(x, y int) int {
	return x + y
}

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量
func split(sum int) (x, y int) {
	x = sum * 4 / 9.0
	y = sum - x
	// 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
	return
}

// 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}
