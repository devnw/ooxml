package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Field struct {
	// Field Index
	XAttr int32
}

func NewCT_Field() *CT_Field {
	ret := &CT_Field{}
	return ret
}

func (m *CT_Field) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "x"},
		Value: fmt.Sprintf("%v", m.XAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Field) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "x" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.XAttr = int32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Field: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Field and its children
func (m *CT_Field) Validate() error {
	return m.ValidateWithPath("CT_Field")
}

// ValidateWithPath validates the CT_Field and its children, prefixing error messages with path
func (m *CT_Field) ValidateWithPath(path string) error {
	return nil
}
