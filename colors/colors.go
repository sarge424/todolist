package colors

import "github.com/fatih/color"

var fg []color.Attribute = []color.Attribute{color.FgWhite, color.FgBlue, color.FgYellow, color.FgRed, color.FgGreen, color.FgBlack}
var bg []color.Attribute = []color.Attribute{color.BgWhite, color.BgBlue, color.BgYellow, color.BgRed, color.BgGreen, color.BgBlack}

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
