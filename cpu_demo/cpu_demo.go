package cpu_demo

import (
	"fmt"
	"runtime"
)

func CpuDemo() {
	c := runtime.NumCPU()
	fmt.Println(c)

	// 设置当前程序可使用的最大核心数
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
}
