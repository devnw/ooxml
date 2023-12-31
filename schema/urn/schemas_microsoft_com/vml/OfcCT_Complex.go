package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcCT_Complex struct {
	ExtAttr ST_Ext
}

func NewOfcCT_Complex() *OfcCT_Complex {
	ret := &OfcCT_Complex{}
	return ret
}

func (m *OfcCT_Complex) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_Complex) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_Complex: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_Complex and its children
func (m *OfcCT_Complex) Validate() error {
	return m.ValidateWithPath("OfcCT_Complex")
}

// ValidateWithPath validates the OfcCT_Complex and its children, prefixing error messages with path
func (m *OfcCT_Complex) ValidateWithPath(path string) error {
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
