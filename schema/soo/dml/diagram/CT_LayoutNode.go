package diagram

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_LayoutNode struct {
	NameAttr     *string
	StyleLblAttr *string
	ChOrderAttr  ST_ChildOrderType
	MoveWithAttr *string
	Alg          []*CT_Algorithm
	Shape        []*CT_Shape
	PresOf       []*CT_PresentationOf
	ConstrLst    []*CT_Constraints
	RuleLst      []*CT_Rules
	VarLst       []*CT_LayoutVariablePropertySet
	ForEach      []*CT_ForEach
	LayoutNode   []*CT_LayoutNode
	Choose       []*CT_Choose
	ExtLst       []*dml.CT_OfficeArtExtensionList
}

func NewCT_LayoutNode() *CT_LayoutNode {
	ret := &CT_LayoutNode{}
	return ret
}

func (m *CT_LayoutNode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.StyleLblAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "styleLbl"},
			Value: fmt.Sprintf("%v", *m.StyleLblAttr)})
	}
	if m.ChOrderAttr != ST_ChildOrderTypeUnset {
		attr, err := m.ChOrderAttr.MarshalXMLAttr(xml.Name{Local: "chOrder"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.MoveWithAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "moveWith"},
			Value: fmt.Sprintf("%v", *m.MoveWithAttr)})
	}
	e.EncodeToken(start)
	if m.Alg != nil {
		sealg := xml.StartElement{Name: xml.Name{Local: "alg"}}
		for _, c := range m.Alg {
			e.EncodeElement(c, sealg)
		}
	}
	if m.Shape != nil {
		seshape := xml.StartElement{Name: xml.Name{Local: "shape"}}
		for _, c := range m.Shape {
			e.EncodeElement(c, seshape)
		}
	}
	if m.PresOf != nil {
		sepresOf := xml.StartElement{Name: xml.Name{Local: "presOf"}}
		for _, c := range m.PresOf {
			e.EncodeElement(c, sepresOf)
		}
	}
	if m.ConstrLst != nil {
		seconstrLst := xml.StartElement{Name: xml.Name{Local: "constrLst"}}
		for _, c := range m.ConstrLst {
			e.EncodeElement(c, seconstrLst)
		}
	}
	if m.RuleLst != nil {
		seruleLst := xml.StartElement{Name: xml.Name{Local: "ruleLst"}}
		for _, c := range m.RuleLst {
			e.EncodeElement(c, seruleLst)
		}
	}
	if m.VarLst != nil {
		sevarLst := xml.StartElement{Name: xml.Name{Local: "varLst"}}
		for _, c := range m.VarLst {
			e.EncodeElement(c, sevarLst)
		}
	}
	if m.ForEach != nil {
		seforEach := xml.StartElement{Name: xml.Name{Local: "forEach"}}
		for _, c := range m.ForEach {
			e.EncodeElement(c, seforEach)
		}
	}
	if m.LayoutNode != nil {
		selayoutNode := xml.StartElement{Name: xml.Name{Local: "layoutNode"}}
		for _, c := range m.LayoutNode {
			e.EncodeElement(c, selayoutNode)
		}
	}
	if m.Choose != nil {
		sechoose := xml.StartElement{Name: xml.Name{Local: "choose"}}
		for _, c := range m.Choose {
			e.EncodeElement(c, sechoose)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		for _, c := range m.ExtLst {
			e.EncodeElement(c, seextLst)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LayoutNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
		if attr.Name.Local == "chOrder" {
			m.ChOrderAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "moveWith" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MoveWithAttr = &parsed
			continue
		}
		if attr.Name.Local == "styleLbl" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleLblAttr = &parsed
			continue
		}
	}
lCT_LayoutNode:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "alg"}:
				tmp := NewCT_Algorithm()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Alg = append(m.Alg, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "shape"}:
				tmp := NewCT_Shape()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Shape = append(m.Shape, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "presOf"}:
				tmp := NewCT_PresentationOf()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.PresOf = append(m.PresOf, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "constrLst"}:
				tmp := NewCT_Constraints()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ConstrLst = append(m.ConstrLst, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "ruleLst"}:
				tmp := NewCT_Rules()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RuleLst = append(m.RuleLst, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "varLst"}:
				tmp := NewCT_LayoutVariablePropertySet()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.VarLst = append(m.VarLst, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "forEach"}:
				tmp := NewCT_ForEach()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ForEach = append(m.ForEach, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "layoutNode"}:
				tmp := NewCT_LayoutNode()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LayoutNode = append(m.LayoutNode, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "choose"}:
				tmp := NewCT_Choose()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Choose = append(m.Choose, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				tmp := dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ExtLst = append(m.ExtLst, tmp)
			default:
				office.Log("skipping unsupported element on CT_LayoutNode %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LayoutNode
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LayoutNode and its children
func (m *CT_LayoutNode) Validate() error {
	return m.ValidateWithPath("CT_LayoutNode")
}

// ValidateWithPath validates the CT_LayoutNode and its children, prefixing error messages with path
func (m *CT_LayoutNode) ValidateWithPath(path string) error {
	if err := m.ChOrderAttr.ValidateWithPath(path + "/ChOrderAttr"); err != nil {
		return err
	}
	for i, v := range m.Alg {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Alg[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Shape {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Shape[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.PresOf {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/PresOf[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ConstrLst {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ConstrLst[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.RuleLst {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RuleLst[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.VarLst {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/VarLst[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ForEach {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ForEach[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.LayoutNode {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LayoutNode[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Choose {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choose[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ExtLst {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ExtLst[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
