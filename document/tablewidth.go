package document

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/measurement"
	"go.devnw.com/ooxml/schema/soo/wml"
)

// TableWidth controls width values in table settings.
type TableWidth struct {
	x *wml.CT_TblWidth
}

// NewTableWidth returns a newly intialized TableWidth
func NewTableWidth() TableWidth {
	return TableWidth{wml.NewCT_TblWidth()}
}

// X returns the inner wrapped XML type.
func (s TableWidth) X() *wml.CT_TblWidth {
	return s.x
}

// SetValue sets the width value.
func (s TableWidth) SetValue(m measurement.Distance) {
	s.x.WAttr = &wml.ST_MeasurementOrPercent{}
	s.x.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	s.x.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = office.Int64(int64(m / measurement.Twips))
	s.x.TypeAttr = wml.ST_TblWidthDxa
}
