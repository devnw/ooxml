package terms

import (
	"encoding/xml"
	"fmt"
)

type UDC struct {
}

func NewUDC() *UDC {
	ret := &UDC{}
	return ret
}

func (m *UDC) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "UDC"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *UDC) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing UDC: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the UDC and its children
func (m *UDC) Validate() error {
	return m.ValidateWithPath("UDC")
}

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (m *UDC) ValidateWithPath(path string) error {
	return nil
}
