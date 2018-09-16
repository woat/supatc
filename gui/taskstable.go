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
	return 1
}

func (mh *modelHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if column == 0 {
		return ui.TableString("oo")
	}
	return ui.TableString("default")
}

func (mh *modelHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {}

func newModelHandler() *modelHandler {
	return new(modelHandler)
}

func newTasksTable() *ui.Table {
	modelHandler := newModelHandler()
	model := ui.NewTableModel(modelHandler)
	table := ui.NewTable(&ui.TableParams{
		Model:                         model,
		RowBackgroundColorModelColumn: -1,
	})

	table.AppendTextColumn("ID", 0, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Keywords", 1, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Size", 2, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Profile", 3, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Status", 4, ui.TableModelColumnNeverEditable, nil)
	return table
}
