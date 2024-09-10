package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (this *User) GetName() {
	fmt.Println("get user name")
}

func getAllMsgByReflect(input interface{}) {
	// 通过反射获取类型，字段值和方法信息
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is ", inputType.Name())

	inputValue := reflect.ValueOf(input)
	fmt.Println("input value is ", inputValue)

	// 获取所有的属性信息
	for i := 0;i < inputType.NumField();i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Println(field.Name, field.Type, value)
	}

	// 获取所有的方法信息
	fmt.Println("method num: ", inputType.NumMethod())
	for i := 0;i < inputType.NumMethod();i++ {
		method := inputType.Method(i)

		fmt.Println(method.Name, " ", method.Type)
	}
}

func main() {
	user := User{1, "kante", 22}
	user.GetName()
	getAllMsgByReflect(user)
}
