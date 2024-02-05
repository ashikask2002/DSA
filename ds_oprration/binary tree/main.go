package main

import "fmt"

type treenode struct {
	Value       int
	Left, Right *treenode
}

func (t *treenode) Insert(value int) *treenode {
	if t == nil {
		return &treenode{Value: value}
	}
	if value <= t.Value {
		t.Left = t.Left.Insert(value)
	} else {
		t.Right = t.Right.Insert(value)
	}
	return t

}
func (t *treenode) Search(value int) bool {
	if t == nil {
		return false
	}
	if value == t.Value {
		return true
	} else if value <= t.Value {
		return t.Left.Search(value)
	} else {
		return t.Right.Search(value)
	}

}

func minvalue(t *treenode) int {
	for t.Left != nil {
		t = t.Left
	}
	return t.Value
}

func (t *treenode) DisplayIn() {
	if t != nil {
		t.Left.DisplayIn()
		fmt.Printf("%d ", t.Value)
		t.Right.DisplayIn()
	}
}

func (t *treenode) Displaypre() {
	if t != nil {
		fmt.Printf("%d ", t.Value)
		t.Left.Displaypre()
		t.Right.Displaypre()
	}
}

func (t *treenode) DisplayPost() {
	if t != nil {
		t.Left.DisplayPost()
		t.Right.DisplayPost()
		fmt.Printf("%d ", t.Value)
	}
}

func main() {
 root := &treenode{Value : 10}
 root.Insert(5)
 root.Insert(2)
 root.Insert(9)
 root.Insert(1)
 root.Insert(3)


 fmt.Println("in order traversal ")
 root.DisplayIn()
 fmt.Println()

 fmt.Println("pre order traversal")
 root.Displaypre()
 fmt.Println()

 fmt.Println("post order traversal")
 root.DisplayPost()
 fmt.Println()
}
