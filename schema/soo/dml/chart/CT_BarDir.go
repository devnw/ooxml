package chart

import (
	"encoding/xml"
	"fmt"
)

type CT_BarDir struct {
	ValAttr ST_BarDir
}

func NewCT_BarDir() *CT_BarDir {
	ret := &CT_BarDir{}
	return ret
}

func (m *CT_BarDir) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_BarDirUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BarDir) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_BarDir: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_BarDir and its children
func (m *CT_BarDir) Validate() error {
	return m.ValidateWithPath("CT_BarDir")
}

// ValidateWithPath validates the CT_BarDir and its children, prefixing error messages with path
func (m *CT_BarDir) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
