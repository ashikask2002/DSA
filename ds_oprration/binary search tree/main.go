package main

import (
	"fmt"
	"math"
)

type treenode struct {
	value       int
	Left, Right *treenode
}

func (t *treenode) Insert(value int) *treenode {
	if t == nil {
		return &treenode{value: value}
	}
	if value <= t.value {
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
	if t.value == value {
		return true
	} else if value < t.value {
		return t.Left.Search(value)
	} else {
		return t.Right.Search(value)
	}
}

func minvalue(t *treenode) int {
	for t.Left != nil {
		t = t.Left
	}
	return t.value
}

func (t *treenode) Delete(value int) *treenode {
	if t == nil {
		return nil
	}
	if value < t.value {
		t.Left = t.Left.Delete(value)
	} else if value > t.value {
		t.Right = t.Right.Delete(value)
	} else {
		if t.Left == nil {
			return t.Right
		} else if t.Right == nil {
			return t.Left
		}
		t.value = minvalue(t.Right)
		t.Right = t.Right.Delete(t.value)
	}
	return t
}

func (t *treenode) findclosest(target int) int {
	closest := t.value
	current := t

	for current != nil {
		if abs(target-closest) > abs(target-current.value) {
			closest = current.value
		}
		if target < current.value {
			current = current.Left
		} else if target > current.value {
			current = current.Right
		} else {
			break
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (t *treenode) validate() bool {
	return t.validatehelper(math.MinInt, math.MaxInt)
}

func (t *treenode) validatehelper(MinInt, MaxInt int) bool {
	if t == nil {
		return true
	}
	if t.value < MinInt || t.value > MaxInt {
		return false
	}

	return t.Left.validatehelper(MinInt, t.value) && t.Right.validatehelper(t.value,MaxInt)
}

func (t *treenode) DisplayIn() {
	if t != nil {
		t.Left.DisplayIn()
		fmt.Printf("%d ", t.value)
		t.Right.DisplayIn()
	}
}

func main() {
	root := &treenode{value: 10}
	root.Insert(4)
	root.Insert(2)
	root.Insert(8)
	root.Insert(1)
	root.Insert(5)
	root.Insert(2)
	root.Insert(7)

	root.DisplayIn()

	todelete := 5
	root.Delete(todelete)
	fmt.Printf("array after deleting %d\n", todelete)
	root.DisplayIn()

	fmt.Println("min", minvalue(root))
	fmt.Println()
    
	target := 9
	closest := root.findclosest(target)
	fmt.Println("closest value is ",closest)

	fmt.Println("is the tree is validater ",root.validate() )
}
