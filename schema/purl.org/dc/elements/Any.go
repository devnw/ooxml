package elements

import (
	"encoding/xml"
	"fmt"
)

type Any struct {
	SimpleLiteral
}

func NewAny() *Any {
	ret := &Any{}
	ret.SimpleLiteral = *NewSimpleLiteral()
	return ret
}

func (m *Any) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.SimpleLiteral.MarshalXML(e, start)
}

func (m *Any) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SimpleLiteral = *NewSimpleLiteral()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Any: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Any and its children
func (m *Any) Validate() error {
	return m.ValidateWithPath("Any")
}

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (m *Any) ValidateWithPath(path string) error {
	if err := m.SimpleLiteral.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
