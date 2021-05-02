package array_demo

import "fmt"

func ArrayDemo() {
	var arr [5]int // 数组 必须给定容量
	for i := range arr {
		arr[i] = i
		fmt.Println(arr[i])
	}
	fmt.Printf("%T\n", arr) // 打印类型

	// var a [3]int
	// var a [4]int
	// a = b  不能这样赋值 因为 [3]int 和 [4]int 并不是同一种类型
}

func ArrayDemo2() {
	var arr = [5]string{"1,2", "Hello", "world", "好", "go"} // 初始化
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println("len=", len(arr))
}

func ArrayDemo3() {
	arr := [5]string{"1,2,3", "Hello", "world", "好"} // 初始化
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println("len=", len(arr))
}

func ArrayDemo4() {
	arr := [...]string{"1,2,3,4", "Hello", "world", "好"} // 自动推断长度
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println("len=", len(arr))
}

func ArrayDemo5() {
	arr := [5]string{1: "1,2,3,4", 4: "Hello"} // 根据索引位置赋值
	for i, v := range arr {                    // 可以同时接收索引与值
		fmt.Println(i, v)
	}
	fmt.Println("len=", len(arr))
}

func ArrayDemo6() {
	arr := [5]string{"0", "Hello", "world", "好"}
	for _, v := range arr {
		fmt.Println(v)
	}
}
