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

	t := ui.NewTab()

	t.Append("Tasks", makeTasksPage())
	t.SetMargined(0, true)
	t.Append("Billing", makeBillingPage())
	t.SetMargined(1, true)
	/*
		t.Append("Cards", makeAccountsPage())
		t.SetMargined(2, true)
		t.Append("Settings", makeSettingsPage)
		t.SetMargined(3, true)
	*/

	mainwin.SetChild(t)
	mainwin.SetMargined(true)

	mainwin.Show()
}

// see taskstable.go
func makeTasksPage() ui.Control {
	v := ui.NewVerticalBox()
	v.SetPadded(true)

	h := ui.NewHorizontalBox()
	h.SetPadded(true)

	h.Append(ui.NewButton("Add"), false)
	h.Append(ui.NewButton("Run"), false)
	h.Append(ui.NewButton("Stop"), false)

	t := newTasksTable()

	v.Append(t, true)
	v.Append(h, false)
	return v
}

func makeBillingPage() ui.Control {
	v := ui.NewVerticalBox()
	v.SetPadded(true)

	h := ui.NewHorizontalBox()
	h.SetPadded(true)
	v.Append(h, false)
	return v
}

func Execute() {
	ui.Main(setupUI)
}
