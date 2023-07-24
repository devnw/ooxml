package dml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/ofc/sharedTypes"
)

type CT_TextField struct {
	IdAttr   string
	TypeAttr *string
	RPr      *CT_TextCharacterProperties
	PPr      *CT_TextParagraphProperties
	T        *string
}

func NewCT_TextField() *CT_TextField {
	ret := &CT_TextField{}
	ret.IdAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}

func (m *CT_TextField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.TypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"},
			Value: fmt.Sprintf("%v", *m.TypeAttr)})
	}
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "a:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	if m.PPr != nil {
		sepPr := xml.StartElement{Name: xml.Name{Local: "a:pPr"}}
		e.EncodeElement(m.PPr, sepPr)
	}
	if m.T != nil {
		set := xml.StartElement{Name: xml.Name{Local: "a:t"}}
		office.AddPreserveSpaceAttr(&set, *m.T)
		e.EncodeElement(m.T, set)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.IdAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
		if attr.Name.Local == "type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypeAttr = &parsed
			continue
		}
	}
lCT_TextField:
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
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "pPr"}:
				m.PPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.PPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "t"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "t"}:
				m.T = new(string)
				if err := d.DecodeElement(m.T, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_TextField %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextField
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextField and its children
func (m *CT_TextField) Validate() error {
	return m.ValidateWithPath("CT_TextField")
}

// ValidateWithPath validates the CT_TextField and its children, prefixing error messages with path
func (m *CT_TextField) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.IdAttr) {
		return fmt.Errorf(`%s/m.IdAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.IdAttr)
	}
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	if m.PPr != nil {
		if err := m.PPr.ValidateWithPath(path + "/PPr"); err != nil {
			return err
		}
	}
	return nil
}
