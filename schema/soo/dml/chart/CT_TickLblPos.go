package chart

import (
	"encoding/xml"
	"fmt"
)

type CT_TickLblPos struct {
	ValAttr ST_TickLblPos
}

func NewCT_TickLblPos() *CT_TickLblPos {
	ret := &CT_TickLblPos{}
	return ret
}

func (m *CT_TickLblPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_TickLblPosUnset {
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

func (m *CT_TickLblPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_TickLblPos: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TickLblPos and its children
func (m *CT_TickLblPos) Validate() error {
	return m.ValidateWithPath("CT_TickLblPos")
}

// ValidateWithPath validates the CT_TickLblPos and its children, prefixing error messages with path
func (m *CT_TickLblPos) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
