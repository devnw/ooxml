package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_TextAlignment struct {
	// Vertical Character Alignment Position
	ValAttr ST_TextAlignment
}

func NewCT_TextAlignment() *CT_TextAlignment {
	ret := &CT_TextAlignment{}
	ret.ValAttr = ST_TextAlignment(1)
	return ret
}

func (m *CT_TextAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_TextAlignment(1)
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
			return fmt.Errorf("parsing CT_TextAlignment: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextAlignment and its children
func (m *CT_TextAlignment) Validate() error {
	return m.ValidateWithPath("CT_TextAlignment")
}

// ValidateWithPath validates the CT_TextAlignment and its children, prefixing error messages with path
func (m *CT_TextAlignment) ValidateWithPath(path string) error {
	if m.ValAttr == ST_TextAlignmentUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
