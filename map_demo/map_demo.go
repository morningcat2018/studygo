package map_demo

import "fmt"

func MapDemo() {
	var myMap map[string]int = make(map[string]int, 10)
	myMap["x"] = 1
	myMap["y"] = 3
	myMap["z"] = 5
	fmt.Println(myMap)

	// 取值
	v, ok := myMap["zs"]
	if !ok {
		fmt.Println("not found this key")
	} else {
		fmt.Println(v)
	}

	// 遍历
	for k, v := range myMap {
		fmt.Println(k, v)
	}

	// 删除
	delete(myMap, "y")
	fmt.Println(myMap)
	delete(myMap, "zs") // 不做操作
	fmt.Println(myMap)
}

// 文档 https://studygolang.com/pkgdoc

func MapDemo2() {
	var s1 = make([]map[string]int, 10)
	s1[0] = make(map[string]int)
	s1[0]["Hello"] = 123
	fmt.Println(s1)

	s1[1]["Hello"] = 123 // error : not make map
	fmt.Println(s1)
}

func MapDemo3() {
	var s2 = make(map[string][]int, 10)
	s2["a"] = []int{1, 2, 3}
	fmt.Println(s2)
}

func MapDemo4() {}
