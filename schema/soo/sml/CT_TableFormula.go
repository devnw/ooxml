package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TableFormula struct {
	ArrayAttr bool
	Content   string
}

func NewCT_TableFormula() *CT_TableFormula {
	ret := &CT_TableFormula{}
	return ret
}

func (m *CT_TableFormula) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "array"},
		Value: fmt.Sprintf("%d", b2i(m.ArrayAttr))})
	e.EncodeElement(m.Content, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableFormula) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "array" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ArrayAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TableFormula: %s", err)
		}
		if cd, ok := tok.(xml.CharData); ok {
			m.Content = string(cd)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TableFormula and its children
func (m *CT_TableFormula) Validate() error {
	return m.ValidateWithPath("CT_TableFormula")
}

// ValidateWithPath validates the CT_TableFormula and its children, prefixing error messages with path
func (m *CT_TableFormula) ValidateWithPath(path string) error {
	return nil
}
