package main

import "fmt"

// 对于结构体来说，不管是值类型还是指针类型，都可以调用值接收者或者指针接收者方法，
// 但是值接收者方法（changeName）对于数据的更改不会生效
//      -                         	     值接收者                                 	                              指针接收者
// 值类型调用者	              方法会使用调用者的一个副本，类似于“传值”	                           使用值的引用来调用方法，上例中，qcrao.growUp() 实际上是 (&qcrao).growUp()
// 指针类型调用者	   指针被解引用为值，上例中，stefno.howOld() 实际上是 (*stefno).howOld()	    实际上也是“传值”，方法里的操作会影响到调用者，类似于指针传参，拷贝了一份指针

// 实现了接收者是值类型的方法，相当于自动实现了接收者是指针类型的方法；而实现了接收者是指针类型的方法，不会自动生成对应接收者是值类型的方法。

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
