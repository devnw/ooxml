package math

import (
	"encoding/xml"
	"fmt"
)

type CT_TopBot struct {
	ValAttr ST_TopBot
}

func NewCT_TopBot() *CT_TopBot {
	ret := &CT_TopBot{}
	ret.ValAttr = ST_TopBot(1)
	return ret
}

func (m *CT_TopBot) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "m:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TopBot) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_TopBot(1)
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
			return fmt.Errorf("parsing CT_TopBot: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TopBot and its children
func (m *CT_TopBot) Validate() error {
	return m.ValidateWithPath("CT_TopBot")
}

// ValidateWithPath validates the CT_TopBot and its children, prefixing error messages with path
func (m *CT_TopBot) ValidateWithPath(path string) error {
	if m.ValAttr == ST_TopBotUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
