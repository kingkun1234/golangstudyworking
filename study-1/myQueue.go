package main

import "fmt"

type QElemType int
type IQueue interface {
	InitQueue(queueCapacity int)     //创建队列
	DestroyQueue()                   //销毁队列
	ClearQueue()                     //清空队列
	QueueEmpty() bool                //判断队列
	QueueFull() bool                 //是否已满
	QueueLength() int                //队列长度
	EnQueue(element QElemType) bool  //插入队列
	DeQueue(element *QElemType) bool //元素出队
	QueueTraverse()                  //遍历队列
}

type Queue struct {
	m_Queue         []QElemType //队列
	m_QueueHeader   int         //队头
	m_QueueTail     int         //队尾
	m_QueueCapacity int         //容量
}

func (queue *Queue) InitQueue(queueCapacity int) {
	queue.m_Queue = make([]QElemType, queueCapacity, queueCapacity)
	queue.m_QueueHeader = 0
	queue.m_QueueTail = 0
	queue.m_QueueCapacity = queueCapacity
}

func (queue *Queue) DestroyQueue() {
	queue = nil
}

func (queue *Queue) ClearQueue() {
	queue.m_QueueHeader = 0
	queue.m_QueueTail = 0
}

func (queue *Queue) QueueEmpty() bool {
	return queue.QueueLength() == 0
}

func (queue *Queue) QueueLength() int {
	return (queue.m_QueueTail - queue.m_QueueHeader + queue.m_QueueCapacity) % queue.m_QueueCapacity
}

func (queue *Queue) EnQueue(element QElemType) bool {
	if queue.QueueFull() {
		return false
	}
	queue.m_Queue[queue.m_QueueTail] = element
	queue.m_QueueTail = (queue.m_QueueTail + 1) % queue.m_QueueCapacity
	return true
}

func (queue *Queue) DeQueue(element *QElemType) bool {
	if queue.QueueEmpty() {
		return false
	}
	element = &queue.m_Queue[queue.m_QueueHeader]
	queue.m_QueueHeader = (queue.m_QueueHeader + 1) % queue.m_QueueCapacity
	return true
}

func (queue *Queue) QueueTraverse() {
	for i := queue.m_QueueHeader; i < queue.m_QueueHeader+len(queue.m_Queue); i++ {
		fmt.Println(queue.m_Queue[i%queue.m_QueueCapacity])
	}
}

func (queue *Queue) QueueFull() bool {
	if (queue.m_QueueTail+1)%cap(queue.m_Queue) == queue.m_QueueHeader {
		return true
	}
	return false
}

var queue Queue

func New() *Queue {
	if queue.m_Queue == nil {
		queue = Queue{}
		queue.InitQueue(4)
	}
	return &queue
}

func main() {
	data := New()
	data.EnQueue(1)
	data.EnQueue(2)
	data.EnQueue(3)
	data.EnQueue(4)
	fmt.Println(data.QueueLength())
	fmt.Println(data.QueueEmpty())
	fmt.Println(data.QueueFull())
	fmt.Println(data.m_QueueHeader)
	fmt.Println(data.m_QueueTail)
	data.QueueTraverse()
	fmt.Println(len(data.m_Queue))
}
