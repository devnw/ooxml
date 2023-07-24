package spreadsheet

import "go.devnw.com/ooxml/schema/soo/sml"

type Table struct {
	x *sml.Table
}

// X returns the inner wrapped XML type.
func (t Table) X() *sml.Table {
	return t.x
}

// Name returns the name of the table
func (t Table) Name() string {
	if t.x.NameAttr != nil {
		return *t.x.NameAttr
	}
	return ""
}

// Reference returns the table reference (the cells within the table)
func (t Table) Reference() string {
	return t.x.RefAttr
}
