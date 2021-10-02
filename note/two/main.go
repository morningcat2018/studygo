package main

import "fmt"

func main() {
	var strValue string = "hello world"
	fmt.Println(strValue)
}

func hello(count int) string {
	if count >= 60 {
		fmt.Println("已及格")
	}
	return ""
}
