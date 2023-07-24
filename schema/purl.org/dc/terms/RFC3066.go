package terms

import (
	"encoding/xml"
	"fmt"
)

type RFC3066 struct {
}

func NewRFC3066() *RFC3066 {
	ret := &RFC3066{}
	return ret
}

func (m *RFC3066) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "RFC3066"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *RFC3066) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing RFC3066: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the RFC3066 and its children
func (m *RFC3066) Validate() error {
	return m.ValidateWithPath("RFC3066")
}

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (m *RFC3066) ValidateWithPath(path string) error {
	return nil
}
