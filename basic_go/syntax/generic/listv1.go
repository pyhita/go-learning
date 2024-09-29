package main

// 泛型接口
type ListV1[T any] interface {
	Add(index int, ele T)
	Delete(index int)
	Append(ele T)
}
