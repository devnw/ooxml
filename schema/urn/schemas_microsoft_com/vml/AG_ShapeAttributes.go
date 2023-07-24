package vml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml/schema/soo/ofc/sharedTypes"
)

type AG_ShapeAttributes struct {
	OpacityAttr      *string
	StrokedAttr      sharedTypes.ST_TrueFalse
	StrokecolorAttr  *string
	StrokeweightAttr *string
	InsetpenAttr     sharedTypes.ST_TrueFalse
	ChromakeyAttr    *string
	FilledAttr       sharedTypes.ST_TrueFalse
	FillcolorAttr    *string
}

func NewAG_ShapeAttributes() *AG_ShapeAttributes {
	ret := &AG_ShapeAttributes{}
	return ret
}

func (m *AG_ShapeAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OpacityAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "opacity"},
			Value: fmt.Sprintf("%v", *m.OpacityAttr)})
	}
	if m.StrokedAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.StrokedAttr.MarshalXMLAttr(xml.Name{Local: "stroked"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StrokecolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strokecolor"},
			Value: fmt.Sprintf("%v", *m.StrokecolorAttr)})
	}
	if m.StrokeweightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strokeweight"},
			Value: fmt.Sprintf("%v", *m.StrokeweightAttr)})
	}
	if m.InsetpenAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.InsetpenAttr.MarshalXMLAttr(xml.Name{Local: "insetpen"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ChromakeyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "chromakey"},
			Value: fmt.Sprintf("%v", *m.ChromakeyAttr)})
	}
	if m.FilledAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.FilledAttr.MarshalXMLAttr(xml.Name{Local: "filled"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FillcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillcolor"},
			Value: fmt.Sprintf("%v", *m.FillcolorAttr)})
	}
	return nil
}

func (m *AG_ShapeAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "opacity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OpacityAttr = &parsed
			continue
		}
		if attr.Name.Local == "stroked" {
			m.StrokedAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "strokecolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokecolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "strokeweight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokeweightAttr = &parsed
			continue
		}
		if attr.Name.Local == "insetpen" {
			m.InsetpenAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "chromakey" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ChromakeyAttr = &parsed
			continue
		}
		if attr.Name.Local == "filled" {
			m.FilledAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "fillcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FillcolorAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_ShapeAttributes: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_ShapeAttributes and its children
func (m *AG_ShapeAttributes) Validate() error {
	return m.ValidateWithPath("AG_ShapeAttributes")
}

// ValidateWithPath validates the AG_ShapeAttributes and its children, prefixing error messages with path
func (m *AG_ShapeAttributes) ValidateWithPath(path string) error {
	if err := m.StrokedAttr.ValidateWithPath(path + "/StrokedAttr"); err != nil {
		return err
	}
	if err := m.InsetpenAttr.ValidateWithPath(path + "/InsetpenAttr"); err != nil {
		return err
	}
	if err := m.FilledAttr.ValidateWithPath(path + "/FilledAttr"); err != nil {
		return err
	}
	return nil
}
