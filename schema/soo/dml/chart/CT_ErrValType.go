package chart

import (
	"encoding/xml"
	"fmt"
)

type CT_ErrValType struct {
	ValAttr ST_ErrValType
}

func NewCT_ErrValType() *CT_ErrValType {
	ret := &CT_ErrValType{}
	return ret
}

func (m *CT_ErrValType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_ErrValTypeUnset {
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

func (m *CT_ErrValType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_ErrValType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ErrValType and its children
func (m *CT_ErrValType) Validate() error {
	return m.ValidateWithPath("CT_ErrValType")
}

// ValidateWithPath validates the CT_ErrValType and its children, prefixing error messages with path
func (m *CT_ErrValType) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
