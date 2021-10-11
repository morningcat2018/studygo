package struct_demo

import (
	"fmt"
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
