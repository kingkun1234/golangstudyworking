package main

func main() {

}

type node struct {
	value      interface{}
	next, prev *node
}

type LinkedQueue struct {
	head, tail *node
	size       int
}

func (queue *LinkedQueue) Size() int {
	return queue.size
}

func (queue *LinkedQueue) Peek() interface{} {
	if queue.head == nil {
		panic("empty queue")
	}
	return queue.head.value
}

func (queue *LinkedQueue) Add(value interface{}) {
	new_node := &node{value, queue.tail, nil}
	if queue.tail == nil {
		queue.head = new_node
		queue.tail = new_node
	} else {
		queue.tail.next = new_node
		queue.tail = new_node
	}
	queue.size++
	new_node = nil
}

func (queue *LinkedQueue) Remove() {
	if queue.head == nil {
		panic("empty queue")
	}
	first_node := queue.head
	queue.head = first_node.next
	first_node.next = nil
	first_node.value = nil
	queue.size--
	first_node = nil
}
