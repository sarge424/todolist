package tasklist

import "main/task"

type Tasklist struct {
	first *tasknode
}

type tasknode struct {
	task task.Task
	next *tasknode
}

func New() *Tasklist {
	tl := Tasklist{}
	tl.first = nil

	return &tl
}

func (tl *Tasklist) Append(t task.Task) {
	//make a node
	nn := tasknode{}
	nn.task = t
	nn.next = nil

	//connect the node
	if tl.first == nil {
		tl.first = &nn
	} else {
		tl.last().next = &nn
	}
}

func (tl *Tasklist) Del(index int) {
	if index == 0 {
		tl.first = tl.first.next
		return
	}

	node := tl.nodeAt(index)
	if node != nil {
		nx := node.next
		tl.nodeAt(index - 1).next = nx
	}
}

func (tl *Tasklist) Len() int {
	i := 0
	node := tl.first
	for node != nil {
		i++
		node = node.next
	}

	return i
}

func (tl *Tasklist) At(index int) *task.Task {
	i := 0
	node := tl.first
	for node != nil && i < index {
		i++
		node = node.next
	}

	return &node.task
}

func (tl *Tasklist) nodeAt(index int) *tasknode {
	i := 0
	node := tl.first
	for node != nil && i < index {
		i++
		node = node.next
	}

	return node
}

func (tl *Tasklist) last() *tasknode {
	node := tl.first
	for node.next != nil {
		node = node.next
	}

	return node
}
