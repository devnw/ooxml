package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_TblLayoutType struct {
	// Table Layout Setting
	TypeAttr ST_TblLayoutType
}

func NewCT_TblLayoutType() *CT_TblLayoutType {
	ret := &CT_TblLayoutType{}
	return ret
}

func (m *CT_TblLayoutType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_TblLayoutTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "w:type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblLayoutType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TblLayoutType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TblLayoutType and its children
func (m *CT_TblLayoutType) Validate() error {
	return m.ValidateWithPath("CT_TblLayoutType")
}

// ValidateWithPath validates the CT_TblLayoutType and its children, prefixing error messages with path
func (m *CT_TblLayoutType) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
