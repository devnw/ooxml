package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_MailMergeDocType struct {
	// Mail Merge Source Document Type
	ValAttr ST_MailMergeDocType
}

func NewCT_MailMergeDocType() *CT_MailMergeDocType {
	ret := &CT_MailMergeDocType{}
	ret.ValAttr = ST_MailMergeDocType(1)
	return ret
}

func (m *CT_MailMergeDocType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MailMergeDocType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_MailMergeDocType(1)
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
			return fmt.Errorf("parsing CT_MailMergeDocType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_MailMergeDocType and its children
func (m *CT_MailMergeDocType) Validate() error {
	return m.ValidateWithPath("CT_MailMergeDocType")
}

// ValidateWithPath validates the CT_MailMergeDocType and its children, prefixing error messages with path
func (m *CT_MailMergeDocType) ValidateWithPath(path string) error {
	if m.ValAttr == ST_MailMergeDocTypeUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
