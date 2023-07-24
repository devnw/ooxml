package spreadsheet

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/sml"
)

// SheetView is a view of a sheet. There is typically one per sheet, though more
// are supported.
type SheetView struct {
	x *sml.CT_SheetView
}

// X returns the inner wrapped XML type.
func (s SheetView) X() *sml.CT_SheetView {
	return s.x
}

// SetState sets the sheet view state (frozen/split/frozen-split)
func (s SheetView) SetState(st sml.ST_PaneState) {
	s.ensurePane()
	s.x.Pane.StateAttr = st
}

func (s SheetView) ensurePane() {
	if s.x.Pane == nil {
		s.x.Pane = sml.NewCT_Pane()
		s.x.Pane.ActivePaneAttr = sml.ST_PaneBottomLeft
	}
}

// SetXSplit sets the column split point
func (s SheetView) SetXSplit(v float64) {
	s.ensurePane()
	s.x.Pane.XSplitAttr = office.Float64(v)
}

// SetYSplit sets the row split point
func (s SheetView) SetYSplit(v float64) {
	s.ensurePane()
	s.x.Pane.YSplitAttr = office.Float64(v)

}

// SetTopLeft sets the top left visible cell after the split.
func (s SheetView) SetTopLeft(cellRef string) {
	s.ensurePane()
	s.x.Pane.TopLeftCellAttr = &cellRef
}

// SetZoom controls the zoom level of the sheet and is measured in percent. The
// default value is 100.
func (s SheetView) SetZoom(pct uint32) {
	s.x.ZoomScaleAttr = &pct
}

// SetShowRuler controls the visibility of the ruler
func (s SheetView) SetShowRuler(b bool) {
	// default is true
	if !b {
		s.x.ShowRulerAttr = office.Bool(false)
	} else {
		s.x.ShowRulerAttr = nil
	}
}
