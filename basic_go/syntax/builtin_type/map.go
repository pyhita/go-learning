package main

func TestMap() {

	m1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}

	// 增
	m1[5] = 5

	// 删除
	delete(m1, 5)

	// 改
	m1[1] = 11

	// 查
	v, ok := m1[1]
	if ok {
		println(v)
	}

	// len
	println("len ", len(m1))
}
