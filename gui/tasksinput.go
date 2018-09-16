package gui

import (
	"github.com/andlabs/ui"
)

func addTasksInput() {
	/*
		window
			groupcontainer
				formgroup
				radiogroup
				o
				o
				o
			buttonscontainer
				buttons
	*/

	window := ui.NewWindow("supatc - Add Tasks", 300, 600, false)
	window.SetMargined(true)
	window.OnClosing(func(*ui.Window) bool {
		return true
	})

	windowContainer := ui.NewVerticalBox()
	windowContainer.SetPadded(true)
	window.SetChild(windowContainer)

	formGroup := ui.NewGroup("tasks")
	formGroup.SetMargined(true)
	formGroupContainer := ui.NewVerticalBox()
	formGroup.SetChild(formGroupContainer)
	windowContainer.Append(formGroup, false)

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	formGroupContainer.Append(entryForm, false)

	entryForm.Append("Product keywords", ui.NewEntry(), false)
	entryForm.Append("Color keywords", ui.NewEntry(), false)
	entryForm.Append("Proxy", ui.NewEntry(), false)

	size := ui.NewCombobox()
	size.Append("S")
	size.Append("M")
	size.Append("L")
	entryForm.Append("Size", size, false)

	profile := ui.NewCombobox()
	profile.Append("TODO: Fill this with profiles")
	entryForm.Append("Profile", profile, false)

	checkout := ui.NewRadioButtons()
	checkout.Append("Lightning")
	checkout.Append("Anti-Bot")
	checkoutContainer := ui.NewVerticalBox()
	checkoutContainer.Append(checkout, false)
	formGroupContainer.Append(checkoutContainer, false)

	buttonsContainer := ui.NewHorizontalBox()
	buttonsContainer.SetPadded(true)

	submit := ui.NewButton("Add")
	submit.OnClicked(func(*ui.Button) {
		window.Destroy()
	})

	cancel := ui.NewButton("Cancel")
	cancel.OnClicked(func(*ui.Button) {
		window.Destroy()
	})

	buttonsContainer.Append(submit, false)
	buttonsContainer.Append(cancel, false)
	windowContainer.Append(buttonsContainer, false)

	window.Show()
}

// func addTasksInput() {
// 	window := ui.NewWindow("supatc - Add Tasks", 300, 600, false)
// 	vbox := ui.NewVerticalBox()
// 	group := ui.NewGroup("Add new tasks, separate keywords with ;")
// 	group.SetMargined(true)

// 	form := ui.NewForm()
// 	form.SetPadded(true)
// 	form.Append("Product keywords", ui.NewEntry(), false)
// 	form.Append("Color keywords", ui.NewEntry(), false)
// 	form.Append("Proxy", ui.NewEntry(), false)

// 	size := ui.NewCombobox()
// 	size.Append("S")
// 	size.Append("M")
// 	size.Append("L")
// 	form.Append("Size", size, false)

// 	profile := ui.NewCombobox()
// 	profile.Append("TODO: Fill this with profiles")
// 	form.Append("Profile", profile, false)

// 	checkout := ui.NewRadioButtons()
// 	checkout.Append("Lightning")
// 	checkout.Append("Anti-Bot")
// 	// TODO bug label shows bottom button instead of top
// 	form.Append("Checkout Type", checkout, false)

// 	profile = ui.NewCombobox()
// 	profile.Append("TODO: Fill this with profiles")
// 	form.Append("Profile", profile, false)

// 	group.SetChild(form)
// 	vbox.Append(group, true)

// 	hbox := ui.NewHorizontalBox()
// 	hbox.SetPadded(true)

// 	submit := ui.NewButton("Add")
// 	submit.OnClicked(func(*ui.Button) {
// 		window.Destroy()
// 	})

// 	cancel := ui.NewButton("Cancel")
// 	cancel.OnClicked(func(*ui.Button) {
// 		window.Destroy()
// 	})

// 	hbox.Append(submit, false)
// 	hbox.Append(cancel, false)
// 	vbox.Append(hbox, false)

// 	window.SetMargined(true)
// 	window.SetChild(vbox)
// 	window.OnClosing(func(*ui.Window) bool {
// 		return true
// 	})
// 	window.Show()
// }
