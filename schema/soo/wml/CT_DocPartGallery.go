package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_DocPartGallery struct {
	// Gallery Value
	ValAttr ST_DocPartGallery
}

func NewCT_DocPartGallery() *CT_DocPartGallery {
	ret := &CT_DocPartGallery{}
	ret.ValAttr = ST_DocPartGallery(1)
	return ret
}

func (m *CT_DocPartGallery) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocPartGallery) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_DocPartGallery(1)
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
			return fmt.Errorf("parsing CT_DocPartGallery: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DocPartGallery and its children
func (m *CT_DocPartGallery) Validate() error {
	return m.ValidateWithPath("CT_DocPartGallery")
}

// ValidateWithPath validates the CT_DocPartGallery and its children, prefixing error messages with path
func (m *CT_DocPartGallery) ValidateWithPath(path string) error {
	if m.ValAttr == ST_DocPartGalleryUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
