package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_ShapeStyle struct {
	LnRef     *CT_StyleMatrixReference
	FillRef   *CT_StyleMatrixReference
	EffectRef *CT_StyleMatrixReference
	FontRef   *CT_FontReference
}

func NewCT_ShapeStyle() *CT_ShapeStyle {
	ret := &CT_ShapeStyle{}
	ret.LnRef = NewCT_StyleMatrixReference()
	ret.FillRef = NewCT_StyleMatrixReference()
	ret.EffectRef = NewCT_StyleMatrixReference()
	ret.FontRef = NewCT_FontReference()
	return ret
}

func (m *CT_ShapeStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	selnRef := xml.StartElement{Name: xml.Name{Local: "a:lnRef"}}
	e.EncodeElement(m.LnRef, selnRef)
	sefillRef := xml.StartElement{Name: xml.Name{Local: "a:fillRef"}}
	e.EncodeElement(m.FillRef, sefillRef)
	seeffectRef := xml.StartElement{Name: xml.Name{Local: "a:effectRef"}}
	e.EncodeElement(m.EffectRef, seeffectRef)
	sefontRef := xml.StartElement{Name: xml.Name{Local: "a:fontRef"}}
	e.EncodeElement(m.FontRef, sefontRef)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ShapeStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.LnRef = NewCT_StyleMatrixReference()
	m.FillRef = NewCT_StyleMatrixReference()
	m.EffectRef = NewCT_StyleMatrixReference()
	m.FontRef = NewCT_FontReference()
lCT_ShapeStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lnRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lnRef"}:
				if err := d.DecodeElement(m.LnRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fillRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fillRef"}:
				if err := d.DecodeElement(m.FillRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effectRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "effectRef"}:
				if err := d.DecodeElement(m.EffectRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fontRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fontRef"}:
				if err := d.DecodeElement(m.FontRef, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_ShapeStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ShapeStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ShapeStyle and its children
func (m *CT_ShapeStyle) Validate() error {
	return m.ValidateWithPath("CT_ShapeStyle")
}

// ValidateWithPath validates the CT_ShapeStyle and its children, prefixing error messages with path
func (m *CT_ShapeStyle) ValidateWithPath(path string) error {
	if err := m.LnRef.ValidateWithPath(path + "/LnRef"); err != nil {
		return err
	}
	if err := m.FillRef.ValidateWithPath(path + "/FillRef"); err != nil {
		return err
	}
	if err := m.EffectRef.ValidateWithPath(path + "/EffectRef"); err != nil {
		return err
	}
	if err := m.FontRef.ValidateWithPath(path + "/FontRef"); err != nil {
		return err
	}
	return nil
}
