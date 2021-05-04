package slice_demo

import "fmt"

// append 追加元素
func SliceAppendDemo1() {
	s1 := []int{1, 3, 5}
	printIntSlice0(s1)
	s1 = append(s1, 7) // 扩容策略 slice.go row:144
	printIntSlice0(s1)
	//s1[4] = 9
	s1 = append(s1, 9)
	printIntSlice0(s1)
}

func SliceAppendDemo2() {
	var arr []int = make([]int, 5, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = i*2 + 1
	}
	printIntSlice0(arr)

	arr = append(arr, 11, 13, 15, 17, 19)
	printIntSlice0(arr)

	arr = append(arr, 21)
	printIntSlice0(arr)

	ss := []int{2, 4, 6}
	arr = append(arr, ss...)
	printIntSlice0(arr)
}

func SliceAppendDemo3() {
	var arr []int = make([]int, 3)

	ss := []int{2, 4, 6}
	arr = append(arr, ss...) // ... 表示将ss分解成单个元素
	printIntSlice0(arr)
}

// 删除
func SliceAppendDemo4() {
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := a1[:]
	fmt.Println(s1) // 删除前

	s1 = append(s1[:1], s1[2:]...) // 切片删除索引1的元素
	fmt.Println(s1)                // 删除后
	fmt.Println(a1)
}

func SliceAppendDemo5() {}

// copy 深度拷贝
// 赋值 浅度拷贝
func SliceCopyDemo1() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3)
	copy(a3, a1)
	a1[0] = 100
	fmt.Println(a1, a2, a3)
}
