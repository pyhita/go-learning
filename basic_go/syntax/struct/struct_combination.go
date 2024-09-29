package main

// 接口可以组合接口
// 结构体可以组合结构体
// 结构体可以组合接口

type List interface {
	Add(ele interface{})
	Delete(ele interface{})
}

type ArrayList interface {
	List
	Get(index int)
}

type Animal struct {
	Name string
	Age  int
	Sex  string
}

func (a *Animal) Eat() {
	println("animal can eat")
}

func (a *Animal) Walk() {
	println("animal can walk")
}

type Fish struct {
	Animal
	SwimSpeed float64
}

func (f *Fish) Eat() {
	println("fish can eat")
}

func (f *Fish) Swim() {
	println("fish can swim")
}

func TestCombination() {
	// Fish 可以调用Animal 实现的方法
	fish := &Fish{}

	fish.Walk()
	// 同名方法，优先调用自己的
	fish.Eat()
}
