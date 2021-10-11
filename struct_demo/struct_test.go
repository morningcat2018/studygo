package struct_demo

import (
	"fmt"
	"testing"
	"time"
)

// 自定义类型 [可以定义方法]
type myInt int

// 结构体 + 自定义类型
type Person struct {
	idNo       string
	name       string
	profession string
	age        int
	birthday   time.Time
}

const YYYY_MM_DD = "2006-01-02"

func getTime(timeString string) time.Time {
	t, _ := time.Parse(YYYY_MM_DD, timeString)
	return t
}

// 结构体的方法
func (p Person) printPerson() {
	fmt.Println(p.idNo, p.name, p.profession, p.age, (&p.birthday).Format(YYYY_MM_DD))
}

////

func NewPerson(idNo, name string) *Person {
	return &Person{idNo: idNo, name: name}
}

// 作用于特定类型的函数
func (p Person) GetPersonStr() string { //  值接收者
	return fmt.Sprintf(p.idNo, p.name, p.profession, p.age, (&p.birthday).Format(YYYY_MM_DD))
}

func TestStruct6(t *testing.T) {
	var p *Person = NewPerson("3419002", "morningcat")
	fmt.Println(p.GetPersonStr())
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
