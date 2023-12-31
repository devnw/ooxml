package diagram

import (
	"encoding/xml"
	"fmt"
)

// ST_ParameterVal is a union type
type ST_ParameterVal struct {
	ST_DiagramHorizontalAlignment ST_DiagramHorizontalAlignment
	ST_VerticalAlignment          ST_VerticalAlignment
	ST_ChildDirection             ST_ChildDirection
	ST_ChildAlignment             ST_ChildAlignment
	ST_SecondaryChildAlignment    ST_SecondaryChildAlignment
	ST_LinearDirection            ST_LinearDirection
	ST_SecondaryLinearDirection   ST_SecondaryLinearDirection
	ST_StartingElement            ST_StartingElement
	ST_BendPoint                  ST_BendPoint
	ST_ConnectorRouting           ST_ConnectorRouting
	ST_ArrowheadStyle             ST_ArrowheadStyle
	ST_ConnectorDimension         ST_ConnectorDimension
	ST_RotationPath               ST_RotationPath
	ST_CenterShapeMapping         ST_CenterShapeMapping
	ST_NodeHorizontalAlignment    ST_NodeHorizontalAlignment
	ST_NodeVerticalAlignment      ST_NodeVerticalAlignment
	ST_FallbackDimension          ST_FallbackDimension
	ST_TextDirection              ST_TextDirection
	ST_PyramidAccentPosition      ST_PyramidAccentPosition
	ST_PyramidAccentTextMargin    ST_PyramidAccentTextMargin
	ST_TextBlockDirection         ST_TextBlockDirection
	ST_TextAnchorHorizontal       ST_TextAnchorHorizontal
	ST_TextAnchorVertical         ST_TextAnchorVertical
	ST_DiagramTextAlignment       ST_DiagramTextAlignment
	ST_AutoTextRotation           ST_AutoTextRotation
	ST_GrowDirection              ST_GrowDirection
	ST_FlowDirection              ST_FlowDirection
	ST_ContinueDirection          ST_ContinueDirection
	ST_Breakpoint                 ST_Breakpoint
	ST_Offset                     ST_Offset
	ST_HierarchyAlignment         ST_HierarchyAlignment
	Int32                         *int32
	Float64                       *float64
	Bool                          *bool
	StringVal                     *string
	ST_ConnectorPoint             ST_ConnectorPoint
}

func (m *ST_ParameterVal) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ParameterVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_DiagramHorizontalAlignment != ST_DiagramHorizontalAlignmentUnset {
		e.EncodeToken(xml.CharData(m.ST_DiagramHorizontalAlignment.String()))
	}
	if m.ST_VerticalAlignment != ST_VerticalAlignmentUnset {
		e.EncodeToken(xml.CharData(m.ST_VerticalAlignment.String()))
	}
	if m.ST_ChildDirection != ST_ChildDirectionUnset {
		e.EncodeToken(xml.CharData(m.ST_ChildDirection.String()))
	}
	if m.ST_ChildAlignment != ST_ChildAlignmentUnset {
		e.EncodeToken(xml.CharData(m.ST_ChildAlignment.String()))
	}
	if m.ST_SecondaryChildAlignment != ST_SecondaryChildAlignmentUnset {
		e.EncodeToken(xml.CharData(m.ST_SecondaryChildAlignment.String()))
	}
	if m.ST_LinearDirection != ST_LinearDirectionUnset {
		e.EncodeToken(xml.CharData(m.ST_LinearDirection.String()))
	}
	if m.ST_SecondaryLinearDirection != ST_SecondaryLinearDirectionUnset {
		e.EncodeToken(xml.CharData(m.ST_SecondaryLinearDirection.String()))
	}
	if m.ST_StartingElement != ST_StartingElementUnset {
		e.EncodeToken(xml.CharData(m.ST_StartingElement.String()))
	}
	if m.ST_BendPoint != ST_BendPointUnset {
		e.EncodeToken(xml.CharData(m.ST_BendPoint.String()))
	}
	if m.ST_ConnectorRouting != ST_ConnectorRoutingUnset {
		e.EncodeToken(xml.CharData(m.ST_ConnectorRouting.String()))
	}
	if m.ST_ArrowheadStyle != ST_ArrowheadStyleUnset {
		e.EncodeToken(xml.CharData(m.ST_ArrowheadStyle.String()))
	}
	if m.ST_ConnectorDimension != ST_ConnectorDimensionUnset {
		e.EncodeToken(xml.CharData(m.ST_ConnectorDimension.String()))
	}
	if m.ST_RotationPath != ST_RotationPathUnset {
		e.EncodeToken(xml.CharData(m.ST_RotationPath.String()))
	}
	if m.ST_CenterShapeMapping != ST_CenterShapeMappingUnset {
		e.EncodeToken(xml.CharData(m.ST_CenterShapeMapping.String()))
	}
	if m.ST_NodeHorizontalAlignment != ST_NodeHorizontalAlignmentUnset {
		e.EncodeToken(xml.CharData(m.ST_NodeHorizontalAlignment.String()))
	}
	if m.ST_NodeVerticalAlignment != ST_NodeVerticalAlignmentUnset {
		e.EncodeToken(xml.CharData(m.ST_NodeVerticalAlignment.String()))
	}
	if m.ST_FallbackDimension != ST_FallbackDimensionUnset {
		e.EncodeToken(xml.CharData(m.ST_FallbackDimension.String()))
	}
	if m.ST_TextDirection != ST_TextDirectionUnset {
		e.EncodeToken(xml.CharData(m.ST_TextDirection.String()))
	}
	if m.ST_PyramidAccentPosition != ST_PyramidAccentPositionUnset {
		e.EncodeToken(xml.CharData(m.ST_PyramidAccentPosition.String()))
	}
	if m.ST_PyramidAccentTextMargin != ST_PyramidAccentTextMarginUnset {
		e.EncodeToken(xml.CharData(m.ST_PyramidAccentTextMargin.String()))
	}
	if m.ST_TextBlockDirection != ST_TextBlockDirectionUnset {
		e.EncodeToken(xml.CharData(m.ST_TextBlockDirection.String()))
	}
	if m.ST_TextAnchorHorizontal != ST_TextAnchorHorizontalUnset {
		e.EncodeToken(xml.CharData(m.ST_TextAnchorHorizontal.String()))
	}
	if m.ST_TextAnchorVertical != ST_TextAnchorVerticalUnset {
		e.EncodeToken(xml.CharData(m.ST_TextAnchorVertical.String()))
	}
	if m.ST_DiagramTextAlignment != ST_DiagramTextAlignmentUnset {
		e.EncodeToken(xml.CharData(m.ST_DiagramTextAlignment.String()))
	}
	if m.ST_AutoTextRotation != ST_AutoTextRotationUnset {
		e.EncodeToken(xml.CharData(m.ST_AutoTextRotation.String()))
	}
	if m.ST_GrowDirection != ST_GrowDirectionUnset {
		e.EncodeToken(xml.CharData(m.ST_GrowDirection.String()))
	}
	if m.ST_FlowDirection != ST_FlowDirectionUnset {
		e.EncodeToken(xml.CharData(m.ST_FlowDirection.String()))
	}
	if m.ST_ContinueDirection != ST_ContinueDirectionUnset {
		e.EncodeToken(xml.CharData(m.ST_ContinueDirection.String()))
	}
	if m.ST_Breakpoint != ST_BreakpointUnset {
		e.EncodeToken(xml.CharData(m.ST_Breakpoint.String()))
	}
	if m.ST_Offset != ST_OffsetUnset {
		e.EncodeToken(xml.CharData(m.ST_Offset.String()))
	}
	if m.ST_HierarchyAlignment != ST_HierarchyAlignmentUnset {
		e.EncodeToken(xml.CharData(m.ST_HierarchyAlignment.String()))
	}
	if m.Int32 != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.Int32)))
	}
	if m.Float64 != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%f", *m.Float64)))
	}
	if m.Bool != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", b2i(*m.Bool))))
	}
	if m.StringVal != nil {
		e.EncodeToken(xml.CharData(*m.StringVal))
	}
	if m.ST_ConnectorPoint != ST_ConnectorPointUnset {
		e.EncodeToken(xml.CharData(m.ST_ConnectorPoint.String()))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_ParameterVal) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_DiagramHorizontalAlignment != ST_DiagramHorizontalAlignmentUnset {
		mems = append(mems, "ST_DiagramHorizontalAlignment")
	}
	if m.ST_VerticalAlignment != ST_VerticalAlignmentUnset {
		mems = append(mems, "ST_VerticalAlignment")
	}
	if m.ST_ChildDirection != ST_ChildDirectionUnset {
		mems = append(mems, "ST_ChildDirection")
	}
	if m.ST_ChildAlignment != ST_ChildAlignmentUnset {
		mems = append(mems, "ST_ChildAlignment")
	}
	if m.ST_SecondaryChildAlignment != ST_SecondaryChildAlignmentUnset {
		mems = append(mems, "ST_SecondaryChildAlignment")
	}
	if m.ST_LinearDirection != ST_LinearDirectionUnset {
		mems = append(mems, "ST_LinearDirection")
	}
	if m.ST_SecondaryLinearDirection != ST_SecondaryLinearDirectionUnset {
		mems = append(mems, "ST_SecondaryLinearDirection")
	}
	if m.ST_StartingElement != ST_StartingElementUnset {
		mems = append(mems, "ST_StartingElement")
	}
	if m.ST_BendPoint != ST_BendPointUnset {
		mems = append(mems, "ST_BendPoint")
	}
	if m.ST_ConnectorRouting != ST_ConnectorRoutingUnset {
		mems = append(mems, "ST_ConnectorRouting")
	}
	if m.ST_ArrowheadStyle != ST_ArrowheadStyleUnset {
		mems = append(mems, "ST_ArrowheadStyle")
	}
	if m.ST_ConnectorDimension != ST_ConnectorDimensionUnset {
		mems = append(mems, "ST_ConnectorDimension")
	}
	if m.ST_RotationPath != ST_RotationPathUnset {
		mems = append(mems, "ST_RotationPath")
	}
	if m.ST_CenterShapeMapping != ST_CenterShapeMappingUnset {
		mems = append(mems, "ST_CenterShapeMapping")
	}
	if m.ST_NodeHorizontalAlignment != ST_NodeHorizontalAlignmentUnset {
		mems = append(mems, "ST_NodeHorizontalAlignment")
	}
	if m.ST_NodeVerticalAlignment != ST_NodeVerticalAlignmentUnset {
		mems = append(mems, "ST_NodeVerticalAlignment")
	}
	if m.ST_FallbackDimension != ST_FallbackDimensionUnset {
		mems = append(mems, "ST_FallbackDimension")
	}
	if m.ST_TextDirection != ST_TextDirectionUnset {
		mems = append(mems, "ST_TextDirection")
	}
	if m.ST_PyramidAccentPosition != ST_PyramidAccentPositionUnset {
		mems = append(mems, "ST_PyramidAccentPosition")
	}
	if m.ST_PyramidAccentTextMargin != ST_PyramidAccentTextMarginUnset {
		mems = append(mems, "ST_PyramidAccentTextMargin")
	}
	if m.ST_TextBlockDirection != ST_TextBlockDirectionUnset {
		mems = append(mems, "ST_TextBlockDirection")
	}
	if m.ST_TextAnchorHorizontal != ST_TextAnchorHorizontalUnset {
		mems = append(mems, "ST_TextAnchorHorizontal")
	}
	if m.ST_TextAnchorVertical != ST_TextAnchorVerticalUnset {
		mems = append(mems, "ST_TextAnchorVertical")
	}
	if m.ST_DiagramTextAlignment != ST_DiagramTextAlignmentUnset {
		mems = append(mems, "ST_DiagramTextAlignment")
	}
	if m.ST_AutoTextRotation != ST_AutoTextRotationUnset {
		mems = append(mems, "ST_AutoTextRotation")
	}
	if m.ST_GrowDirection != ST_GrowDirectionUnset {
		mems = append(mems, "ST_GrowDirection")
	}
	if m.ST_FlowDirection != ST_FlowDirectionUnset {
		mems = append(mems, "ST_FlowDirection")
	}
	if m.ST_ContinueDirection != ST_ContinueDirectionUnset {
		mems = append(mems, "ST_ContinueDirection")
	}
	if m.ST_Breakpoint != ST_BreakpointUnset {
		mems = append(mems, "ST_Breakpoint")
	}
	if m.ST_Offset != ST_OffsetUnset {
		mems = append(mems, "ST_Offset")
	}
	if m.ST_HierarchyAlignment != ST_HierarchyAlignmentUnset {
		mems = append(mems, "ST_HierarchyAlignment")
	}
	if m.Int32 != nil {
		mems = append(mems, "Int32")
	}
	if m.Float64 != nil {
		mems = append(mems, "Float64")
	}
	if m.Bool != nil {
		mems = append(mems, "Bool")
	}
	if m.StringVal != nil {
		mems = append(mems, "StringVal")
	}
	if m.ST_ConnectorPoint != ST_ConnectorPointUnset {
		mems = append(mems, "ST_ConnectorPoint")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_ParameterVal) String() string {
	if m.ST_DiagramHorizontalAlignment != ST_DiagramHorizontalAlignmentUnset {
		return m.ST_DiagramHorizontalAlignment.String()
	}
	if m.ST_VerticalAlignment != ST_VerticalAlignmentUnset {
		return m.ST_VerticalAlignment.String()
	}
	if m.ST_ChildDirection != ST_ChildDirectionUnset {
		return m.ST_ChildDirection.String()
	}
	if m.ST_ChildAlignment != ST_ChildAlignmentUnset {
		return m.ST_ChildAlignment.String()
	}
	if m.ST_SecondaryChildAlignment != ST_SecondaryChildAlignmentUnset {
		return m.ST_SecondaryChildAlignment.String()
	}
	if m.ST_LinearDirection != ST_LinearDirectionUnset {
		return m.ST_LinearDirection.String()
	}
	if m.ST_SecondaryLinearDirection != ST_SecondaryLinearDirectionUnset {
		return m.ST_SecondaryLinearDirection.String()
	}
	if m.ST_StartingElement != ST_StartingElementUnset {
		return m.ST_StartingElement.String()
	}
	if m.ST_BendPoint != ST_BendPointUnset {
		return m.ST_BendPoint.String()
	}
	if m.ST_ConnectorRouting != ST_ConnectorRoutingUnset {
		return m.ST_ConnectorRouting.String()
	}
	if m.ST_ArrowheadStyle != ST_ArrowheadStyleUnset {
		return m.ST_ArrowheadStyle.String()
	}
	if m.ST_ConnectorDimension != ST_ConnectorDimensionUnset {
		return m.ST_ConnectorDimension.String()
	}
	if m.ST_RotationPath != ST_RotationPathUnset {
		return m.ST_RotationPath.String()
	}
	if m.ST_CenterShapeMapping != ST_CenterShapeMappingUnset {
		return m.ST_CenterShapeMapping.String()
	}
	if m.ST_NodeHorizontalAlignment != ST_NodeHorizontalAlignmentUnset {
		return m.ST_NodeHorizontalAlignment.String()
	}
	if m.ST_NodeVerticalAlignment != ST_NodeVerticalAlignmentUnset {
		return m.ST_NodeVerticalAlignment.String()
	}
	if m.ST_FallbackDimension != ST_FallbackDimensionUnset {
		return m.ST_FallbackDimension.String()
	}
	if m.ST_TextDirection != ST_TextDirectionUnset {
		return m.ST_TextDirection.String()
	}
	if m.ST_PyramidAccentPosition != ST_PyramidAccentPositionUnset {
		return m.ST_PyramidAccentPosition.String()
	}
	if m.ST_PyramidAccentTextMargin != ST_PyramidAccentTextMarginUnset {
		return m.ST_PyramidAccentTextMargin.String()
	}
	if m.ST_TextBlockDirection != ST_TextBlockDirectionUnset {
		return m.ST_TextBlockDirection.String()
	}
	if m.ST_TextAnchorHorizontal != ST_TextAnchorHorizontalUnset {
		return m.ST_TextAnchorHorizontal.String()
	}
	if m.ST_TextAnchorVertical != ST_TextAnchorVerticalUnset {
		return m.ST_TextAnchorVertical.String()
	}
	if m.ST_DiagramTextAlignment != ST_DiagramTextAlignmentUnset {
		return m.ST_DiagramTextAlignment.String()
	}
	if m.ST_AutoTextRotation != ST_AutoTextRotationUnset {
		return m.ST_AutoTextRotation.String()
	}
	if m.ST_GrowDirection != ST_GrowDirectionUnset {
		return m.ST_GrowDirection.String()
	}
	if m.ST_FlowDirection != ST_FlowDirectionUnset {
		return m.ST_FlowDirection.String()
	}
	if m.ST_ContinueDirection != ST_ContinueDirectionUnset {
		return m.ST_ContinueDirection.String()
	}
	if m.ST_Breakpoint != ST_BreakpointUnset {
		return m.ST_Breakpoint.String()
	}
	if m.ST_Offset != ST_OffsetUnset {
		return m.ST_Offset.String()
	}
	if m.ST_HierarchyAlignment != ST_HierarchyAlignmentUnset {
		return m.ST_HierarchyAlignment.String()
	}
	if m.Int32 != nil {
		return fmt.Sprintf("%v", *m.Int32)
	}
	if m.Float64 != nil {
		return fmt.Sprintf("%v", *m.Float64)
	}
	if m.Bool != nil {
		return fmt.Sprintf("%v", *m.Bool)
	}
	if m.StringVal != nil {
		return fmt.Sprintf("%v", *m.StringVal)
	}
	if m.ST_ConnectorPoint != ST_ConnectorPointUnset {
		return m.ST_ConnectorPoint.String()
	}
	return ""
}
