package gui

import (
	"github.com/andlabs/ui"
)

func addTasksInput() {
	label := ui.NewLabel("Add Tasks Pop-up Page")
	vbox := ui.NewVerticalBox()
	vbox.Append(label, false)

	window := ui.NewWindow("supatc - Add Tasks", 300, 300, false)
	window.SetMargined(true)
	window.SetChild(vbox)
	window.OnClosing(func(*ui.Window) bool {
		return true
	})
	window.Show()
}
