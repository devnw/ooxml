package spreadsheet

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/measurement"
	"go.devnw.com/ooxml/schema/soo/sml"
)

// Column represents a column within a sheet. It's only used for formatting
// purposes, so it's possible to construct a sheet without configuring columns.
type Column struct {
	x *sml.CT_Col
}

// X returns the inner wrapped XML type.
func (c Column) X() *sml.CT_Col {
	return c.x
}

// SetWidth controls the width of a column.
func (c Column) SetWidth(w measurement.Distance) {
	c.x.WidthAttr = office.Float64(float64(w / measurement.Character))
}

// SetStyle sets the cell style for an entire column.
func (c Column) SetStyle(cs CellStyle) {
	c.x.StyleAttr = office.Uint32(cs.Index())
}

// SetHidden controls the visibility of a column.
func (c Column) SetHidden(b bool) {
	if !b {
		c.x.HiddenAttr = nil
	} else {
		c.x.HiddenAttr = office.Bool(true)
	}
}
