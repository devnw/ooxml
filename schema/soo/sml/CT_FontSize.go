package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_FontSize struct {
	// Value
	ValAttr float64
}

func NewCT_FontSize() *CT_FontSize {
	ret := &CT_FontSize{}
	return ret
}

func (m *CT_FontSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FontSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FontSize: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_FontSize and its children
func (m *CT_FontSize) Validate() error {
	return m.ValidateWithPath("CT_FontSize")
}

// ValidateWithPath validates the CT_FontSize and its children, prefixing error messages with path
func (m *CT_FontSize) ValidateWithPath(path string) error {
	return nil
}
