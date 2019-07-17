package main

import (
	"fmt"
)

func main() {
	dq := dequeue{}

	for i := 0; i < 10; i++ {
		dq.add_rear(i)
	}
	fmt.Println(dq)
	fmt.Println(dq.remove_rear)
	fmt.Println(dq.remove_front)
	dq.add_front("hello")
	fmt.Println(dq)
	fmt.Println(dq.is_empty)
	fmt.Println(dq.size)
	fmt.Println(dq.peek(true))
}

type dequeue struct {
	values []interface{}
}

func (dq *dequeue) add_rear(value interface{}) {
	dq.values = append(dq.values, value)
}

func (dq *dequeue) remove_rear() interface{} {
	val := dq.values[len(dq.values)-1]
	dq.values = dq.values[:len(dq.values)-1]
	return val
}

func (dq dequeue) add_front(values interface{}) {
	var n [100]interface{}
	for i:=0; i<len(dq.values) ; i++  {
		n[i+1]=dq.values[i]
	}
	dq.values = dq.values[:len(dq.values)+1]
	n[0] = values
	for i := 0; i < len(dq.values); i++ {
		dq.values[i] = n[i]
	}
}

func (dq *dequeue) remove_front() interface{} {
	val := dq.values[0]
	var n [100]interface{}
	for i := 1; i < len(dq.values); i++ {
		n[i-1] = dq.values[i]
	}
	dq.values = dq.values[:len(dq.values)-1]
	for i := 0; i < len(dq.values); i++ {
		dq.values[i] = n[i]
	}
	return val
}

func (dq *dequeue) is_empty() bool{
	if len(dq.values) == 0{
		return false
	}
	return true
}

func (dq *dequeue) size() int{
	l := len(dq.values)
	return l
}

func (dq *dequeue) peek(value bool) interface{} {
	var node interface{}
	if value == true {
		node = dq.values[len(dq.values)-1]
	} else {
		node = dq.values[0]
	}
	return node
}
