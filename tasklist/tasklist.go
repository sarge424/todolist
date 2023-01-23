package tasklist

import "main/task"

type Tasklist struct {
	first *tasknode
}

type tasknode struct {
	task task.Task
	sub  *Tasklist
	prev *tasknode
	next *tasknode
}

func New() *Tasklist {
	tl := Tasklist{}
	return &tl
}

func (tl *Tasklist) Append(t task.Task) {
	//make a node
	nn := tasknode{}
	nn.task = t

	//connect the node
	if tl.first == nil {
		tl.first = &nn
	} else {
		tl.last().next = &nn
		nn.prev = tl.last()
	}
}

func (tl *Tasklist) Del(node *tasknode) {
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

func (tl *Tasklist) Swap(i1 int, i2 int) {
	if i1 < 0 || i2 < 0 || i1 >= tl.Len() || i2 >= tl.Len() {
		return
	}

	tmp := tl.nodeAt(i1).task
	tl.nodeAt(i1).task = tl.nodeAt(i2).task
	tl.nodeAt(i2).task = tmp
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
