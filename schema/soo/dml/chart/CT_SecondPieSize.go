package chart

import (
	"encoding/xml"
	"fmt"
)

type CT_SecondPieSize struct {
	ValAttr *ST_SecondPieSize
}

func NewCT_SecondPieSize() *CT_SecondPieSize {
	ret := &CT_SecondPieSize{}
	return ret
}

func (m *CT_SecondPieSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SecondPieSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_SecondPieSize(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SecondPieSize: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SecondPieSize and its children
func (m *CT_SecondPieSize) Validate() error {
	return m.ValidateWithPath("CT_SecondPieSize")
}

// ValidateWithPath validates the CT_SecondPieSize and its children, prefixing error messages with path
func (m *CT_SecondPieSize) ValidateWithPath(path string) error {
	if m.ValAttr != nil {
		if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
			return err
		}
	}
	return nil
}
