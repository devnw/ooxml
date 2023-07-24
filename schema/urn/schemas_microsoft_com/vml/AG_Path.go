package vml

import (
	"encoding/xml"
	"fmt"
)

type AG_Path struct {
	PathAttr *string
}

func NewAG_Path() *AG_Path {
	ret := &AG_Path{}
	return ret
}

func (m *AG_Path) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PathAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "path"},
			Value: fmt.Sprintf("%v", *m.PathAttr)})
	}
	return nil
}

func (m *AG_Path) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "path" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PathAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Path: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Path and its children
func (m *AG_Path) Validate() error {
	return m.ValidateWithPath("AG_Path")
}

// ValidateWithPath validates the AG_Path and its children, prefixing error messages with path
func (m *AG_Path) ValidateWithPath(path string) error {
	return nil
}
