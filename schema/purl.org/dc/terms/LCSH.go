package terms

import (
	"encoding/xml"
	"fmt"
)

type LCSH struct {
}

func NewLCSH() *LCSH {
	ret := &LCSH{}
	return ret
}

func (m *LCSH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "LCSH"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *LCSH) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing LCSH: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the LCSH and its children
func (m *LCSH) Validate() error {
	return m.ValidateWithPath("LCSH")
}

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (m *LCSH) ValidateWithPath(path string) error {
	return nil
}
