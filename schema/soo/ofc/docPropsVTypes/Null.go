package docPropsVTypes

import (
	"encoding/xml"
	"fmt"
)

type Null struct {
	CT_Null
}

func NewNull() *Null {
	ret := &Null{}
	ret.CT_Null = *NewCT_Null()
	return ret
}

func (m *Null) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Null.MarshalXML(e, start)
}

func (m *Null) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Null = *NewCT_Null()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Null: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Null and its children
func (m *Null) Validate() error {
	return m.ValidateWithPath("Null")
}

// ValidateWithPath validates the Null and its children, prefixing error messages with path
func (m *Null) ValidateWithPath(path string) error {
	if err := m.CT_Null.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
