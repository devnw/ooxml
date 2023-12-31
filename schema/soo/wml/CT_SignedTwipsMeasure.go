package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_SignedTwipsMeasure struct {
	// Positive or Negative Value in Twentieths of a Point
	ValAttr ST_SignedTwipsMeasure
}

func NewCT_SignedTwipsMeasure() *CT_SignedTwipsMeasure {
	ret := &CT_SignedTwipsMeasure{}
	return ret
}

func (m *CT_SignedTwipsMeasure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SignedTwipsMeasure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
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
			return fmt.Errorf("parsing CT_SignedTwipsMeasure: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SignedTwipsMeasure and its children
func (m *CT_SignedTwipsMeasure) Validate() error {
	return m.ValidateWithPath("CT_SignedTwipsMeasure")
}

// ValidateWithPath validates the CT_SignedTwipsMeasure and its children, prefixing error messages with path
func (m *CT_SignedTwipsMeasure) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
