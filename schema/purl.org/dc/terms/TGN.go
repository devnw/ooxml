package terms

import (
	"encoding/xml"
	"fmt"
)

type TGN struct {
}

func NewTGN() *TGN {
	ret := &TGN{}
	return ret
}

func (m *TGN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "TGN"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *TGN) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing TGN: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the TGN and its children
func (m *TGN) Validate() error {
	return m.ValidateWithPath("TGN")
}

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (m *TGN) ValidateWithPath(path string) error {
	return nil
}
