package math

import (
	"encoding/xml"
	"fmt"
)

type CT_OMathJc struct {
	ValAttr ST_Jc
}

func NewCT_OMathJc() *CT_OMathJc {
	ret := &CT_OMathJc{}
	return ret
}

func (m *CT_OMathJc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_JcUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "m:val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OMathJc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
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
			return fmt.Errorf("parsing CT_OMathJc: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OMathJc and its children
func (m *CT_OMathJc) Validate() error {
	return m.ValidateWithPath("CT_OMathJc")
}

// ValidateWithPath validates the CT_OMathJc and its children, prefixing error messages with path
func (m *CT_OMathJc) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
