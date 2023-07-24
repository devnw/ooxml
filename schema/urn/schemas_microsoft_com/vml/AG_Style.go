package vml

import (
	"encoding/xml"
	"fmt"
)

type AG_Style struct {
	StyleAttr *string
}

func NewAG_Style() *AG_Style {
	ret := &AG_Style{}
	return ret
}

func (m *AG_Style) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.StyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"},
			Value: fmt.Sprintf("%v", *m.StyleAttr)})
	}
	return nil
}

func (m *AG_Style) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Style: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Style and its children
func (m *AG_Style) Validate() error {
	return m.ValidateWithPath("AG_Style")
}

// ValidateWithPath validates the AG_Style and its children, prefixing error messages with path
func (m *AG_Style) ValidateWithPath(path string) error {
	return nil
}
