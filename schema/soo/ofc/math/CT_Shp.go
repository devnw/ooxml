package math

import (
	"encoding/xml"
	"fmt"
)

type CT_Shp struct {
	ValAttr ST_Shp
}

func NewCT_Shp() *CT_Shp {
	ret := &CT_Shp{}
	ret.ValAttr = ST_Shp(1)
	return ret
}

func (m *CT_Shp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "m:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Shp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_Shp(1)
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
			return fmt.Errorf("parsing CT_Shp: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Shp and its children
func (m *CT_Shp) Validate() error {
	return m.ValidateWithPath("CT_Shp")
}

// ValidateWithPath validates the CT_Shp and its children, prefixing error messages with path
func (m *CT_Shp) ValidateWithPath(path string) error {
	if m.ValAttr == ST_ShpUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
