package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_TextEffect struct {
	// Animated Text Effect Type
	ValAttr ST_TextEffect
}

func NewCT_TextEffect() *CT_TextEffect {
	ret := &CT_TextEffect{}
	ret.ValAttr = ST_TextEffect(1)
	return ret
}

func (m *CT_TextEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_TextEffect(1)
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
			return fmt.Errorf("parsing CT_TextEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextEffect and its children
func (m *CT_TextEffect) Validate() error {
	return m.ValidateWithPath("CT_TextEffect")
}

// ValidateWithPath validates the CT_TextEffect and its children, prefixing error messages with path
func (m *CT_TextEffect) ValidateWithPath(path string) error {
	if m.ValAttr == ST_TextEffectUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
