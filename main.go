package main

import (
	"studygo/array_demo"
	"studygo/channel_demo"
	"studygo/closure_demo"
	"studygo/map_demo"
	"studygo/slice_demo"
	// interface_demo "github.com/morningcat2018/studygo/interface_demo" 自定义导入 即给导入的包起别名
	// _ "xxx/yyy" 匿名导入 不使用包中的模块 例如 mysql 只使用了 init() 方法
)

func main() {
	// fmt.Println("Hello world")
	// time_demo.TimeFormat()
	// time_demo.Time2()
	// base.StringDemo5()
	closure_demo.TestClosure3()
	array_demo.ArrayDemo8()
	slice_demo.SliceAppendDemo5()
	map_demo.MapDemo4()
	// func_demo.Add5(1, 2, 4, 6)
	// func_demo.ClosureDemo5()
	// func_demo.Demo5Main()

	// struct_demo.StructDemo10()
	// interface_demo.InterfaceDemo1()
	// interface_demo.BlankInterfaceDemo2("hello", "world")
	// interface_demo.BlankInterfaceDemo4()

	// book_manage.DataPanel()
	//package_demo.PackageDemo0()

	channel_demo.ChanDemo()
}
