package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_GvmlTextShapeChoice struct {
	UseSpRect *CT_GvmlUseShapeRectangle
	Xfrm      *CT_Transform2D
}

func NewCT_GvmlTextShapeChoice() *CT_GvmlTextShapeChoice {
	ret := &CT_GvmlTextShapeChoice{}
	return ret
}

func (m *CT_GvmlTextShapeChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UseSpRect != nil {
		seuseSpRect := xml.StartElement{Name: xml.Name{Local: "a:useSpRect"}}
		e.EncodeElement(m.UseSpRect, seuseSpRect)
	}
	if m.Xfrm != nil {
		sexfrm := xml.StartElement{Name: xml.Name{Local: "a:xfrm"}}
		e.EncodeElement(m.Xfrm, sexfrm)
	}
	return nil
}

func (m *CT_GvmlTextShapeChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GvmlTextShapeChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "useSpRect"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "useSpRect"}:
				m.UseSpRect = NewCT_GvmlUseShapeRectangle()
				if err := d.DecodeElement(m.UseSpRect, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "xfrm"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "xfrm"}:
				m.Xfrm = NewCT_Transform2D()
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_GvmlTextShapeChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlTextShapeChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlTextShapeChoice and its children
func (m *CT_GvmlTextShapeChoice) Validate() error {
	return m.ValidateWithPath("CT_GvmlTextShapeChoice")
}

// ValidateWithPath validates the CT_GvmlTextShapeChoice and its children, prefixing error messages with path
func (m *CT_GvmlTextShapeChoice) ValidateWithPath(path string) error {
	if m.UseSpRect != nil {
		if err := m.UseSpRect.ValidateWithPath(path + "/UseSpRect"); err != nil {
			return err
		}
	}
	if m.Xfrm != nil {
		if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
			return err
		}
	}
	return nil
}
