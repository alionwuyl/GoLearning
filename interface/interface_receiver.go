package main

import "fmt"

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	// 这里的初始化是核心代码，c是指针
	var c coder = &Gopher{"Go"}
	c.code()
	c.debug()
}

// 上面的代码输出：
// I am coding Go language
// I am debuging Go language

// 如果将`var c coder = &Gopher{"Go"}` 改为 `var c coder = Gopher{"Go"}`
// 代码会无法通过编译
//  src/main.go:23:6: cannot use Gopher literal (type Gopher) as type coder in assignment:
// 	Gopher does not implement coder (debug method has pointer receiver)

// 如果方法的接收者是值类型，无论调用者是对象还是对象指针，修改的都是对象的副本，不影响调用者；如果方法的接收者是指针类型，则调用者修改的是指针指向的对象本身。
// 使用指针作为方法的接收者的理由：
// 方法能够修改接收者指向的值。
// 避免在每次调用方法时复制该值，在值的类型为大型结构体时，这样做会更加高效。
