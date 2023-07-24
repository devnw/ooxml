package chart

import (
	"go.devnw.com/ooxml/drawing"
	"go.devnw.com/ooxml/schema/soo/dml"
	crt "go.devnw.com/ooxml/schema/soo/dml/chart"
)

// AreaChartSeries is a series to be used on an area chart.
type AreaChartSeries struct {
	x *crt.CT_AreaSer
}

// X returns the inner wrapped XML type.
func (c AreaChartSeries) X() *crt.CT_AreaSer {
	return c.x
}

// InitializeDefaults initializes an area series to the default values.
func (c AreaChartSeries) InitializeDefaults() {
}

// SetText sets the series text.
func (c AreaChartSeries) SetText(s string) {
	c.x.Tx = crt.NewCT_SerTx()
	c.x.Tx.Choice.V = &s
}

// CategoryAxis returns the category data source.
func (c AreaChartSeries) CategoryAxis() CategoryAxisDataSource {
	if c.x.Cat == nil {
		c.x.Cat = crt.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(c.x.Cat)
}

// Values returns the value data source.
func (c AreaChartSeries) Values() NumberDataSource {
	if c.x.Val == nil {
		c.x.Val = crt.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(c.x.Val)
}

// Properties returns the bar chart series shape properties.
func (c AreaChartSeries) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}
