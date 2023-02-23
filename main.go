package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) changeName(newName string) {
	p.name = newName
}

func (p *Person) changeAge(newAge int) {
	p.age = newAge
}

func main() {
	person1 := Person{}
	person2 := &Person{}

	person1.changeAge(123)
	person1.changeName("alion")

	person2.changeAge(321)
	person2.changeName("wu")

	fmt.Printf("%v -- %v", person1, person2)
	//输出：{ 123} -- &{ 321}

}
