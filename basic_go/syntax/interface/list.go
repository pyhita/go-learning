package main

type List interface {
	Add(index int, val any)
	Append(val any) error
	Delete(index int) (any, error)
}
