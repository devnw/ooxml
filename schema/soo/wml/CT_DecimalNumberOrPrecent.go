package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_DecimalNumberOrPrecent struct {
	// Value in Percent
	ValAttr ST_DecimalNumberOrPercent
}

func NewCT_DecimalNumberOrPrecent() *CT_DecimalNumberOrPrecent {
	ret := &CT_DecimalNumberOrPrecent{}
	return ret
}

func (m *CT_DecimalNumberOrPrecent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DecimalNumberOrPrecent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_DecimalNumberOrPercent(attr.Value)
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
			return fmt.Errorf("parsing CT_DecimalNumberOrPrecent: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DecimalNumberOrPrecent and its children
func (m *CT_DecimalNumberOrPrecent) Validate() error {
	return m.ValidateWithPath("CT_DecimalNumberOrPrecent")
}

// ValidateWithPath validates the CT_DecimalNumberOrPrecent and its children, prefixing error messages with path
func (m *CT_DecimalNumberOrPrecent) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
