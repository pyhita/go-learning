package main

import "fmt"

// 定义接口，实际就是一个指针
type Animal interface {
	GetColor() string
	SetColor(color string)
}

//只要某个类实现了全部的方法，就说明这个类实现了这个接口
type Dog struct {
	color string
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) SetColor(color string) {
	this.color = color
}

//
type Cat struct {
	color string
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) SetColor(color string) {
	this.color = color
}

func showAnimal(animal Animal) {
	fmt.Println("--------")
	fmt.Println("animal's color: ", animal.GetColor())
}

func main() {
	cat := Cat {"red"}
	// 接口实际上是一个指针
	showAnimal(&cat)

	dog := Dog {"yellow"}
	// 接口实际上是一个指针
	showAnimal(&dog)
}
