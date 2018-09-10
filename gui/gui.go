package gui

import (
	"github.com/andlabs/ui"
)

var mainwin *ui.Window

func setupUI() {
	mainwin = ui.NewWindow("supatc", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tabs := ui.NewTab()

	tabs.Append("Tasks", makeTasksPage())
	tabs.SetMargined(0, true)
	tabs.Append("Billing", makeBillingPage())
	tabs.SetMargined(1, true)
	/*
		tabs.Append("Cards", makeAccountsPage())
		tabs.SetMargined(2, true)
		tabs.Append("Settings", makeSettingsPage)
		tabs.SetMargined(3, true)
	*/

	mainwin.SetChild(tabs)
	mainwin.SetMargined(true)

	mainwin.Show()
}

func makeTasksPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	add := ui.NewButton("Add")
	add.OnClicked(func(*ui.Button) {
		addTasksInput()
	})
	hbox.Append(add, false)
	hbox.Append(ui.NewButton("Run"), false)
	hbox.Append(ui.NewButton("Stop"), false)

	table := newTasksTable()

	vbox.Append(table, true)
	vbox.Append(hbox, false)
	return vbox
}

func makeBillingPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)
	return vbox
}

func Execute() {
	ui.Main(setupUI)
}
