package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_BaseStylesOverride struct {
	ClrScheme  *CT_ColorScheme
	FontScheme *CT_FontScheme
	FmtScheme  *CT_StyleMatrix
}

func NewCT_BaseStylesOverride() *CT_BaseStylesOverride {
	ret := &CT_BaseStylesOverride{}
	return ret
}

func (m *CT_BaseStylesOverride) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ClrScheme != nil {
		seclrScheme := xml.StartElement{Name: xml.Name{Local: "a:clrScheme"}}
		e.EncodeElement(m.ClrScheme, seclrScheme)
	}
	if m.FontScheme != nil {
		sefontScheme := xml.StartElement{Name: xml.Name{Local: "a:fontScheme"}}
		e.EncodeElement(m.FontScheme, sefontScheme)
	}
	if m.FmtScheme != nil {
		sefmtScheme := xml.StartElement{Name: xml.Name{Local: "a:fmtScheme"}}
		e.EncodeElement(m.FmtScheme, sefmtScheme)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BaseStylesOverride) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BaseStylesOverride:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "clrScheme"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "clrScheme"}:
				m.ClrScheme = NewCT_ColorScheme()
				if err := d.DecodeElement(m.ClrScheme, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fontScheme"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fontScheme"}:
				m.FontScheme = NewCT_FontScheme()
				if err := d.DecodeElement(m.FontScheme, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fmtScheme"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fmtScheme"}:
				m.FmtScheme = NewCT_StyleMatrix()
				if err := d.DecodeElement(m.FmtScheme, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_BaseStylesOverride %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BaseStylesOverride
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BaseStylesOverride and its children
func (m *CT_BaseStylesOverride) Validate() error {
	return m.ValidateWithPath("CT_BaseStylesOverride")
}

// ValidateWithPath validates the CT_BaseStylesOverride and its children, prefixing error messages with path
func (m *CT_BaseStylesOverride) ValidateWithPath(path string) error {
	if m.ClrScheme != nil {
		if err := m.ClrScheme.ValidateWithPath(path + "/ClrScheme"); err != nil {
			return err
		}
	}
	if m.FontScheme != nil {
		if err := m.FontScheme.ValidateWithPath(path + "/FontScheme"); err != nil {
			return err
		}
	}
	if m.FmtScheme != nil {
		if err := m.FmtScheme.ValidateWithPath(path + "/FmtScheme"); err != nil {
			return err
		}
	}
	return nil
}
