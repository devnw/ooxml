package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_NoFillProperties struct {
}

func NewCT_NoFillProperties() *CT_NoFillProperties {
	ret := &CT_NoFillProperties{}
	return ret
}

func (m *CT_NoFillProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NoFillProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_NoFillProperties: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_NoFillProperties and its children
func (m *CT_NoFillProperties) Validate() error {
	return m.ValidateWithPath("CT_NoFillProperties")
}

// ValidateWithPath validates the CT_NoFillProperties and its children, prefixing error messages with path
func (m *CT_NoFillProperties) ValidateWithPath(path string) error {
	return nil
}
