package main

import "fmt"

// 利用struct 来实现封装特性

type Person struct {
	Name string
	Age int
	// 私有属性首字母小写，其他包不可见
	no int
}

// 定义Person的私有方法
func (this *Person) GetName() string {
	return this.Name
}

func (this *Person) SetName(Name string) {
	this.Name = Name
}

func (this *Person) Show() {
	fmt.Println("name: ", this.Name, "age: ", this.Age)
}

func main() {
	p := Person {
		Name: "kante",
		Age: 22,
	}

	p.Show()
	p.SetName("liam")
	p.Show()
}

