package base

import (
	"fmt"
	"strings"
)

func StringDemo() {
	var str string = "衣香人影太匆匆"
	fmt.Printf("%s\n", str)
	var str2 string = `
		枯蓬老树昏鸦
		小桥流水人家
		古道西风瘦
	`
	fmt.Println(str2)
	fmt.Println(len(str2))
}

func StringDemo2() {
	s1 := "hello "
	s2 := "world"
	s3 := s1 + s2
	fmt.Println(s3)
}

func StringDemo3() {
	s1 := "hello"
	s2 := "world"
	s3 := fmt.Sprintf("%s %s ，你好 世界", s1, s2)
	fmt.Println(s3)
}

func StringDemo4() {
	s1 := "hello world This is golang"
	arr := strings.Split(s1, " ")
	fmt.Println(arr)
	fmt.Println(len(arr))
}

func StringDemo5() {
	s1 := "hello世界"
	for i, v := range s1 {
		fmt.Printf("%d %c\n", i, v)
	}
}
