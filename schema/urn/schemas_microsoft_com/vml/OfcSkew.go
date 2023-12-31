package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcSkew struct {
	OfcCT_Skew
}

func NewOfcSkew() *OfcSkew {
	ret := &OfcSkew{}
	ret.OfcCT_Skew = *NewOfcCT_Skew()
	return ret
}

func (m *OfcSkew) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.OfcCT_Skew.MarshalXML(e, start)
}

func (m *OfcSkew) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_Skew = *NewOfcCT_Skew()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "offset" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OffsetAttr = &parsed
			continue
		}
		if attr.Name.Local == "origin" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OriginAttr = &parsed
			continue
		}
		if attr.Name.Local == "matrix" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MatrixAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcSkew: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcSkew and its children
func (m *OfcSkew) Validate() error {
	return m.ValidateWithPath("OfcSkew")
}

// ValidateWithPath validates the OfcSkew and its children, prefixing error messages with path
func (m *OfcSkew) ValidateWithPath(path string) error {
	if err := m.OfcCT_Skew.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
