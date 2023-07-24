package sml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_Extension struct {
	// URI
	UriAttr *string
	Any     office.Any
}

func NewCT_Extension() *CT_Extension {
	ret := &CT_Extension{}
	return ret
}

func (m *CT_Extension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UriAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uri"},
			Value: fmt.Sprintf("%v", *m.UriAttr)})
	}
	e.EncodeToken(start)
	if m.Any != nil {
		m.Any.MarshalXML(e, xml.StartElement{})
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Extension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uri" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UriAttr = &parsed
			continue
		}
	}
lCT_Extension:
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
					m.Any = anyEl
				}
			}
		case xml.EndElement:
			break lCT_Extension
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Extension and its children
func (m *CT_Extension) Validate() error {
	return m.ValidateWithPath("CT_Extension")
}

// ValidateWithPath validates the CT_Extension and its children, prefixing error messages with path
func (m *CT_Extension) ValidateWithPath(path string) error {
	return nil
}
