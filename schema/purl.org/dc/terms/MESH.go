package terms

import (
	"encoding/xml"
	"fmt"
)

type MESH struct {
}

func NewMESH() *MESH {
	ret := &MESH{}
	return ret
}

func (m *MESH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "MESH"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *MESH) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing MESH: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the MESH and its children
func (m *MESH) Validate() error {
	return m.ValidateWithPath("MESH")
}

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (m *MESH) ValidateWithPath(path string) error {
	return nil
}
