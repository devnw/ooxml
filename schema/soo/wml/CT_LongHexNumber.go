package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_LongHexNumber struct {
	// Long Hexadecimal Number Value
	ValAttr string
}

func NewCT_LongHexNumber() *CT_LongHexNumber {
	ret := &CT_LongHexNumber{}
	return ret
}

func (m *CT_LongHexNumber) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LongHexNumber) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_LongHexNumber: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LongHexNumber and its children
func (m *CT_LongHexNumber) Validate() error {
	return m.ValidateWithPath("CT_LongHexNumber")
}

// ValidateWithPath validates the CT_LongHexNumber and its children, prefixing error messages with path
func (m *CT_LongHexNumber) ValidateWithPath(path string) error {
	return nil
}
