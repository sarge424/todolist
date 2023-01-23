package tasklist

import (
	"main/task"

	"github.com/fatih/color"
)

var fg []color.Attribute = []color.Attribute{color.FgWhite, color.FgBlue, color.FgYellow, color.FgRed, color.FgGreen, color.FgBlack}
var bg []color.Attribute = []color.Attribute{color.BgWhite, color.BgBlue, color.BgYellow, color.BgRed, color.BgGreen, color.BgBlack}

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
		if node.prev != nil {
			node.prev.next = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		}
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

func (tl *Tasklist) Swap(node *Tasknode, up bool) bool {
	if node == nil {
		return false
	}

	if up && node.prev != nil {
		tmpt := node.prev.task
		tmpl := node.prev.sub

		node.prev.task = node.task
		node.prev.sub = node.sub

		node.task = tmpt
		node.sub = tmpl

		return true
	} else if !up && node.next != nil {
		tmpt := node.next.task
		tmpl := node.next.sub

		node.next.task = node.task
		node.next.sub = node.sub

		node.task = tmpt
		node.sub = tmpl

		return true
	}

	return false
}

func (tl *Tasklist) Nest(node *Tasknode) {
	if node != nil && node.prev != nil {
		tmp := node.prev
		tl.Del(node)
		if tmp.sub == nil {
			tmp.sub = New()
			tmp.sub.first = node

			node.prev = nil
			node.next = nil
		} else {
			b := tmp.sub.bottom()
			b.next = node

			node.prev = b
			node.next = nil
		}
	}
}

func (tl *Tasklist) At(index int) *task.Task {
	i := 0
	node := tl.first
	for node != nil {
		if i == index {
			return &node.task
		}

		if node.sub != nil {
			if i+node.sub.Len() >= index {
				return node.sub.At(index - i - 1)
			} else {
				i += node.sub.Len()
			}
		}

		if i == index {
			return &node.task
		}

		i++
		node = node.next
	}

	return &node.task
}

func (tl *Tasklist) NodeAt(index int) *Tasknode {
	i := 0
	node := tl.first
	for node != nil {
		if i == index {
			return node
		}

		if node.sub != nil {
			if i+node.sub.Len() >= index {
				return node.sub.NodeAt(index - i - 1)
			} else {
				i += node.sub.Len()
			}
		}

		if i == index {
			return node
		}

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

func (tl *Tasklist) bottom() *Tasknode {
	node := tl.first
	for node.next != nil {
		node = node.next
	}

	if node.sub != nil {
		return node.sub.bottom()
	}

	return node
}

func (tl *Tasklist) DeepDisplay(sel *Tasknode, w int, depth string) {
	tmp := tl.first
	for tmp != nil {
		selected := sel == tmp
		col := GetColor(tmp.task.GetColorIndex(), selected)
		col.Println(depth + tmp.task.GetString(w-len(depth)))
		if tmp.sub != nil {
			tmp.sub.DeepDisplay(sel, w, depth+depth)
		}
		tmp = tmp.next
	}
}

func GetColor(colIndex int, sel bool) *color.Color {
	if sel {
		return color.New(fg[5], bg[colIndex])
	} else {
		if colIndex == 0 {
			return color.New(fg[colIndex], bg[5], color.Bold)
		} else {
			return color.New(fg[colIndex], bg[5])
		}
	}
}

func (t *Tasknode) Disp() string {
	if t != nil {
		return t.task.GetText()
	}

	return ""
}

func (t *Tasknode) Lines() int {
	if t == nil || t.sub == nil {
		return 0
	} else {
		return t.sub.Len()
	}
}

func (t *Tasknode) GetPrev() *Tasknode {
	return t.prev
}

func (t *Tasknode) GetNext() *Tasknode {
	return t.next
}
