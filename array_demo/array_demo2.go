package array_demo

import "fmt"

func Array2Demo1() {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 7
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}
}

// for-range
func Array2Demo2() {
	var arr1 [5]int
	for i, _ := range arr1 {
		arr1[i] = i * 7
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}
}

// Go 语言中的数组是一种 值类型（不像 C/C++ 中是指向首元素的指针）
func Array2Demo3() {
	//a := [...]string{"a", "b", "c", "d"}
	var a [4]string = [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}

func Array2Demo4() {
	var arr1 = new([5]int) // arr1 的类型是 *[5]int
	var arr2 [5]int        // arr2 的类型是 [5]int
	for i, _ := range arr1 {
		arr1[i] = i * 7
		arr2[i] = i * 7
	}
	arr3 := *arr1
	arr3[2] = 300
	fmt.Println(arr3[2])
	fmt.Println(arr1[2])

	arr4 := arr2
	arr4[2] = 400
	fmt.Println(arr4[2])
	fmt.Println(arr2[2])

	var arr5 *[5]int = arr1
	arr5[2] = 500 // 会改变值
	fmt.Println(arr5[2])
	fmt.Println(arr1[2])

	var arr6 *[5]int = &arr2
	arr6[2] = 600 // 会改变值
	fmt.Println(arr6[2])
	fmt.Println(arr2[2])
}

func Array2Demo4_2() {
	var arr1 = new([5]int) // arr1 的类型是 *[5]int
	var arr2 [5]int        // arr2 的类型是 [5]int
	var len int = 5
	var slice1 []int = make([]int, len) // slice1 的类型是 *[]int
	for i, _ := range arr1 {
		arr1[i] = i * 7
		arr2[i] = i * 7
		slice1[i] = i * 7
	}
	eDemo2(arr1)
	printSlice2(*arr1)
	eDemo2(&arr2)
	printSlice2(arr2)
	eDemo(slice1)
	printSlice(slice1)
}

func eDemo(arr1 []int) {
	arr1[2] = 200 // 测试 arr1 是否为指针
}

func eDemo2(arr1 *[5]int) {
	arr1[2] = 200 // 测试 arr1 是否为指针
}

func printSlice(arr1 []int) {
	for i, _ := range arr1 {
		fmt.Printf("%d\t", arr1[i])
	}
	fmt.Println()
}

func printSlice2(arr1 [5]int) {
	for i, _ := range arr1 {
		fmt.Printf("%d\t", arr1[i])
	}
	fmt.Println()
}

func Array2Demo5() {
	var ar [3]int = [...]int{1, 2, 3}
	f(ar) // passes a copy of ar
	fmt.Println(ar[0])
	fp(&ar) // passes a pointer to ar
	fmt.Println(ar[0])
}
func f(a [3]int) {
	fmt.Println(a)
	a[0] = 100
}
func fp(a *[3]int) {
	fmt.Println(a)
	a[0] = 100
}

func Array2Demo6() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	fmt.Println(arrAge)
	var arrKeyValue = [5]string{0: "Hello", 3: "Chris", 4: "Ron"}
	fmt.Println(arrKeyValue)

	var arrLazy = [...]int{5, 6, 7, 8, 22}
	fmt.Println(arrLazy)
}

/*
把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
传递数组的指针
使用数组的切片
*/
func SliceDemo5() {
	//var slice1 []type = make([]type, len)
	var len int = 10
	var slice1 []int = make([]int, len)
	for i, _ := range slice1 {
		slice1[i] = i
	}
	eDemo(slice1)
	printSlice(slice1)
}
