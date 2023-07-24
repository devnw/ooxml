package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_Base64Binary struct {
	ValAttr string
}

func NewCT_Base64Binary() *CT_Base64Binary {
	ret := &CT_Base64Binary{}
	return ret
}

func (m *CT_Base64Binary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Base64Binary) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_Base64Binary: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Base64Binary and its children
func (m *CT_Base64Binary) Validate() error {
	return m.ValidateWithPath("CT_Base64Binary")
}

// ValidateWithPath validates the CT_Base64Binary and its children, prefixing error messages with path
func (m *CT_Base64Binary) ValidateWithPath(path string) error {
	return nil
}
