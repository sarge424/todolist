package main

import (
	"fmt"
	"main/colors"
	"main/task"
	"os"
	"os/exec"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/fatih/color"
	"github.com/nathan-fiscaletti/consolesize-go"
)

var tasklist []task.Task
var w int
var h int
var line int

var mode int //0 for view, 1 for edit
var inp []byte

func init() {
	tasklist = make([]task.Task, 0, 5)
	inp = nil
	cursor.Hide()

	w, h = consolesize.GetConsoleSize()
	line = 0
}

func main() {
	refresh()

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if mode == 0 {
			if key.Code == keys.RuneKey {
				if key.String() == "q" {
					return true, nil //return true to exit
				}

			} else {
				if key.String() == "up" && line > 0 {
					line--
					refresh()
				} else if key.String() == "down" && line < len(tasklist) {
					line++
					refresh()
				} else if key.String() == "left" && line < len(tasklist) {
					tasklist[line].ShiftPriority(-1)
					refresh()
				} else if key.String() == "right" && line < len(tasklist) {
					tasklist[line].ShiftPriority(1)
					refresh()
				} else if key.String() == "space" && line < len(tasklist) {
					tasklist[line].Toggle()
					refresh()
				} else if key.String() == "backspace" && line < len(tasklist) {
					tasklist = append(tasklist[:line], tasklist[line+1:]...)
					refresh()
				} else if key.String() == "enter" {
					cursor.Up(len(tasklist) - line + 1)
					color.Set(color.FgBlack, color.BgWhite)
					fillLine()
					cursor.Left(w - 1)
					fmt.Printf("[ ]%d: ", line)
					cursor.Show()

					inp = nil
					mode = 1
				}
			}

		} else if mode == 1 {
			if key.Code == keys.RuneKey || key.String() == "space" || key.String() == "backspace" {
				if key.String() == "space" {
					fmt.Print(" ")
					inp = append(inp, byte(' '))
				} else if key.String() == "backspace" && len(inp) > 0 {
					cursor.Left(1)
					fmt.Print(" ")
					cursor.Left(1)
					inp = inp[:len(inp)-1]
				} else if key.Code == keys.RuneKey {
					fmt.Print(key.String())
					inp = append(inp, []byte(key.String())...)
				}
			} else {
				if key.String() == "enter" {
					if line < len(tasklist) {
						tasklist[line].SetText(string(inp))
					} else if line == len(tasklist) {
						if len(inp) > 0 {
							tasklist = append(tasklist, task.New(len(tasklist), string(inp), false))
						} else {
							line--
						}
					}

					color.Set(color.FgRed, color.BgBlack)
					mode = 0
					line++
					cursor.Hide()
					refresh()
				}
			}
		}

		return false, nil // Return false to continue listening
	})

}

func refresh() {
	cls()
	w, h = consolesize.GetConsoleSize()

	for l, t := range tasklist {
		colors.GetColor(t.GetColor(), l == line).Println(t.GetString(w))
	}
	colors.GetColor(0, line == len(tasklist)).Println("  +  new")
}

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func fillLine() {
	for i := 0; i < w-1; i++ {
		fmt.Print(" ")
	}
}
