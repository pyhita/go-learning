package main

type ArrayList struct {
}

func (a ArrayList) Add(index int, val any) {
	println("arraylist add method ...")
}

func (a ArrayList) Append(val any) error {
	println("arraylist append method ...")
	return nil
}

func (a ArrayList) Delete(index int) (any, error) {
	println("arraylist delete method ...")
	return nil, nil
}
