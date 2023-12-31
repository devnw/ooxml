package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_Jc struct {
	// Alignment Type
	ValAttr ST_Jc
}

func NewCT_Jc() *CT_Jc {
	ret := &CT_Jc{}
	ret.ValAttr = ST_Jc(1)
	return ret
}

func (m *CT_Jc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Jc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_Jc(1)
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
			return fmt.Errorf("parsing CT_Jc: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Jc and its children
func (m *CT_Jc) Validate() error {
	return m.ValidateWithPath("CT_Jc")
}

// ValidateWithPath validates the CT_Jc and its children, prefixing error messages with path
func (m *CT_Jc) ValidateWithPath(path string) error {
	if m.ValAttr == ST_JcUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
