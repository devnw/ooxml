package math

import (
	"encoding/xml"
	"fmt"
)

type CT_Script struct {
	ValAttr ST_Script
}

func NewCT_Script() *CT_Script {
	ret := &CT_Script{}
	return ret
}

func (m *CT_Script) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_ScriptUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "m:val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Script) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_Script: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Script and its children
func (m *CT_Script) Validate() error {
	return m.ValidateWithPath("CT_Script")
}

// ValidateWithPath validates the CT_Script and its children, prefixing error messages with path
func (m *CT_Script) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
