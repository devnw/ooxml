package diagram

import (
	"encoding/xml"
	"fmt"
)

type CT_Direction struct {
	ValAttr ST_Direction
}

func NewCT_Direction() *CT_Direction {
	ret := &CT_Direction{}
	return ret
}

func (m *CT_Direction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_DirectionUnset {
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

func (m *CT_Direction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_Direction: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Direction and its children
func (m *CT_Direction) Validate() error {
	return m.ValidateWithPath("CT_Direction")
}

// ValidateWithPath validates the CT_Direction and its children, prefixing error messages with path
func (m *CT_Direction) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
