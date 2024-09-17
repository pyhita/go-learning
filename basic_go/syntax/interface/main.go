package main

func main() {
	testList(LinkedList{})
	testList(ArrayList{})
}

// 测试接口
func testList(list List) {
	list.Add(0, 0)
}
