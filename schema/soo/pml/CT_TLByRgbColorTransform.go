package pml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_TLByRgbColorTransform struct {
	// Red
	RAttr dml.ST_FixedPercentage
	// Green
	GAttr dml.ST_FixedPercentage
	// Blue
	BAttr dml.ST_FixedPercentage
}

func NewCT_TLByRgbColorTransform() *CT_TLByRgbColorTransform {
	ret := &CT_TLByRgbColorTransform{}
	return ret
}

func (m *CT_TLByRgbColorTransform) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "g"},
		Value: fmt.Sprintf("%v", m.GAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
		Value: fmt.Sprintf("%v", m.BAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLByRgbColorTransform) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.RAttr = parsed
			continue
		}
		if attr.Name.Local == "g" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.GAttr = parsed
			continue
		}
		if attr.Name.Local == "b" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.BAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLByRgbColorTransform: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TLByRgbColorTransform and its children
func (m *CT_TLByRgbColorTransform) Validate() error {
	return m.ValidateWithPath("CT_TLByRgbColorTransform")
}

// ValidateWithPath validates the CT_TLByRgbColorTransform and its children, prefixing error messages with path
func (m *CT_TLByRgbColorTransform) ValidateWithPath(path string) error {
	if err := m.RAttr.ValidateWithPath(path + "/RAttr"); err != nil {
		return err
	}
	if err := m.GAttr.ValidateWithPath(path + "/GAttr"); err != nil {
		return err
	}
	if err := m.BAttr.ValidateWithPath(path + "/BAttr"); err != nil {
		return err
	}
	return nil
}
