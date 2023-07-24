package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_GammaTransform struct {
}

func NewCT_GammaTransform() *CT_GammaTransform {
	ret := &CT_GammaTransform{}
	return ret
}

func (m *CT_GammaTransform) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GammaTransform) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_GammaTransform: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_GammaTransform and its children
func (m *CT_GammaTransform) Validate() error {
	return m.ValidateWithPath("CT_GammaTransform")
}

// ValidateWithPath validates the CT_GammaTransform and its children, prefixing error messages with path
func (m *CT_GammaTransform) ValidateWithPath(path string) error {
	return nil
}
