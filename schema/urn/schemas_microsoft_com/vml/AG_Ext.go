package vml

import (
	"encoding/xml"
	"fmt"
)

type AG_Ext struct {
	ExtAttr ST_Ext
}

func NewAG_Ext() *AG_Ext {
	ret := &AG_Ext{}
	ret.ExtAttr = ST_Ext(1)
	return ret
}

func (m *AG_Ext) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	start.Name.Local = "v:AG_Ext"
	return nil
}

func (m *AG_Ext) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ExtAttr = ST_Ext(1)
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
			return fmt.Errorf("parsing AG_Ext: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Ext and its children
func (m *AG_Ext) Validate() error {
	return m.ValidateWithPath("AG_Ext")
}

// ValidateWithPath validates the AG_Ext and its children, prefixing error messages with path
func (m *AG_Ext) ValidateWithPath(path string) error {
	if m.ExtAttr == ST_ExtUnset {
		return fmt.Errorf("%s/ExtAttr is a mandatory field", path)
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
