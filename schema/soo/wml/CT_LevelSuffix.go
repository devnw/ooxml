package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_LevelSuffix struct {
	// Character Type Between Numbering and Text
	ValAttr ST_LevelSuffix
}

func NewCT_LevelSuffix() *CT_LevelSuffix {
	ret := &CT_LevelSuffix{}
	ret.ValAttr = ST_LevelSuffix(1)
	return ret
}

func (m *CT_LevelSuffix) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LevelSuffix) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_LevelSuffix(1)
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
			return fmt.Errorf("parsing CT_LevelSuffix: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LevelSuffix and its children
func (m *CT_LevelSuffix) Validate() error {
	return m.ValidateWithPath("CT_LevelSuffix")
}

// ValidateWithPath validates the CT_LevelSuffix and its children, prefixing error messages with path
func (m *CT_LevelSuffix) ValidateWithPath(path string) error {
	if m.ValAttr == ST_LevelSuffixUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
