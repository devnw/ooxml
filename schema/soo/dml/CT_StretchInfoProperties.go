package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_StretchInfoProperties struct {
	FillRect *CT_RelativeRect
}

func NewCT_StretchInfoProperties() *CT_StretchInfoProperties {
	ret := &CT_StretchInfoProperties{}
	return ret
}

func (m *CT_StretchInfoProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.FillRect != nil {
		sefillRect := xml.StartElement{Name: xml.Name{Local: "a:fillRect"}}
		e.EncodeElement(m.FillRect, sefillRect)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_StretchInfoProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_StretchInfoProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fillRect"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fillRect"}:
				m.FillRect = NewCT_RelativeRect()
				if err := d.DecodeElement(m.FillRect, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_StretchInfoProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_StretchInfoProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_StretchInfoProperties and its children
func (m *CT_StretchInfoProperties) Validate() error {
	return m.ValidateWithPath("CT_StretchInfoProperties")
}

// ValidateWithPath validates the CT_StretchInfoProperties and its children, prefixing error messages with path
func (m *CT_StretchInfoProperties) ValidateWithPath(path string) error {
	if m.FillRect != nil {
		if err := m.FillRect.ValidateWithPath(path + "/FillRect"); err != nil {
			return err
		}
	}
	return nil
}
