package pml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_TLByAnimateColorTransform struct {
	// RGB
	Rgb *CT_TLByRgbColorTransform
	// HSL
	Hsl *CT_TLByHslColorTransform
}

func NewCT_TLByAnimateColorTransform() *CT_TLByAnimateColorTransform {
	ret := &CT_TLByAnimateColorTransform{}
	return ret
}

func (m *CT_TLByAnimateColorTransform) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Rgb != nil {
		sergb := xml.StartElement{Name: xml.Name{Local: "p:rgb"}}
		e.EncodeElement(m.Rgb, sergb)
	}
	if m.Hsl != nil {
		sehsl := xml.StartElement{Name: xml.Name{Local: "p:hsl"}}
		e.EncodeElement(m.Hsl, sehsl)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLByAnimateColorTransform) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TLByAnimateColorTransform:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "rgb"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "rgb"}:
				m.Rgb = NewCT_TLByRgbColorTransform()
				if err := d.DecodeElement(m.Rgb, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "hsl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "hsl"}:
				m.Hsl = NewCT_TLByHslColorTransform()
				if err := d.DecodeElement(m.Hsl, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_TLByAnimateColorTransform %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLByAnimateColorTransform
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLByAnimateColorTransform and its children
func (m *CT_TLByAnimateColorTransform) Validate() error {
	return m.ValidateWithPath("CT_TLByAnimateColorTransform")
}

// ValidateWithPath validates the CT_TLByAnimateColorTransform and its children, prefixing error messages with path
func (m *CT_TLByAnimateColorTransform) ValidateWithPath(path string) error {
	if m.Rgb != nil {
		if err := m.Rgb.ValidateWithPath(path + "/Rgb"); err != nil {
			return err
		}
	}
	if m.Hsl != nil {
		if err := m.Hsl.ValidateWithPath(path + "/Hsl"); err != nil {
			return err
		}
	}
	return nil
}
