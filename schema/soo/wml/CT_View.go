package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_View struct {
	// Document View Setting Value
	ValAttr ST_View
}

func NewCT_View() *CT_View {
	ret := &CT_View{}
	ret.ValAttr = ST_View(1)
	return ret
}

func (m *CT_View) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_View) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_View(1)
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
			return fmt.Errorf("parsing CT_View: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_View and its children
func (m *CT_View) Validate() error {
	return m.ValidateWithPath("CT_View")
}

// ValidateWithPath validates the CT_View and its children, prefixing error messages with path
func (m *CT_View) ValidateWithPath(path string) error {
	if m.ValAttr == ST_ViewUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
