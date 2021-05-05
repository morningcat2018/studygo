package struct_demo

import (
	"fmt"
	"studygo/time_demo"
	"time"
)

type person struct {
	idNo       string
	name       string
	profession string
	age        int
	birthday   time.Time
}

func StructDemo() {
	var person1 person
	person1.idNo = "341993001"
	person1.name = "morningcat"
	person1.profession = "go programmer"
	person1.birthday = time_demo.GetTime("1993-04-05")
	fmt.Println(person1)
}
