package slice_demo

import "fmt"

/**
切片是一个引用类型
切片是一个 长度可变的数组
因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中 切片比数组更常用
声明切片的格式是： var identifierName []type（不需要说明长度）
*/

func SliceDemo() {
	lens := 10
	var arr []int = make([]int, lens)
	for i := 0; i < len(arr); i++ {
		arr[i] = i

	}
	printIntSlice("arr", arr)
}

func printIntSlice0(arr []int) {
	printIntSlice("name", arr)
}

func printIntSlice(name string, arr []int) {
	fmt.Println(arr)
	fmt.Printf("%s ->len=%d cap=%d\n", name, len(arr), cap(arr)) // 长度  容量
}

/**
切片（Slice）是一个拥有相同类型元素的可变长度的序列
基于数组类型做的一层封装
支持自动扩容
它的内部结构包含地址、长度和容量
*/

func SliceDemo2() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15}
	s1 := a1[1:5] // 1～7
	printIntSlice("s1", s1)
	s2 := a1[:5] // 0～7
	printIntSlice("s2", s2)
	s3 := a1[5:] // 5～7
	printIntSlice("s3", s3)
	s4 := a1[:] //0～7
	printIntSlice("s4", s4)

	// 切片再切片
	s6 := s1[1:3]
	printIntSlice("s6", s6)
}

// 底层数组从切片的第一个位置到最后位置的数量

func SliceDemo3() {
	var arr []int = make([]int, 5, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	printIntSlice("arr", arr)
}

// 切片只是一个引用 真正的数据都保存在底层数组中

func SliceDemo0() {
	var arr []int // nil 值的切片没有底层数组 长度和容量都是0
	fmt.Println(arr == nil)
}

// 两种遍历方式
func SliceDemo4() {
	var arr []int = make([]int, 5, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 5
	}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
