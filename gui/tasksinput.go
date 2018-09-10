package gui

import (
	"github.com/andlabs/ui"
)

func addTasksInput() {
	window := ui.NewWindow("supatc - Add Tasks", 300, 300, false)
	vbox := ui.NewVerticalBox()
	group := ui.NewGroup("Add new tasks, separate keywords with ;")
	group.SetMargined(true)

	form := ui.NewForm()
	form.SetPadded(true)
	form.Append("Product keywords", ui.NewEntry(), false)
	form.Append("Color keywords", ui.NewEntry(), false)
	form.Append("Proxy", ui.NewEntry(), false)

	size := ui.NewCombobox()
	size.Append("S")
	size.Append("M")
	size.Append("L")
	form.Append("Size", size, false)

	profile := ui.NewCombobox()
	profile.Append("TODO: Fill this with profiles")
	form.Append("Profile", profile, false)

	checkout := ui.NewRadioButtons()
	checkout.Append("Lightning")
	checkout.Append("Anti-Bot")
	// TODO bug label shows bottom button instead of top
	form.Append("Checkout Type", checkout, false)

	group.SetChild(form)
	vbox.Append(group, true)

	hbox := ui.NewHorizontalBox()
	vbox.Append(hbox, false)

	submit := ui.NewButton("Add")
	submit.OnClicked(func(*ui.Button) {
		window.Destroy()
	})

	cancel := ui.NewButton("Cancel")
	cancel.OnClicked(func(*ui.Button) {
		window.Destroy()
	})

	hbox.Append(submit, false)
	hbox.Append(cancel, false)

	window.SetMargined(true)
	window.SetChild(vbox)
	window.OnClosing(func(*ui.Window) bool {
		return true
	})
	window.Show()
}
