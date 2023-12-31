package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_TextDirection struct {
	// Direction of Text Flow
	ValAttr ST_TextDirection
}

func NewCT_TextDirection() *CT_TextDirection {
	ret := &CT_TextDirection{}
	ret.ValAttr = ST_TextDirection(1)
	return ret
}

func (m *CT_TextDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_TextDirection(1)
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
			return fmt.Errorf("parsing CT_TextDirection: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextDirection and its children
func (m *CT_TextDirection) Validate() error {
	return m.ValidateWithPath("CT_TextDirection")
}

// ValidateWithPath validates the CT_TextDirection and its children, prefixing error messages with path
func (m *CT_TextDirection) ValidateWithPath(path string) error {
	if m.ValAttr == ST_TextDirectionUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
