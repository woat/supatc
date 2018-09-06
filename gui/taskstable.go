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

func (mh *modelHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {}

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
