package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_RegularTextRun struct {
	RPr *CT_TextCharacterProperties
	T   string
}

func NewCT_RegularTextRun() *CT_RegularTextRun {
	ret := &CT_RegularTextRun{}
	return ret
}

func (m *CT_RegularTextRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "a:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	set := xml.StartElement{Name: xml.Name{Local: "a:t"}}
	office.AddPreserveSpaceAttr(&set, m.T)
	e.EncodeElement(m.T, set)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RegularTextRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_RegularTextRun:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "rPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "rPr"}:
				m.RPr = NewCT_TextCharacterProperties()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "t"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "t"}:
				if err := d.DecodeElement(&m.T, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_RegularTextRun %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RegularTextRun
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RegularTextRun and its children
func (m *CT_RegularTextRun) Validate() error {
	return m.ValidateWithPath("CT_RegularTextRun")
}

// ValidateWithPath validates the CT_RegularTextRun and its children, prefixing error messages with path
func (m *CT_RegularTextRun) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	return nil
}
