package vml

import (
	"encoding/xml"
	"fmt"
)

type AG_Id struct {
	IdAttr *string
}

func NewAG_Id() *AG_Id {
	ret := &AG_Id{}
	return ret
}

func (m *AG_Id) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	return nil
}

func (m *AG_Id) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Id: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Id and its children
func (m *AG_Id) Validate() error {
	return m.ValidateWithPath("AG_Id")
}

// ValidateWithPath validates the AG_Id and its children, prefixing error messages with path
func (m *AG_Id) ValidateWithPath(path string) error {
	return nil
}
