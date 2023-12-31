package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_GroupFillProperties struct {
}

func NewCT_GroupFillProperties() *CT_GroupFillProperties {
	ret := &CT_GroupFillProperties{}
	return ret
}

func (m *CT_GroupFillProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GroupFillProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_GroupFillProperties: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_GroupFillProperties and its children
func (m *CT_GroupFillProperties) Validate() error {
	return m.ValidateWithPath("CT_GroupFillProperties")
}

// ValidateWithPath validates the CT_GroupFillProperties and its children, prefixing error messages with path
func (m *CT_GroupFillProperties) ValidateWithPath(path string) error {
	return nil
}
