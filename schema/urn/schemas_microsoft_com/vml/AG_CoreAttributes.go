package vml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml/schema/soo/ofc/sharedTypes"
)

type AG_CoreAttributes struct {
	HrefAttr        *string
	TargetAttr      *string
	ClassAttr       *string
	TitleAttr       *string
	AltAttr         *string
	CoordsizeAttr   *string
	CoordoriginAttr *string
	WrapcoordsAttr  *string
	PrintAttr       sharedTypes.ST_TrueFalse
	IdAttr          *string
	StyleAttr       *string
}

func NewAG_CoreAttributes() *AG_CoreAttributes {
	ret := &AG_CoreAttributes{}
	return ret
}

func (m *AG_CoreAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.HrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "href"},
			Value: fmt.Sprintf("%v", *m.HrefAttr)})
	}
	if m.TargetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "target"},
			Value: fmt.Sprintf("%v", *m.TargetAttr)})
	}
	if m.ClassAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "class"},
			Value: fmt.Sprintf("%v", *m.ClassAttr)})
	}
	if m.TitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "title"},
			Value: fmt.Sprintf("%v", *m.TitleAttr)})
	}
	if m.AltAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "alt"},
			Value: fmt.Sprintf("%v", *m.AltAttr)})
	}
	if m.CoordsizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coordsize"},
			Value: fmt.Sprintf("%v", *m.CoordsizeAttr)})
	}
	if m.CoordoriginAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coordorigin"},
			Value: fmt.Sprintf("%v", *m.CoordoriginAttr)})
	}
	if m.WrapcoordsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "wrapcoords"},
			Value: fmt.Sprintf("%v", *m.WrapcoordsAttr)})
	}
	if m.PrintAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.PrintAttr.MarshalXMLAttr(xml.Name{Local: "print"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.StyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"},
			Value: fmt.Sprintf("%v", *m.StyleAttr)})
	}
	return nil
}

func (m *AG_CoreAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HrefAttr = &parsed
			continue
		}
		if attr.Name.Local == "target" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TargetAttr = &parsed
			continue
		}
		if attr.Name.Local == "class" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ClassAttr = &parsed
			continue
		}
		if attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
			continue
		}
		if attr.Name.Local == "alt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AltAttr = &parsed
			continue
		}
		if attr.Name.Local == "coordsize" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CoordsizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "coordorigin" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CoordoriginAttr = &parsed
			continue
		}
		if attr.Name.Local == "wrapcoords" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WrapcoordsAttr = &parsed
			continue
		}
		if attr.Name.Local == "print" {
			m.PrintAttr.UnmarshalXMLAttr(attr)
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
			return fmt.Errorf("parsing AG_CoreAttributes: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_CoreAttributes and its children
func (m *AG_CoreAttributes) Validate() error {
	return m.ValidateWithPath("AG_CoreAttributes")
}

// ValidateWithPath validates the AG_CoreAttributes and its children, prefixing error messages with path
func (m *AG_CoreAttributes) ValidateWithPath(path string) error {
	if err := m.PrintAttr.ValidateWithPath(path + "/PrintAttr"); err != nil {
		return err
	}
	return nil
}
