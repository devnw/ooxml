package vml

import (
	"encoding/xml"
	"fmt"
)

type Shadow struct {
	CT_Shadow
}

func NewShadow() *Shadow {
	ret := &Shadow{}
	ret.CT_Shadow = *NewCT_Shadow()
	return ret
}

func (m *Shadow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Shadow.MarshalXML(e, start)
}

func (m *Shadow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Shadow = *NewCT_Shadow()
	for _, attr := range start.Attr {
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "obscured" {
			m.ObscuredAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "color" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColorAttr = &parsed
			continue
		}
		if attr.Name.Local == "opacity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OpacityAttr = &parsed
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
		if attr.Name.Local == "color2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Color2Attr = &parsed
			continue
		}
		if attr.Name.Local == "offset2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Offset2Attr = &parsed
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
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Shadow: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Shadow and its children
func (m *Shadow) Validate() error {
	return m.ValidateWithPath("Shadow")
}

// ValidateWithPath validates the Shadow and its children, prefixing error messages with path
func (m *Shadow) ValidateWithPath(path string) error {
	if err := m.CT_Shadow.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
