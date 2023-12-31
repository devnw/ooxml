package sml

import (
	"encoding/xml"
	"fmt"
)

type CT_OleSize struct {
	// Reference
	RefAttr string
}

func NewCT_OleSize() *CT_OleSize {
	ret := &CT_OleSize{}
	return ret
}

func (m *CT_OleSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
		Value: fmt.Sprintf("%v", m.RefAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_OleSize: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OleSize and its children
func (m *CT_OleSize) Validate() error {
	return m.ValidateWithPath("CT_OleSize")
}

// ValidateWithPath validates the CT_OleSize and its children, prefixing error messages with path
func (m *CT_OleSize) ValidateWithPath(path string) error {
	return nil
}
