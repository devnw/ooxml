package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_InverseTransform struct {
}

func NewCT_InverseTransform() *CT_InverseTransform {
	ret := &CT_InverseTransform{}
	return ret
}

func (m *CT_InverseTransform) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_InverseTransform) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_InverseTransform: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_InverseTransform and its children
func (m *CT_InverseTransform) Validate() error {
	return m.ValidateWithPath("CT_InverseTransform")
}

// ValidateWithPath validates the CT_InverseTransform and its children, prefixing error messages with path
func (m *CT_InverseTransform) ValidateWithPath(path string) error {
	return nil
}
