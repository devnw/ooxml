package chart

import (
	"encoding/xml"
	"fmt"
)

type CT_ErrDir struct {
	ValAttr ST_ErrDir
}

func NewCT_ErrDir() *CT_ErrDir {
	ret := &CT_ErrDir{}
	ret.ValAttr = ST_ErrDir(1)
	return ret
}

func (m *CT_ErrDir) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ErrDir) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_ErrDir(1)
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
			return fmt.Errorf("parsing CT_ErrDir: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ErrDir and its children
func (m *CT_ErrDir) Validate() error {
	return m.ValidateWithPath("CT_ErrDir")
}

// ValidateWithPath validates the CT_ErrDir and its children, prefixing error messages with path
func (m *CT_ErrDir) ValidateWithPath(path string) error {
	if m.ValAttr == ST_ErrDirUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
