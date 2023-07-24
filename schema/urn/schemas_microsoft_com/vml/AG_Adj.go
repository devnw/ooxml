package vml

import (
	"encoding/xml"
	"fmt"
)

type AG_Adj struct {
	AdjAttr *string
}

func NewAG_Adj() *AG_Adj {
	ret := &AG_Adj{}
	return ret
}

func (m *AG_Adj) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AdjAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "adj"},
			Value: fmt.Sprintf("%v", *m.AdjAttr)})
	}
	return nil
}

func (m *AG_Adj) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "adj" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AdjAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Adj: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Adj and its children
func (m *AG_Adj) Validate() error {
	return m.ValidateWithPath("AG_Adj")
}

// ValidateWithPath validates the AG_Adj and its children, prefixing error messages with path
func (m *AG_Adj) ValidateWithPath(path string) error {
	return nil
}
