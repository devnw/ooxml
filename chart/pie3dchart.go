package chart

import crt "go.devnw.com/ooxml/schema/soo/dml/chart"
import "go.devnw.com/ooxml"

// Pie3DChart is a Pie3D chart.
type Pie3DChart struct {
	chartBase
	x *crt.CT_Pie3DChart
}

// X returns the inner wrapped XML type.
func (c Pie3DChart) X() *crt.CT_Pie3DChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c Pie3DChart) InitializeDefaults() {
	c.x.VaryColors = crt.NewCT_Boolean()
	c.x.VaryColors.ValAttr = office.Bool(true)
}

// AddSeries adds a default series to an Pie3D chart.
func (c Pie3DChart) AddSeries() PieChartSeries {
	ser := crt.NewCT_PieSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := PieChartSeries{ser}
	bs.InitializeDefaults()
	return bs
}
