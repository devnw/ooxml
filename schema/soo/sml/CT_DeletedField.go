package sml

import (
	"encoding/xml"
	"fmt"
)

type CT_DeletedField struct {
	// Deleted Fields Name
	NameAttr string
}

func NewCT_DeletedField() *CT_DeletedField {
	ret := &CT_DeletedField{}
	return ret
}

func (m *CT_DeletedField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DeletedField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DeletedField: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DeletedField and its children
func (m *CT_DeletedField) Validate() error {
	return m.ValidateWithPath("CT_DeletedField")
}

// ValidateWithPath validates the CT_DeletedField and its children, prefixing error messages with path
func (m *CT_DeletedField) ValidateWithPath(path string) error {
	return nil
}
