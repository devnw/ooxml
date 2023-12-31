package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_DocPartBehavior struct {
	// Insertion Behavior Value
	ValAttr ST_DocPartBehavior
}

func NewCT_DocPartBehavior() *CT_DocPartBehavior {
	ret := &CT_DocPartBehavior{}
	ret.ValAttr = ST_DocPartBehavior(1)
	return ret
}

func (m *CT_DocPartBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocPartBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_DocPartBehavior(1)
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
			return fmt.Errorf("parsing CT_DocPartBehavior: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DocPartBehavior and its children
func (m *CT_DocPartBehavior) Validate() error {
	return m.ValidateWithPath("CT_DocPartBehavior")
}

// ValidateWithPath validates the CT_DocPartBehavior and its children, prefixing error messages with path
func (m *CT_DocPartBehavior) ValidateWithPath(path string) error {
	if m.ValAttr == ST_DocPartBehaviorUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
