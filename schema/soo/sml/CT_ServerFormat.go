package sml

import (
	"encoding/xml"
	"fmt"
)

type CT_ServerFormat struct {
	// Culture
	CultureAttr *string
	// Format
	FormatAttr *string
}

func NewCT_ServerFormat() *CT_ServerFormat {
	ret := &CT_ServerFormat{}
	return ret
}

func (m *CT_ServerFormat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CultureAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "culture"},
			Value: fmt.Sprintf("%v", *m.CultureAttr)})
	}
	if m.FormatAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "format"},
			Value: fmt.Sprintf("%v", *m.FormatAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ServerFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "culture" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CultureAttr = &parsed
			continue
		}
		if attr.Name.Local == "format" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FormatAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ServerFormat: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ServerFormat and its children
func (m *CT_ServerFormat) Validate() error {
	return m.ValidateWithPath("CT_ServerFormat")
}

// ValidateWithPath validates the CT_ServerFormat and its children, prefixing error messages with path
func (m *CT_ServerFormat) ValidateWithPath(path string) error {
	return nil
}
