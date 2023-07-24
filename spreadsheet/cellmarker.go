package spreadsheet

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/measurement"
	sd "go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

// CellMarker represents a cell position
type CellMarker struct {
	x *sd.CT_Marker
}

// X returns the inner wrapped XML type.
func (c CellMarker) X() *sd.CT_Marker {
	return c.x
}

// Col returns the column of the cell marker.
func (c CellMarker) Col() int32 {
	return c.x.Col
}

// SetCol set the column of the cell marker.
func (c CellMarker) SetCol(col int32) {
	c.x.Col = col
}

// Row returns the row of the cell marker.
func (c CellMarker) Row() int32 {
	return c.x.Row
}

// SetRow set the row of the cell marker.
func (c CellMarker) SetRow(row int32) {
	c.x.Row = row
}

// SetColOffset sets a column offset in absolute distance.
func (c CellMarker) SetColOffset(m measurement.Distance) {
	c.x.ColOff.ST_CoordinateUnqualified = office.Int64(int64(m / measurement.EMU))
}

// ColOffset returns the offset from the row cell.
func (c CellMarker) ColOffset() measurement.Distance {
	if c.x.RowOff.ST_CoordinateUnqualified == nil {
		return 0
	}
	return measurement.Distance(float64(*c.x.ColOff.ST_CoordinateUnqualified) * measurement.EMU)
}

// SetRowOffset sets a column offset in absolute distance.
func (c CellMarker) SetRowOffset(m measurement.Distance) {
	c.x.RowOff.ST_CoordinateUnqualified = office.Int64(int64(m / measurement.EMU))
}

// RowOffset returns the offset from the row cell.
func (c CellMarker) RowOffset() measurement.Distance {
	if c.x.RowOff.ST_CoordinateUnqualified == nil {
		return 0
	}
	return measurement.Distance(float64(*c.x.RowOff.ST_CoordinateUnqualified) * measurement.EMU)
}
