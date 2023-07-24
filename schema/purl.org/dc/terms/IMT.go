package terms

import (
	"encoding/xml"
	"fmt"
)

type IMT struct {
}

func NewIMT() *IMT {
	ret := &IMT{}
	return ret
}

func (m *IMT) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "IMT"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *IMT) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing IMT: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the IMT and its children
func (m *IMT) Validate() error {
	return m.ValidateWithPath("IMT")
}

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (m *IMT) ValidateWithPath(path string) error {
	return nil
}
