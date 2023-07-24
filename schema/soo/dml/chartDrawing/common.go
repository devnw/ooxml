package chartDrawing

import "go.devnw.com/ooxml"

func b2i(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_ShapeNonVisual", NewCT_ShapeNonVisual)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Shape", NewCT_Shape)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_ConnectorNonVisual", NewCT_ConnectorNonVisual)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Connector", NewCT_Connector)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_PictureNonVisual", NewCT_PictureNonVisual)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Picture", NewCT_Picture)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GraphicFrameNonVisual", NewCT_GraphicFrameNonVisual)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GraphicFrame", NewCT_GraphicFrame)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GroupShapeNonVisual", NewCT_GroupShapeNonVisual)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GroupShape", NewCT_GroupShape)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Marker", NewCT_Marker)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_RelSizeAnchor", NewCT_RelSizeAnchor)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_AbsSizeAnchor", NewCT_AbsSizeAnchor)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Drawing", NewCT_Drawing)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "EG_ObjectChoices", NewEG_ObjectChoices)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "EG_Anchor", NewEG_Anchor)
}
