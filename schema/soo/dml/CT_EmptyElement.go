package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_EmptyElement struct {
}

func NewCT_EmptyElement() *CT_EmptyElement {
	ret := &CT_EmptyElement{}
	return ret
}

func (m *CT_EmptyElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EmptyElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_EmptyElement: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_EmptyElement and its children
func (m *CT_EmptyElement) Validate() error {
	return m.ValidateWithPath("CT_EmptyElement")
}

// ValidateWithPath validates the CT_EmptyElement and its children, prefixing error messages with path
func (m *CT_EmptyElement) ValidateWithPath(path string) error {
	return nil
}
