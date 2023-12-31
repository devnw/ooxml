package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcCallout struct {
	OfcCT_Callout
}

func NewOfcCallout() *OfcCallout {
	ret := &OfcCallout{}
	ret.OfcCT_Callout = *NewOfcCT_Callout()
	return ret
}

func (m *OfcCallout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.OfcCT_Callout.MarshalXML(e, start)
}

func (m *OfcCallout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_Callout = *NewOfcCT_Callout()
	for _, attr := range start.Attr {
		if attr.Name.Local == "lengthspecified" {
			m.LengthspecifiedAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "gap" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GapAttr = &parsed
			continue
		}
		if attr.Name.Local == "angle" {
			m.AngleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dropauto" {
			m.DropautoAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "drop" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DropAttr = &parsed
			continue
		}
		if attr.Name.Local == "distance" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DistanceAttr = &parsed
			continue
		}
		if attr.Name.Local == "type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypeAttr = &parsed
			continue
		}
		if attr.Name.Local == "length" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LengthAttr = &parsed
			continue
		}
		if attr.Name.Local == "accentbar" {
			m.AccentbarAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "textborder" {
			m.TextborderAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "minusx" {
			m.MinusxAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "minusy" {
			m.MinusyAttr.UnmarshalXMLAttr(attr)
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
			return fmt.Errorf("parsing OfcCallout: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCallout and its children
func (m *OfcCallout) Validate() error {
	return m.ValidateWithPath("OfcCallout")
}

// ValidateWithPath validates the OfcCallout and its children, prefixing error messages with path
func (m *OfcCallout) ValidateWithPath(path string) error {
	if err := m.OfcCT_Callout.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
