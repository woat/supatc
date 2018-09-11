package gui

import (
	"github.com/andlabs/ui"
)

func addTasksInput() {
	window := ui.NewWindow("supatc - Add Tasks", 300, 600, false)
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

	vb := ui.NewVerticalBox()
	vb.SetPadded(true)
	checkout := ui.NewRadioButtons()
	checkout.Append("Lightning")
	checkout.Append("Anti-Bot")
	// TODO bug label shows bottom button instead of top
	vb.Append(checkout, false)
	group.Append("Checkout Type", vb, false)

	group.SetChild(form)
	vbox.Append(group, true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

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
	vbox.Append(hbox, false)

	window.SetMargined(true)
	window.SetChild(vbox)
	window.OnClosing(func(*ui.Window) bool {
		return true
	})
	window.Show()
}
