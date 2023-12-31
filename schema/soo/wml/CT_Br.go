package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_Br struct {
	// Break Type
	TypeAttr ST_BrType
	// Restart Location For Text Wrapping Break
	ClearAttr ST_BrClear
}

func NewCT_Br() *CT_Br {
	ret := &CT_Br{}
	return ret
}

func (m *CT_Br) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_BrTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "w:type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ClearAttr != ST_BrClearUnset {
		attr, err := m.ClearAttr.MarshalXMLAttr(xml.Name{Local: "w:clear"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Br) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "clear" {
			m.ClearAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Br: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Br and its children
func (m *CT_Br) Validate() error {
	return m.ValidateWithPath("CT_Br")
}

// ValidateWithPath validates the CT_Br and its children, prefixing error messages with path
func (m *CT_Br) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.ClearAttr.ValidateWithPath(path + "/ClearAttr"); err != nil {
		return err
	}
	return nil
}
