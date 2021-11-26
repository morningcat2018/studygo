package main

import (
	"fmt"
	"studygo/bitmap"
)

func main() {
	// fmt.Println("Hello world")
	// time_demo.TimeFormat()
	// time_demo.Time2()
	// base.StringDemo5()
	//closure_demo.TestClosure3()
	//array_demo.ArrayDemo8()
	//slice_demo.SliceAppendDemo5()
	//map_demo.MapDemo4()
	// func_demo.Add5(1, 2, 4, 6)
	// func_demo.ClosureDemo5()
	// func_demo.Demo5Main()

	// struct_demo.StructDemo10()
	// interface_demo.InterfaceDemo1()
	// interface_demo.BlankInterfaceDemo2("hello", "world")
	// interface_demo.BlankInterfaceDemo4()

	// book_manage.DataPanel()
	//package_demo.PackageDemo0()

	//channel_demo.ChanDemo2()

	//mysql_db.DbMain()
	//struct_demo.Main()

	TestBitMap()

}

func TestBitMap() {
	bitMap := bitmap.New(0)
	bitMap.Set(1000)
	bitMap.Set(10001)
	fmt.Println(bitMap.Get(1001))
	fmt.Println(bitMap.Get(10001))
}
