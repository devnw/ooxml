package dml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_OfficeArtExtension struct {
	UriAttr string
	Any     []office.Any
}

func NewCT_OfficeArtExtension() *CT_OfficeArtExtension {
	ret := &CT_OfficeArtExtension{}
	return ret
}

func (m *CT_OfficeArtExtension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uri"},
		Value: fmt.Sprintf("%v", m.UriAttr)})
	e.EncodeToken(start)
	if m.Any != nil {
		for _, c := range m.Any {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OfficeArtExtension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uri" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UriAttr = parsed
			continue
		}
	}
lCT_OfficeArtExtension:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			default:
				if anyEl, err := office.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = append(m.Any, anyEl)
				}
			}
		case xml.EndElement:
			break lCT_OfficeArtExtension
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OfficeArtExtension and its children
func (m *CT_OfficeArtExtension) Validate() error {
	return m.ValidateWithPath("CT_OfficeArtExtension")
}

// ValidateWithPath validates the CT_OfficeArtExtension and its children, prefixing error messages with path
func (m *CT_OfficeArtExtension) ValidateWithPath(path string) error {
	return nil
}
