package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TLAnimVariantIntegerVal struct {
	// Value
	ValAttr int32
}

func NewCT_TLAnimVariantIntegerVal() *CT_TLAnimVariantIntegerVal {
	ret := &CT_TLAnimVariantIntegerVal{}
	return ret
}

func (m *CT_TLAnimVariantIntegerVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLAnimVariantIntegerVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ValAttr = int32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLAnimVariantIntegerVal: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TLAnimVariantIntegerVal and its children
func (m *CT_TLAnimVariantIntegerVal) Validate() error {
	return m.ValidateWithPath("CT_TLAnimVariantIntegerVal")
}

// ValidateWithPath validates the CT_TLAnimVariantIntegerVal and its children, prefixing error messages with path
func (m *CT_TLAnimVariantIntegerVal) ValidateWithPath(path string) error {
	return nil
}
