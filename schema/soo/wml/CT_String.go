package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_String struct {
	// String Value
	ValAttr string
}

func NewCT_String() *CT_String {
	ret := &CT_String{}
	return ret
}

func (m *CT_String) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_String) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_String: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_String and its children
func (m *CT_String) Validate() error {
	return m.ValidateWithPath("CT_String")
}

// ValidateWithPath validates the CT_String and its children, prefixing error messages with path
func (m *CT_String) ValidateWithPath(path string) error {
	return nil
}
