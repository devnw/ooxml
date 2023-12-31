package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_SectType struct {
	// Section Type Setting
	ValAttr ST_SectionMark
}

func NewCT_SectType() *CT_SectType {
	ret := &CT_SectType{}
	return ret
}

func (m *CT_SectType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_SectionMarkUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SectType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SectType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SectType and its children
func (m *CT_SectType) Validate() error {
	return m.ValidateWithPath("CT_SectType")
}

// ValidateWithPath validates the CT_SectType and its children, prefixing error messages with path
func (m *CT_SectType) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
