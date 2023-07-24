package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_PathShadeProperties struct {
	PathAttr   ST_PathShadeType
	FillToRect *CT_RelativeRect
}

func NewCT_PathShadeProperties() *CT_PathShadeProperties {
	ret := &CT_PathShadeProperties{}
	return ret
}

func (m *CT_PathShadeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PathAttr != ST_PathShadeTypeUnset {
		attr, err := m.PathAttr.MarshalXMLAttr(xml.Name{Local: "path"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.FillToRect != nil {
		sefillToRect := xml.StartElement{Name: xml.Name{Local: "a:fillToRect"}}
		e.EncodeElement(m.FillToRect, sefillToRect)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PathShadeProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "path" {
			m.PathAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_PathShadeProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fillToRect"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fillToRect"}:
				m.FillToRect = NewCT_RelativeRect()
				if err := d.DecodeElement(m.FillToRect, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_PathShadeProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PathShadeProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PathShadeProperties and its children
func (m *CT_PathShadeProperties) Validate() error {
	return m.ValidateWithPath("CT_PathShadeProperties")
}

// ValidateWithPath validates the CT_PathShadeProperties and its children, prefixing error messages with path
func (m *CT_PathShadeProperties) ValidateWithPath(path string) error {
	if err := m.PathAttr.ValidateWithPath(path + "/PathAttr"); err != nil {
		return err
	}
	if m.FillToRect != nil {
		if err := m.FillToRect.ValidateWithPath(path + "/FillToRect"); err != nil {
			return err
		}
	}
	return nil
}
