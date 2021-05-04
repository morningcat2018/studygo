package func_demo

import "fmt"

func Add(x, y int) int {
	var sum = x + y
	fmt.Println(sum)
	return sum
}

func Add2(x, y int) (sum int) { // 命名返回值
	sum = x + y
	fmt.Println(sum)
	return // 不写返回值也可 因为返回值已经被命名
}

func Add3(x, y int) (int, error) { // 多返回值
	var sum = x + y
	return sum, nil
}

func Add4(x, y int) (sum int, err error) {
	sum = x + y
	err = nil
	return
}

// 可变长参数 必须放在参数列表的最后
func Add5(x ...int) {
	// x 是一个切片类型 ，[]int
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
}
