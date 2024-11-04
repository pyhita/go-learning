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

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("word already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dict map[string]string

func (d Dict) Search(key string) (string, error) {
	_, ok := d[key]

	if ok {
		return d[key], nil
	} else {
		return "", ErrNotFound
	}
}

func (d Dict) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dict) Update(key, value string) {
	d[key] = value
}
