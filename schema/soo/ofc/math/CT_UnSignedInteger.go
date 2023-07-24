package math

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_UnSignedInteger struct {
	ValAttr uint32
}

func NewCT_UnSignedInteger() *CT_UnSignedInteger {
	ret := &CT_UnSignedInteger{}
	return ret
}

func (m *CT_UnSignedInteger) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "m:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_UnSignedInteger) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ValAttr = uint32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_UnSignedInteger: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_UnSignedInteger and its children
func (m *CT_UnSignedInteger) Validate() error {
	return m.ValidateWithPath("CT_UnSignedInteger")
}

// ValidateWithPath validates the CT_UnSignedInteger and its children, prefixing error messages with path
func (m *CT_UnSignedInteger) ValidateWithPath(path string) error {
	return nil
}
