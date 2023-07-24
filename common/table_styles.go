package common

import (
	"go.devnw.com/ooxml/schema/soo/dml"
)

// TableStyles contains document specific properties.
type TableStyles struct {
	x *dml.TblStyleLst
}

// NewTableStyles constructs a new TableStyles.
func NewTableStyles() TableStyles {
	return TableStyles{x: dml.NewTblStyleLst()}
}

// X returns the inner wrapped XML type.
func (t TableStyles) X() *dml.TblStyleLst {
	return t.x
}

// DefAttr returns the DefAttr property.
func (t TableStyles) DefAttr() string {
	return t.x.DefAttr
}

// TblStyle returns the TblStyle property.
func (t TableStyles) TblStyle() []*dml.CT_TableStyle {
	return t.x.TblStyle
}
