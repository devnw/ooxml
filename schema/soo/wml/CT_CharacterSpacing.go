package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_CharacterSpacing struct {
	// Value
	ValAttr ST_CharacterSpacing
}

func NewCT_CharacterSpacing() *CT_CharacterSpacing {
	ret := &CT_CharacterSpacing{}
	ret.ValAttr = ST_CharacterSpacing(1)
	return ret
}

func (m *CT_CharacterSpacing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CharacterSpacing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_CharacterSpacing(1)
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
			return fmt.Errorf("parsing CT_CharacterSpacing: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CharacterSpacing and its children
func (m *CT_CharacterSpacing) Validate() error {
	return m.ValidateWithPath("CT_CharacterSpacing")
}

// ValidateWithPath validates the CT_CharacterSpacing and its children, prefixing error messages with path
func (m *CT_CharacterSpacing) ValidateWithPath(path string) error {
	if m.ValAttr == ST_CharacterSpacingUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
