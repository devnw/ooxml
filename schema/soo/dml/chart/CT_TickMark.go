package chart

import (
	"encoding/xml"
	"fmt"
)

type CT_TickMark struct {
	ValAttr ST_TickMark
}

func NewCT_TickMark() *CT_TickMark {
	ret := &CT_TickMark{}
	return ret
}

func (m *CT_TickMark) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_TickMarkUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TickMark) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_TickMark: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TickMark and its children
func (m *CT_TickMark) Validate() error {
	return m.ValidateWithPath("CT_TickMark")
}

// ValidateWithPath validates the CT_TickMark and its children, prefixing error messages with path
func (m *CT_TickMark) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
