package struct_demo

import (
	"fmt"
	"studygo/time_demo"
	"time"
)

// 自定义类型 [可以定义方法]
type myInt int

// 类型别名
type youInt = int

func StructDemo0() {
	var x myInt = 100
	fmt.Printf("%T\n", x) // struct_demo.myInt
	var y youInt = 100
	fmt.Printf("%T\n", y) // int
}

// 结构体
func StructDemo1() {
	var x struct {
		idNo string
		name string
	}
	x.idNo = "340001"
}

// 结构体 + 自定义类型
type Person struct {
	idNo       string
	name       string
	profession string
	age        myInt
	birthday   time.Time
}

type office struct {
	officeNo, officeName string
	persons              []Person
}

// 结构体的方法
func (p Person) printPerson() {
	fmt.Println(p.idNo, p.name, p.profession, p.age, time_demo.GetTimeString(&p.birthday))
}

func StructDemo2() {
	var person1 Person
	person1.idNo = "341993001"
	person1.name = "morningcat"
	person1.profession = "go programmer"
	person1.birthday = time_demo.GetTime("1993-04-05")
	person1.age = 28
	fmt.Println(person1)
	person1.printPerson()
}

func changeName(p Person) { // 值拷贝
	p.name = "hello"
}

func changeName2(p *Person) { // 引用
	(*p).name = "hello"
	//p.name = "hello" // 语法糖 解析成上诉语句
}

func StructDemo3() {
	var person1 Person
	person1.name = "morningcat"
	fmt.Println(person1.name)
	changeName(person1)
	fmt.Println(person1.name)

	changeName2(&person1)
	fmt.Println(person1.name)
}

func StructDemo4() {
	var person2 = new(Person) // Person指针 （结构体指针）
	person2.name = "morningcat"
	fmt.Println(person2.name)

	changeName2(person2)
	fmt.Println(person2.name)
}

func StructDemo5() {
	// 初始化
	var person3 = Person{idNo: "3419901", name: "mengzhang6", profession: "java", age: 28, birthday: time_demo.GetTime("1993-12-01")}
	person3.printPerson()

	// 格式化
	var person4 = Person{
		idNo:       "3419901",
		name:       "mengzhang6",
		profession: "java",
		age:        28,
		birthday:   time_demo.GetTime("1993-12-01"),
	}
	person4.printPerson()

	// 全量赋值 可省略key
	var person5 = Person{"3419901", "mengzhang6", "java", 28, time_demo.GetTime("1993-12-01")}
	person5.printPerson()
}

func NewPerson(idNo, name string) *Person {
	return &Person{idNo: idNo, name: name}
}

func StructDemo6() {
	var p = NewPerson("3419002", "morningcat")
	p.printPerson()
	fmt.Println(p.GetPersonStr())
}

// 方法 作用于特定类型的函数

func (p Person) GetPersonStr() string { //  值接收者
	return fmt.Sprintf(p.idNo, p.name, p.profession, p.age, time_demo.GetTimeString(&p.birthday))
}

func (p *Person) ChangrPersonName(name string) { // 指针接收者
	p.name = name
}

func StructDemo7() {
	var p = NewPerson("3419002", "morningcat")
	fmt.Println(p.name)
	p.ChangrPersonName("world")
	fmt.Println(p.name)
}

// 任意自定义类型都可以加方法

func (x myInt) helloInt() {
	fmt.Printf("Hello %d\n", x)
}

func StructDemo8() {
	var x myInt = 12
	x.helloInt()
}

// 只能给自己 package 中的`自定义类型`添加方法

func StructDemo10() {
	var x int32 = 12
	var y = int32(12) // 另一种定义方法
	fmt.Println(x, y)
}
