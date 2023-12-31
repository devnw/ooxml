package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcCT_Fill struct {
	TypeAttr OfcST_FillType
	ExtAttr  ST_Ext
}

func NewOfcCT_Fill() *OfcCT_Fill {
	ret := &OfcCT_Fill{}
	return ret
}

func (m *OfcCT_Fill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != OfcST_FillTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
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

func (m *OfcCT_Fill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_Fill: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_Fill and its children
func (m *OfcCT_Fill) Validate() error {
	return m.ValidateWithPath("OfcCT_Fill")
}

// ValidateWithPath validates the OfcCT_Fill and its children, prefixing error messages with path
func (m *OfcCT_Fill) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
