package chart

import (
	"go.devnw.com/ooxml/color"
	"go.devnw.com/ooxml/drawing"
	"go.devnw.com/ooxml/measurement"
	"go.devnw.com/ooxml/schema/soo/dml"
	crt "go.devnw.com/ooxml/schema/soo/dml/chart"
)

type SurfaceChartSeries struct {
	x *crt.CT_SurfaceSer
}

// X returns the inner wrapped XML type.
func (c SurfaceChartSeries) X() *crt.CT_SurfaceSer {
	return c.x
}

// Index returns the index of the series
func (c SurfaceChartSeries) Index() uint32 {
	return c.x.Idx.ValAttr
}

// SetIndex sets the index of the series
func (c SurfaceChartSeries) SetIndex(idx uint32) {
	c.x.Idx.ValAttr = idx
}

// Order returns the order of the series
func (c SurfaceChartSeries) Order() uint32 {
	return c.x.Order.ValAttr
}

// SetOrder sets the order of the series
func (c SurfaceChartSeries) SetOrder(idx uint32) {
	c.x.Order.ValAttr = idx
}

// SetText sets the series text
func (c SurfaceChartSeries) SetText(s string) {
	c.x.Tx = crt.NewCT_SerTx()
	c.x.Tx.Choice.V = &s
}

// Properties returns the line chart series shape properties.
func (c SurfaceChartSeries) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}

func (c SurfaceChartSeries) CategoryAxis() CategoryAxisDataSource {
	if c.x.Cat == nil {
		c.x.Cat = crt.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(c.x.Cat)
}

func (c SurfaceChartSeries) Values() NumberDataSource {
	if c.x.Val == nil {
		c.x.Val = crt.NewCT_NumDataSource()
	}
	ns := MakeNumberDataSource(c.x.Val)
	// Mac Excel requires this on surface charts
	ns.CreateEmptyNumberCache()
	return ns
}

func (c SurfaceChartSeries) InitializeDefaults() {
	c.Properties().LineProperties().SetWidth(1 * measurement.Point)
	c.Properties().LineProperties().SetSolidFill(color.Black)
	c.Properties().LineProperties().SetJoin(drawing.LineJoinRound)
}
