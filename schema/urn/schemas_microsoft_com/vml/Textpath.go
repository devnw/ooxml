package vml

import (
	"encoding/xml"
	"fmt"
)

type Textpath struct {
	CT_TextPath
}

func NewTextpath() *Textpath {
	ret := &Textpath{}
	ret.CT_TextPath = *NewCT_TextPath()
	return ret
}

func (m *Textpath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_TextPath.MarshalXML(e, start)
}

func (m *Textpath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_TextPath = *NewCT_TextPath()
	for _, attr := range start.Attr {
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "fitshape" {
			m.FitshapeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "fitpath" {
			m.FitpathAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "trim" {
			m.TrimAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "xscale" {
			m.XscaleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "string" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StringAttr = &parsed
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
		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Textpath: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Textpath and its children
func (m *Textpath) Validate() error {
	return m.ValidateWithPath("Textpath")
}

// ValidateWithPath validates the Textpath and its children, prefixing error messages with path
func (m *Textpath) ValidateWithPath(path string) error {
	if err := m.CT_TextPath.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
