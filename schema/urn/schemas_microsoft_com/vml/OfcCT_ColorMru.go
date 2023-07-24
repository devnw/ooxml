package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcCT_ColorMru struct {
	ColorsAttr *string
	ExtAttr    ST_Ext
}

func NewOfcCT_ColorMru() *OfcCT_ColorMru {
	ret := &OfcCT_ColorMru{}
	return ret
}

func (m *OfcCT_ColorMru) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ColorsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "colors"},
			Value: fmt.Sprintf("%v", *m.ColorsAttr)})
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

func (m *OfcCT_ColorMru) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "colors" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColorsAttr = &parsed
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
			return fmt.Errorf("parsing OfcCT_ColorMru: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_ColorMru and its children
func (m *OfcCT_ColorMru) Validate() error {
	return m.ValidateWithPath("OfcCT_ColorMru")
}

// ValidateWithPath validates the OfcCT_ColorMru and its children, prefixing error messages with path
func (m *OfcCT_ColorMru) ValidateWithPath(path string) error {
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
