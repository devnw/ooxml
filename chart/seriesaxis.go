package chart

import crt "go.devnw.com/ooxml/schema/soo/dml/chart"

type SeriesAxis struct {
	x *crt.CT_SerAx
}

func MakeSeriesAxis(x *crt.CT_SerAx) SeriesAxis {
	return SeriesAxis{x}
}

// X returns the inner wrapped XML type.
func (s SeriesAxis) X() *crt.CT_SerAx {
	return s.x
}

func (s SeriesAxis) InitializeDefaults() {

}

func (s SeriesAxis) AxisID() uint32 {
	return s.x.AxId.ValAttr
}

func (s SeriesAxis) SetCrosses(axis Axis) {
	s.x.CrossAx.ValAttr = axis.AxisID()
}
