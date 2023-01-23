package tasklist

import (
	"main/task"
)

type Tasklist struct {
	first *Tasknode
}

type Tasknode struct {
	task task.Task
	sub  *Tasklist
	prev *Tasknode
	next *Tasknode
}

func New() *Tasklist {
	tl := Tasklist{}
	return &tl
}

func (tl *Tasklist) Append(t task.Task) {
	//make a node
	nn := Tasknode{}
	nn.task = t

	//connect the node
	if tl.first == nil {
		tl.first = &nn
	} else {
		nn.prev = tl.last()
		tl.last().next = &nn
	}
}

func (tl *Tasklist) Del(node *Tasknode) {
	if node != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}

func (tl *Tasklist) Len() int {
	i := 0
	node := tl.first
	for node != nil {
		i++
		if node.sub != nil {
			i += node.sub.Len()
		}
		node = node.next
	}

	return i
}

func (tl *Tasklist) Swap(node *Tasknode, up bool) {
	if node == nil {
		return
	}

	if up && node.prev != nil {
		tmp := node.task
		node.task = node.prev.task
		node.prev.task = tmp

	} else if !up && node.next != nil {
		tmp := node.next.task
		node.next.task = node.task
		node.task = tmp
	}
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

func (tl *Tasklist) NodeAt(index int) *Tasknode {
	i := 0
	node := tl.first
	for node != nil && i < index {
		i++
		node = node.next
	}

	return node
}

func (tl *Tasklist) last() *Tasknode {
	node := tl.first
	for node.next != nil {
		node = node.next
	}

	return node
}
