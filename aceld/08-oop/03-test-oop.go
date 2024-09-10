package main

import "fmt"

// 实现golang 继承

type Human struct {
	name string
	sex string
}
// 定义父类方法
func (this *Human) Walk() {
	fmt.Println("Human Walk")
}

func (this *Human) Eat() {
	fmt.Println("Human Eat")
}

// 定义SuperMan 继承自Human
type SuperMan struct {
	Human
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eat")
}

func main() {
	human := Human {"zhangsan", "male"}
	human.Eat()

	super := SuperMan {Human {"lisi", "male"}, 99}

	// 调用父类方法
	super.Walk()
	// 调用重写方法
	super.Eat()
}