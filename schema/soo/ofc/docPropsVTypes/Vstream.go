package docPropsVTypes

import (
	"encoding/xml"
	"fmt"
)

type Vstream struct {
	CT_Vstream
}

func NewVstream() *Vstream {
	ret := &Vstream{}
	ret.CT_Vstream = *NewCT_Vstream()
	return ret
}

func (m *Vstream) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Vstream.MarshalXML(e, start)
}

func (m *Vstream) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Vstream = *NewCT_Vstream()
	for _, attr := range start.Attr {
		if attr.Name.Local == "version" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VersionAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Vstream: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Vstream and its children
func (m *Vstream) Validate() error {
	return m.ValidateWithPath("Vstream")
}

// ValidateWithPath validates the Vstream and its children, prefixing error messages with path
func (m *Vstream) ValidateWithPath(path string) error {
	if err := m.CT_Vstream.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
