package math

import (
	"encoding/xml"
	"fmt"
)

type CT_LimLoc struct {
	ValAttr ST_LimLoc
}

func NewCT_LimLoc() *CT_LimLoc {
	ret := &CT_LimLoc{}
	ret.ValAttr = ST_LimLoc(1)
	return ret
}

func (m *CT_LimLoc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "m:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LimLoc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_LimLoc(1)
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
			return fmt.Errorf("parsing CT_LimLoc: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LimLoc and its children
func (m *CT_LimLoc) Validate() error {
	return m.ValidateWithPath("CT_LimLoc")
}

// ValidateWithPath validates the CT_LimLoc and its children, prefixing error messages with path
func (m *CT_LimLoc) ValidateWithPath(path string) error {
	if m.ValAttr == ST_LimLocUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
