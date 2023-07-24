package chart

import (
	"go.devnw.com/ooxml/drawing"

	"go.devnw.com/ooxml/schema/soo/dml"
	crt "go.devnw.com/ooxml/schema/soo/dml/chart"
)

type Marker struct {
	x *crt.CT_Marker
}

func MakeMarker(x *crt.CT_Marker) Marker {
	return Marker{x}
}

// X returns the inner wrapped XML type.
func (m Marker) X() *crt.CT_Marker {
	return m.x
}

func (m Marker) SetSymbol(s crt.ST_MarkerStyle) {
	if s == crt.ST_MarkerStyleUnset {
		m.x.Symbol = nil
	} else {
		m.x.Symbol = crt.NewCT_MarkerStyle()
		m.x.Symbol.ValAttr = s
	}
}
func (m Marker) SetSize(sz uint8) {
	m.x.Size = crt.NewCT_MarkerSize()
	m.x.Size.ValAttr = &sz
}
func (m Marker) Properties() drawing.ShapeProperties {
	if m.x.SpPr == nil {
		m.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(m.x.SpPr)
}
