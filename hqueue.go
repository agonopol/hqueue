package hqueue

import "container/heap"

type Less interface {
	Less(interface{}) bool
}

type Queue struct {
	data []Less
}

func NewQueue() *Queue {
	this := new(Queue)
	this.data = make([]Less, 0)
	return this
}

func (this *Queue) Len() int {
	return len(this.data)
}

func (this *Queue) Less(i, j int) bool {
	return this.data[i].Less(this.data[j])
}

func (this *Queue) Pop() interface{} {
	x := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return x
}

func (this *Queue) Push(i interface{}) {
	this.data = append(this.data, i.(Less))
}

func (this *Queue) Swap(i, j int) {
	tmp := this.data[i]
	this.data[i] = this.data[j]
	this.data[j] = tmp
}

type HQueue struct {
	queue *Queue
}

func New() *HQueue {
	this := &HQueue{NewQueue()}
	heap.Init(this.queue)
	return this
}

func (this *HQueue) Push(i Less) {
	heap.Push(this.queue,i)
}

func (this *HQueue) Pop() interface{} {
	return heap.Pop(this.queue)
}

func (this *HQueue) Peek() interface {} {
	if this.queue.Len() > 0 {
		return this.queue.data[0]
	}
	return nil
}

func (this *HQueue) Len() int {
	return this.queue.Len()
}
