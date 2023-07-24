package chart

import crt "go.devnw.com/ooxml/schema/soo/dml/chart"
import "go.devnw.com/ooxml"

// PieChart is a Pie chart.
type PieChart struct {
	chartBase
	x *crt.CT_PieChart
}

// X returns the inner wrapped XML type.
func (c PieChart) X() *crt.CT_PieChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c PieChart) InitializeDefaults() {
	c.x.VaryColors = crt.NewCT_Boolean()
	c.x.VaryColors.ValAttr = office.Bool(true)

}

// AddSeries adds a default series to an Pie chart.
func (c PieChart) AddSeries() PieChartSeries {
	ser := crt.NewCT_PieSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := PieChartSeries{ser}
	bs.InitializeDefaults()
	return bs
}
