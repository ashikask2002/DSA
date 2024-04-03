package main

import "fmt"

const size = 26

type Node struct {
	children [size]*Node
	isend    bool
}

type trie struct {
	root *Node
}

func (t *trie) Insert(value string) {
	n := len(value)
	current := t.root
	for i := 0; i < n; i++ {
		curnInds := value[i] - 'a'
		if current.children[curnInds] == nil {
			current.children[curnInds] = &Node{}
		}
		current = current.children[curnInds]
	}
	current.isend = true

}

func (t *trie) Search(value string) bool {
	n := len(value)
	currentnode := t.root
	for i := 0; i < n; i++ {
		curnInds := value[i] - 'a'
		if currentnode.children[curnInds] == nil {
			return false
		}
		currentnode = currentnode.children[curnInds]
	}
	if currentnode.isend == true {
		return true
	}
	return false
}

func main() {
	t := &trie{root: &Node{}}
	toAdd := []string{
		"ashik",
		"ashi",
		"ash",
		"sdj",
	}
	for _, val := range toAdd {
		t.Insert(val)
	}
	b := t.Search("ashik")
	fmt.Println(" ", b)
	j := t.Search("h")
	fmt.Println(" ",j)
}
