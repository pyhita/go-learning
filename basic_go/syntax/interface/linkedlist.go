package main

type LinkedList struct {
	head *node
	tail *node
}

// 实现了接口的所有方法，就是实现了这个接口
func (l LinkedList) Add(index int, val any) {
	println("linkedlist add method ...")
}

func (l LinkedList) Append(val any) error {
	println("linkedlist append method ...")
	return nil
}

func (l LinkedList) Delete(index int) (any, error) {
	println("linkedlist delete method ...")
	return nil, nil
}

type node struct {
}
