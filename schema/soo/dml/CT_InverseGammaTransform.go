package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_InverseGammaTransform struct {
}

func NewCT_InverseGammaTransform() *CT_InverseGammaTransform {
	ret := &CT_InverseGammaTransform{}
	return ret
}

func (m *CT_InverseGammaTransform) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_InverseGammaTransform) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_InverseGammaTransform: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_InverseGammaTransform and its children
func (m *CT_InverseGammaTransform) Validate() error {
	return m.ValidateWithPath("CT_InverseGammaTransform")
}

// ValidateWithPath validates the CT_InverseGammaTransform and its children, prefixing error messages with path
func (m *CT_InverseGammaTransform) ValidateWithPath(path string) error {
	return nil
}
