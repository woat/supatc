package gui

import (
	"github.com/andlabs/ui"
)

type modelHandler struct{}

func (mh *modelHandler) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{
		ui.TableString(""),
		ui.TableString(""),
		ui.TableString(""),
		ui.TableString(""),
		ui.TableString(""),
	}
}

func (mh *modelHandler) NumRows(m *ui.TableModel) int {
	return 5
}

func (mh *modelHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if column == 0 {
		return ui.TableString("owo")
	}
	return ui.TableString("default")
}

func (mh *modelHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {

}

func newModelHandler() *modelHandler {
	return new(modelHandler)
}

func newTasksTable() *ui.Table {
	h := newModelHandler()
	m := ui.NewTableModel(h)
	t := ui.NewTable(&ui.TableParams{
		Model:                         m,
		RowBackgroundColorModelColumn: -1,
	})

	t.AppendTextColumn("ID", 0, ui.TableModelColumnNeverEditable, nil)
	t.AppendTextColumn("Product", 1, ui.TableModelColumnNeverEditable, nil)
	t.AppendTextColumn("Size", 2, ui.TableModelColumnNeverEditable, nil)
	t.AppendTextColumn("Profile", 3, ui.TableModelColumnNeverEditable, nil)
	t.AppendTextColumn("Status", 4, ui.TableModelColumnNeverEditable, nil)
	return t
}

/*
	The time of writing the following is at 01:48 and I am unsure about
	what I have done.

	Given Go's package scope, this might be frowned upon, but it feels
	wrong to be able to call functions in different files. So in this
	case I am using the `_` convention to notate when calling a method
	not found in it's immediate file. Knowing this, I create a function
	referring to it's filename to then return a struct that contains
	methods that would be used elsewhere.

	Why not a package?

	In this specific case, notice how this file is no more than just
	utility/helper. I don't think it would be necessary to put it within
	it's own package just yet because I do not forsee enough complexity
	to dictate that.

	I'm not sure.
*/

type _tt struct{}

func (_t _tt) newTasksTable() *ui.Table {
	return newTasksTable()
}

func _taskstables() *_tt {
	return new(_tt)
}
