package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_Sym struct {
	// Symbol Character Font
	FontAttr *string
	// Symbol Character Code
	CharAttr *string
}

func NewCT_Sym() *CT_Sym {
	ret := &CT_Sym{}
	return ret
}

func (m *CT_Sym) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FontAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:font"},
			Value: fmt.Sprintf("%v", *m.FontAttr)})
	}
	if m.CharAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:char"},
			Value: fmt.Sprintf("%v", *m.CharAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Sym) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "font" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FontAttr = &parsed
			continue
		}
		if attr.Name.Local == "char" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CharAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Sym: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Sym and its children
func (m *CT_Sym) Validate() error {
	return m.ValidateWithPath("CT_Sym")
}

// ValidateWithPath validates the CT_Sym and its children, prefixing error messages with path
func (m *CT_Sym) ValidateWithPath(path string) error {
	return nil
}
