package chart

import (
	"go.devnw.com/ooxml/drawing"
	"go.devnw.com/ooxml/schema/soo/dml"
	crt "go.devnw.com/ooxml/schema/soo/dml/chart"
)

// BarChartSeries is a series to be used on a bar chart.
type BarChartSeries struct {
	x *crt.CT_BarSer
}

// X returns the inner wrapped XML type.
func (c BarChartSeries) X() *crt.CT_BarSer {
	return c.x
}

// InitializeDefaults initializes a bar chart series to the default values.
func (c BarChartSeries) InitializeDefaults() {
}

// SetText sets the series text.
func (c BarChartSeries) SetText(s string) {
	c.x.Tx = crt.NewCT_SerTx()
	c.x.Tx.Choice.V = &s
}

// CategoryAxis returns the category data source.
func (c BarChartSeries) CategoryAxis() CategoryAxisDataSource {
	if c.x.Cat == nil {
		c.x.Cat = crt.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(c.x.Cat)
}

// Values returns the value data source.
func (c BarChartSeries) Values() NumberDataSource {
	if c.x.Val == nil {
		c.x.Val = crt.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(c.x.Val)
}

// Properties returns the bar chart series shape properties.
func (c BarChartSeries) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}
