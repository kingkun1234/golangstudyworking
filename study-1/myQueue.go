package main

import "fmt"

type IQueue interface {
	InitQueue(queueCapacity int) //创建队列
	DestroyQueue()               //销毁队列
	ClearQueue()                 //清空队列
	QueueEmpty() bool            //判断队列
	QueueLength() int            //队列长度
	EnQueue(item int) bool       //插入队列
	DeQueue() bool               //元素出队
	QueueTraverse()              //遍历队列
}

type Queue struct {
	m_Queue       []int
	m_QueueLen    int
	m_QueueCap    int
	m_QueueHeader int
	m_QueueTail   int
}

func (queue *Queue) InitQueue(queueCapacity int) {
	queue.m_Queue = make([]int, queueCapacity, queueCapacity)
	queue.m_QueueLen = 0
	queue.m_QueueCap = queueCapacity
	queue.m_QueueHeader = 0
	queue.m_QueueTail = 0
}

func (queue *Queue) DestroyQueue() {
	queue = nil
}

func (queue *Queue) ClearQueue() {
	queue.m_QueueHeader = 0
	queue.m_QueueTail = 0
	queue.m_QueueLen = 0
}

func (queue *Queue) QueueEmpty() bool {
	return len(queue.m_Queue) == 0
}

func (queue *Queue) QueueLength() int {
	return len(queue.m_Queue)
}

func (queue *Queue) EnQueue(item int) bool {
	len := queue.m_QueueLen
	queue.m_Queue = append(queue.m_Queue, item)
	queue.m_QueueLen += 1
	queue.m_QueueCap -= 1
	if queue.m_QueueLen-len == 1 {
		return true
	}
	return false
}

func (queue *Queue) DeQueue() int {
	item := queue.m_Queue[0]
	queue.m_Queue = queue.m_Queue[1:len(queue.m_Queue)]
	queue.m_QueueLen -= 1
	queue.m_QueueCap += 1
	return item
}

func (queue *Queue) QueueTraverse() {
	for k, v := range queue.m_Queue {
		fmt.Println(v)
	}
}

var queue Queue

func New() *Queue {
	if queue.m_Queue == nil {
		queue = Queue{}
		queue.InitQueue(10)
	}
	return &queue
}
