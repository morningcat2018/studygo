package map_demo

import (
	"fmt"
	"testing"
)

func TestMap1(t *testing.T) {
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

func TestMap2(t *testing.T) {
	// 文档 https://studygolang.com/pkgdoc
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

// https://www.liwenzhou.com/posts/Go/08_map/
func PointerDemo3() {
	// map
	var myMap map[string]int = make(map[string]int)
	myMap["x"] = 1
	myMap["x"] = 2
	myMap["y"] = 3
	fmt.Println(myMap)
}
