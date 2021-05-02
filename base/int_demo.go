package base

import "fmt"

func IntDemo() {
	var number int = 100
	fmt.Printf("%d\n", number)
	fmt.Printf("%b\n", number) // 二进制
	fmt.Printf("%o\n", number) // 八进制
	fmt.Printf("%x\n", number) // ff
	fmt.Printf("%X\n", number) // FF
}

func IntDemo2() {
	a := 100
	b := 0b1001
	c := 072
	d := 0xFF
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
}
