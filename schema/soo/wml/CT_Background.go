package wml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_Background struct {
	// Background Color
	ColorAttr *ST_HexColor
	// Background Theme Color
	ThemeColorAttr ST_ThemeColor
	// Background Theme Color Tint
	ThemeTintAttr *string
	// Background Theme Color Shade
	ThemeShadeAttr *string
	Drawing        *CT_Drawing
}

func NewCT_Background() *CT_Background {
	ret := &CT_Background{}
	return ret
}

func (m *CT_Background) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ColorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:color"},
			Value: fmt.Sprintf("%v", *m.ColorAttr)})
	}
	if m.ThemeColorAttr != ST_ThemeColorUnset {
		attr, err := m.ThemeColorAttr.MarshalXMLAttr(xml.Name{Local: "w:themeColor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ThemeTintAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeTint"},
			Value: fmt.Sprintf("%v", *m.ThemeTintAttr)})
	}
	if m.ThemeShadeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeShade"},
			Value: fmt.Sprintf("%v", *m.ThemeShadeAttr)})
	}
	e.EncodeToken(start)
	if m.Drawing != nil {
		sedrawing := xml.StartElement{Name: xml.Name{Local: "w:drawing"}}
		e.EncodeElement(m.Drawing, sedrawing)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Background) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "color" {
			parsed, err := ParseUnionST_HexColor(attr.Value)
			if err != nil {
				return err
			}
			m.ColorAttr = &parsed
			continue
		}
		if attr.Name.Local == "themeColor" {
			m.ThemeColorAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "themeTint" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThemeTintAttr = &parsed
			continue
		}
		if attr.Name.Local == "themeShade" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThemeShadeAttr = &parsed
			continue
		}
	}
lCT_Background:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "drawing"}:
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_Background %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Background
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Background and its children
func (m *CT_Background) Validate() error {
	return m.ValidateWithPath("CT_Background")
}

// ValidateWithPath validates the CT_Background and its children, prefixing error messages with path
func (m *CT_Background) ValidateWithPath(path string) error {
	if m.ColorAttr != nil {
		if err := m.ColorAttr.ValidateWithPath(path + "/ColorAttr"); err != nil {
			return err
		}
	}
	if err := m.ThemeColorAttr.ValidateWithPath(path + "/ThemeColorAttr"); err != nil {
		return err
	}
	if m.Drawing != nil {
		if err := m.Drawing.ValidateWithPath(path + "/Drawing"); err != nil {
			return err
		}
	}
	return nil
}
