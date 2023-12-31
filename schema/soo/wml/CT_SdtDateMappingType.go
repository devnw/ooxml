package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_SdtDateMappingType struct {
	// Date Storage Type
	ValAttr ST_SdtDateMappingType
}

func NewCT_SdtDateMappingType() *CT_SdtDateMappingType {
	ret := &CT_SdtDateMappingType{}
	return ret
}

func (m *CT_SdtDateMappingType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_SdtDateMappingTypeUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SdtDateMappingType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_SdtDateMappingType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SdtDateMappingType and its children
func (m *CT_SdtDateMappingType) Validate() error {
	return m.ValidateWithPath("CT_SdtDateMappingType")
}

// ValidateWithPath validates the CT_SdtDateMappingType and its children, prefixing error messages with path
func (m *CT_SdtDateMappingType) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
