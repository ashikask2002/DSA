package main

import "fmt"

const arraysize = 7

type Hashtable struct {
	array [arraysize]*bucket
}

type bucket struct {
	head *bucketnode
}

type bucketnode struct {
	key      string
	nextnode *bucketnode
}

func (h *Hashtable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *Hashtable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *Hashtable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newnode := &bucketnode{key: k}
		newnode.nextnode = b.head
		b.head = newnode
	} else {
		fmt.Println("already exists")
	}
}

func (b *bucket) search(k string) bool {
	currentnode := b.head
	for currentnode != nil {
		if currentnode.key == k {
			return true
		}
		currentnode = currentnode.nextnode
	}
	return false

}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.nextnode
		return
	}
	previousnode := b.head
	for previousnode != nil {
		if previousnode.nextnode.key == k {
			previousnode.nextnode = previousnode.nextnode.nextnode
		}
		previousnode = previousnode.nextnode
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraysize
}

func Init() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashtable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",	
	}
	for _, v := range list {
		testHashtable.Insert(v)
	}
	fmt.Println(testHashtable.Search("ERIC"))

}
