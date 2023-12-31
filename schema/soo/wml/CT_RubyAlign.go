package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_RubyAlign struct {
	// Phonetic Guide Text Alignment Value
	ValAttr ST_RubyAlign
}

func NewCT_RubyAlign() *CT_RubyAlign {
	ret := &CT_RubyAlign{}
	ret.ValAttr = ST_RubyAlign(1)
	return ret
}

func (m *CT_RubyAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RubyAlign) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_RubyAlign(1)
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
			return fmt.Errorf("parsing CT_RubyAlign: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_RubyAlign and its children
func (m *CT_RubyAlign) Validate() error {
	return m.ValidateWithPath("CT_RubyAlign")
}

// ValidateWithPath validates the CT_RubyAlign and its children, prefixing error messages with path
func (m *CT_RubyAlign) ValidateWithPath(path string) error {
	if m.ValAttr == ST_RubyAlignUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
