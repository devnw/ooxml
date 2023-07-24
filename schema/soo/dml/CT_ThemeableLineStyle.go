package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_ThemeableLineStyle struct {
	Ln    *CT_LineProperties
	LnRef *CT_StyleMatrixReference
}

func NewCT_ThemeableLineStyle() *CT_ThemeableLineStyle {
	ret := &CT_ThemeableLineStyle{}
	return ret
}

func (m *CT_ThemeableLineStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Ln != nil {
		seln := xml.StartElement{Name: xml.Name{Local: "a:ln"}}
		e.EncodeElement(m.Ln, seln)
	}
	if m.LnRef != nil {
		selnRef := xml.StartElement{Name: xml.Name{Local: "a:lnRef"}}
		e.EncodeElement(m.LnRef, selnRef)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ThemeableLineStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ThemeableLineStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "ln"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "ln"}:
				m.Ln = NewCT_LineProperties()
				if err := d.DecodeElement(m.Ln, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lnRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lnRef"}:
				m.LnRef = NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.LnRef, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_ThemeableLineStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ThemeableLineStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ThemeableLineStyle and its children
func (m *CT_ThemeableLineStyle) Validate() error {
	return m.ValidateWithPath("CT_ThemeableLineStyle")
}

// ValidateWithPath validates the CT_ThemeableLineStyle and its children, prefixing error messages with path
func (m *CT_ThemeableLineStyle) ValidateWithPath(path string) error {
	if m.Ln != nil {
		if err := m.Ln.ValidateWithPath(path + "/Ln"); err != nil {
			return err
		}
	}
	if m.LnRef != nil {
		if err := m.LnRef.ValidateWithPath(path + "/LnRef"); err != nil {
			return err
		}
	}
	return nil
}
