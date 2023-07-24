package math

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml/schema/soo/ofc/sharedTypes"
)

type CT_YAlign struct {
	ValAttr sharedTypes.ST_YAlign
}

func NewCT_YAlign() *CT_YAlign {
	ret := &CT_YAlign{}
	ret.ValAttr = sharedTypes.ST_YAlign(1)
	return ret
}

func (m *CT_YAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "m:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_YAlign) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = sharedTypes.ST_YAlign(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_YAlign: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_YAlign and its children
func (m *CT_YAlign) Validate() error {
	return m.ValidateWithPath("CT_YAlign")
}

// ValidateWithPath validates the CT_YAlign and its children, prefixing error messages with path
func (m *CT_YAlign) ValidateWithPath(path string) error {
	if m.ValAttr == sharedTypes.ST_YAlignUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
