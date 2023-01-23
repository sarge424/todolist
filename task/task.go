package task

import "fmt"

type Task struct {
	id       int
	priority int
	text     string
	done     bool
}

func New(id int, tx string, d bool) Task {
	t := Task{}
	t.id = id
	t.done = d
	t.priority = 0
	t.SetText(tx)

	return t
}

func (t *Task) SetText(tx string) {
	if len(tx) > 0 {
		t.text = tx
	}
}

func (t *Task) Toggle() {
	t.done = !t.done
}

func (t *Task) GetString(width int) string {
	c := ' '
	if t.done {
		c = 'x'
	}
	ans := fmt.Sprintf("[%c]%d: %s", c, t.id, t.text)

	for len(ans) < width {
		ans += " "
	}

	return string(ans[:width])
}

func (t *Task) GetColor() int {
	if t.done {
		return 4
	} else {
		return t.priority
	}
}

func (t *Task) ShiftPriority(offset int) {
	t.priority = (t.priority + offset + 4) % 4

}
