package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_LineEndProperties struct {
	TypeAttr ST_LineEndType
	WAttr    ST_LineEndWidth
	LenAttr  ST_LineEndLength
}

func NewCT_LineEndProperties() *CT_LineEndProperties {
	ret := &CT_LineEndProperties{}
	return ret
}

func (m *CT_LineEndProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_LineEndTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.WAttr != ST_LineEndWidthUnset {
		attr, err := m.WAttr.MarshalXMLAttr(xml.Name{Local: "w"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.LenAttr != ST_LineEndLengthUnset {
		attr, err := m.LenAttr.MarshalXMLAttr(xml.Name{Local: "len"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LineEndProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "w" {
			m.WAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "len" {
			m.LenAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LineEndProperties: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LineEndProperties and its children
func (m *CT_LineEndProperties) Validate() error {
	return m.ValidateWithPath("CT_LineEndProperties")
}

// ValidateWithPath validates the CT_LineEndProperties and its children, prefixing error messages with path
func (m *CT_LineEndProperties) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.WAttr.ValidateWithPath(path + "/WAttr"); err != nil {
		return err
	}
	if err := m.LenAttr.ValidateWithPath(path + "/LenAttr"); err != nil {
		return err
	}
	return nil
}
