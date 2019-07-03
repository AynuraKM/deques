package main

import "fmt"

func main() {
	dq := dequeue{}
	for i := 0; i < 10; i++ {
		dq.add_rear(i)
	}
	fmt.Println(dq)
}

type dequeue struct {
	slice []interface{}
}
func (dq *dequeue) add_rear(item interface{}) {
	dq.slice[len(dq.slice)+1] = append(dq.slice[:len(dq.slice)+1], item)
}
