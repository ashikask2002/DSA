package main

import "fmt"

type queue struct {
	data []int
}

func main() {
l := &queue{}
l.enqueue(7)
l.enqueue(5)
l.enqueue(1)
l.enqueue(4)
l.enqueue(2)
l.display()
value := 4
l.removeelement(value)
l.display()
}

func (q *queue) enqueue(a int) {
	q.data = append(q.data, a)

}

func (q *queue) dequeue() int {
	toremove := q.data[0]
	q.data = q.data[1:]
	return toremove
}

func (q *queue) display() {
	fmt.Println(q.data)
}

func (q *queue) removeelement(val int){
	if len(q.data) ==0{
		return
	}
	var newarr []int
	for _,v := range q.data{
		if v != val {
			newarr = append(newarr,v)
		}
	}
	q.data = newarr
}
