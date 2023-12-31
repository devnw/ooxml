package elements

import (
	"encoding/xml"
	"fmt"
)

type SimpleLiteral struct {
}

func NewSimpleLiteral() *SimpleLiteral {
	ret := &SimpleLiteral{}
	return ret
}

func (m *SimpleLiteral) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *SimpleLiteral) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing SimpleLiteral: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the SimpleLiteral and its children
func (m *SimpleLiteral) Validate() error {
	return m.ValidateWithPath("SimpleLiteral")
}

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (m *SimpleLiteral) ValidateWithPath(path string) error {
	return nil
}
