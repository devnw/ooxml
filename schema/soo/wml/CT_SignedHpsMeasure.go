package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_SignedHpsMeasure struct {
	// Signed Half-Point Measurement
	ValAttr ST_SignedHpsMeasure
}

func NewCT_SignedHpsMeasure() *CT_SignedHpsMeasure {
	ret := &CT_SignedHpsMeasure{}
	return ret
}

func (m *CT_SignedHpsMeasure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SignedHpsMeasure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_SignedHpsMeasure(attr.Value)
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
			return fmt.Errorf("parsing CT_SignedHpsMeasure: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SignedHpsMeasure and its children
func (m *CT_SignedHpsMeasure) Validate() error {
	return m.ValidateWithPath("CT_SignedHpsMeasure")
}

// ValidateWithPath validates the CT_SignedHpsMeasure and its children, prefixing error messages with path
func (m *CT_SignedHpsMeasure) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
