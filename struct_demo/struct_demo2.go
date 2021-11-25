package struct_demo

import "fmt"

type Animal struct {
	Name string
}

type Partner struct {
	Link string
}

type Dog struct {
	Animal
	Partner
	Id    int
	Color string
}

func (a *Animal) Eat() {
	fmt.Println(a.Name, "is eating")
}

func (p *Partner) Accompany() {
	fmt.Println("My partner is with me:", p.Link)
}

func (d *Dog) Run() {
	fmt.Println("No.", d.Id, "is running")
}

func Main() {
	var dog Dog
	dog.Id = 1
	dog.Name = "小花"
	dog.Color = "blank"
	dog.Link = "小伙伴"
	dog.Run()
	dog.Eat()
	dog.Accompany()
}
