package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_MailMergeDest struct {
	// Mail Merge Merged Document Type
	ValAttr ST_MailMergeDest
}

func NewCT_MailMergeDest() *CT_MailMergeDest {
	ret := &CT_MailMergeDest{}
	ret.ValAttr = ST_MailMergeDest(1)
	return ret
}

func (m *CT_MailMergeDest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MailMergeDest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_MailMergeDest(1)
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
			return fmt.Errorf("parsing CT_MailMergeDest: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_MailMergeDest and its children
func (m *CT_MailMergeDest) Validate() error {
	return m.ValidateWithPath("CT_MailMergeDest")
}

// ValidateWithPath validates the CT_MailMergeDest and its children, prefixing error messages with path
func (m *CT_MailMergeDest) ValidateWithPath(path string) error {
	if m.ValAttr == ST_MailMergeDestUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
