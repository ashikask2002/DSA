package main

import "fmt"

type treenod struct {
	data        int
	Left, Right *treenod
}

func (t *treenod) Insert(data int) *treenod {
	if t == nil {
		return &treenod{data: data}
	}
	if data < t.data {
		t.Left = t.Left.Insert(data)
	} else {
		t.Right = t.Right.Insert(data)
	}
	return t
}

func (t *treenod) Search(data int) bool {
	if t == nil {
		return false
	}
	if data == t.data {
		return true
	} else if data < t.data {
		return t.Left.Search(data)
	} else {
		return t.Right.Search(data)
	}
}

func (t *treenod) DisplayIn() {
	if t != nil {
		t.Left.DisplayIn()
		fmt.Print(" ", t.data)
		t.Right.DisplayIn()
	}
}

func main() {
	l := &treenod{data: 10}
	l.Insert(4)
	l.Insert(7)
	l.Insert(9)
	l.Insert(2)
	l.Insert(10)
	l.Insert(101)

	l.DisplayIn()
	Searchvalue := 101
	fmt.Println("result of seardh is")
	fmt.Println(l.Search(Searchvalue))
}
