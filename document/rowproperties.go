package document

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/measurement"
	"go.devnw.com/ooxml/schema/soo/ofc/sharedTypes"
	"go.devnw.com/ooxml/schema/soo/wml"
)

// RowProperties are the properties for a row within a table
type RowProperties struct {
	x *wml.CT_TrPr
}

// SetHeight allows controlling the height of a row within a table.
func (r RowProperties) SetHeight(ht measurement.Distance, rule wml.ST_HeightRule) {
	if rule == wml.ST_HeightRuleUnset {
		r.x.TrHeight = nil
	} else {
		htv := wml.NewCT_Height()
		htv.HRuleAttr = rule
		htv.ValAttr = &sharedTypes.ST_TwipsMeasure{}
		htv.ValAttr.ST_UnsignedDecimalNumber = office.Uint64(uint64(ht / measurement.Twips))
		r.x.TrHeight = []*wml.CT_Height{htv}
	}
}
