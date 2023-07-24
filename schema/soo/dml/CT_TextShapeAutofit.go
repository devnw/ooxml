package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_TextShapeAutofit struct {
}

func NewCT_TextShapeAutofit() *CT_TextShapeAutofit {
	ret := &CT_TextShapeAutofit{}
	return ret
}

func (m *CT_TextShapeAutofit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextShapeAutofit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextShapeAutofit: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextShapeAutofit and its children
func (m *CT_TextShapeAutofit) Validate() error {
	return m.ValidateWithPath("CT_TextShapeAutofit")
}

// ValidateWithPath validates the CT_TextShapeAutofit and its children, prefixing error messages with path
func (m *CT_TextShapeAutofit) ValidateWithPath(path string) error {
	return nil
}
