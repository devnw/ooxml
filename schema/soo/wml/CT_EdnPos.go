package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_EdnPos struct {
	// Endnote Position Type
	ValAttr ST_EdnPos
}

func NewCT_EdnPos() *CT_EdnPos {
	ret := &CT_EdnPos{}
	ret.ValAttr = ST_EdnPos(1)
	return ret
}

func (m *CT_EdnPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EdnPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_EdnPos(1)
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
			return fmt.Errorf("parsing CT_EdnPos: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_EdnPos and its children
func (m *CT_EdnPos) Validate() error {
	return m.ValidateWithPath("CT_EdnPos")
}

// ValidateWithPath validates the CT_EdnPos and its children, prefixing error messages with path
func (m *CT_EdnPos) ValidateWithPath(path string) error {
	if m.ValAttr == ST_EdnPosUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
