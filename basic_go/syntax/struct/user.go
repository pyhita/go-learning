package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Phone int
	Addr  string
}

// 结构体绑定方法，并注意指针和值的差别
func (this *User) SetName(name string) {
	this.Name = name
}

func (this User) SetAge(age int) {
	this.Age = age
}

func ChangeUser() {
	user := &User{Name: "kante", Age: 22}
	fmt.Printf("user, %+v \n", user)
	user.SetName("liam")
	user.SetAge(26)
	fmt.Printf("user, %+v \n", user)

}

func TestUser() {
	// 声明并初始化结构体
	u1 := &User{Name: "kante", Age: 22, Addr: "shandong"}
	fmt.Printf("u1 %+v \n", u1)

	// 声明结构体
	u2 := &User{}
	u2.Name = "liam"
	u2.Age = 25
	fmt.Printf("u2 %+v \n", u2)
}
