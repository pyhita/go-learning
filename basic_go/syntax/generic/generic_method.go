package main

func Sum[T Number](data ...T) T {
	var res T
	for _, v := range data {
		res += v
	}
	return res
}

type Number interface {
	int | uint
}
