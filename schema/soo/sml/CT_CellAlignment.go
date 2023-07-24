package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_CellAlignment struct {
	// Horizontal Alignment
	HorizontalAttr ST_HorizontalAlignment
	// Vertical Alignment
	VerticalAttr ST_VerticalAlignment
	// Text Rotation
	TextRotationAttr *uint8
	// Wrap Text
	WrapTextAttr *bool
	// Indent
	IndentAttr *uint32
	// Relative Indent
	RelativeIndentAttr *int32
	// Justify Last Line
	JustifyLastLineAttr *bool
	// Shrink To Fit
	ShrinkToFitAttr *bool
	// Reading Order
	ReadingOrderAttr *uint32
}

func NewCT_CellAlignment() *CT_CellAlignment {
	ret := &CT_CellAlignment{}
	return ret
}

func (m *CT_CellAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.HorizontalAttr != ST_HorizontalAlignmentUnset {
		attr, err := m.HorizontalAttr.MarshalXMLAttr(xml.Name{Local: "horizontal"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.VerticalAttr != ST_VerticalAlignmentUnset {
		attr, err := m.VerticalAttr.MarshalXMLAttr(xml.Name{Local: "vertical"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TextRotationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "textRotation"},
			Value: fmt.Sprintf("%v", *m.TextRotationAttr)})
	}
	if m.WrapTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "wrapText"},
			Value: fmt.Sprintf("%d", b2i(*m.WrapTextAttr))})
	}
	if m.IndentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "indent"},
			Value: fmt.Sprintf("%v", *m.IndentAttr)})
	}
	if m.RelativeIndentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeIndent"},
			Value: fmt.Sprintf("%v", *m.RelativeIndentAttr)})
	}
	if m.JustifyLastLineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "justifyLastLine"},
			Value: fmt.Sprintf("%d", b2i(*m.JustifyLastLineAttr))})
	}
	if m.ShrinkToFitAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "shrinkToFit"},
			Value: fmt.Sprintf("%d", b2i(*m.ShrinkToFitAttr))})
	}
	if m.ReadingOrderAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "readingOrder"},
			Value: fmt.Sprintf("%v", *m.ReadingOrderAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "horizontal" {
			m.HorizontalAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "vertical" {
			m.VerticalAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "textRotation" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.TextRotationAttr = &pt
			continue
		}
		if attr.Name.Local == "wrapText" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.WrapTextAttr = &parsed
			continue
		}
		if attr.Name.Local == "indent" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.IndentAttr = &pt
			continue
		}
		if attr.Name.Local == "relativeIndent" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.RelativeIndentAttr = &pt
			continue
		}
		if attr.Name.Local == "justifyLastLine" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.JustifyLastLineAttr = &parsed
			continue
		}
		if attr.Name.Local == "shrinkToFit" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShrinkToFitAttr = &parsed
			continue
		}
		if attr.Name.Local == "readingOrder" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ReadingOrderAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CellAlignment: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CellAlignment and its children
func (m *CT_CellAlignment) Validate() error {
	return m.ValidateWithPath("CT_CellAlignment")
}

// ValidateWithPath validates the CT_CellAlignment and its children, prefixing error messages with path
func (m *CT_CellAlignment) ValidateWithPath(path string) error {
	if err := m.HorizontalAttr.ValidateWithPath(path + "/HorizontalAttr"); err != nil {
		return err
	}
	if err := m.VerticalAttr.ValidateWithPath(path + "/VerticalAttr"); err != nil {
		return err
	}
	return nil
}
