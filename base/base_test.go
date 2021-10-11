package base

import (
	"fmt"
	"testing"
)

func TestBaseType2(t *testing.T) {

	// 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明
	// 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用
	j := 6
	fmt.Println(j)

	// 类型转换
	k := 42
	fl := float64(k)
	ux := uint(fl)
	fmt.Println(k, fl, ux)

}
