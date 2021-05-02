package slice_demo

import "fmt"

/**
切片是一个引用类型
切片是一个 长度可变的数组
因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中 切片比数组更常用
声明切片的格式是： var identifierName []type（不需要说明长度）
*/

func SliceDemo() {
	len := 10
	var arr []int = make([]int, len)
	for i := 0; i < 10; i++ {
		arr[i] = i
		fmt.Println(arr[i])
	}
}
