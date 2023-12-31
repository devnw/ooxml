package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcInk struct {
	OfcCT_Ink
}

func NewOfcInk() *OfcInk {
	ret := &OfcInk{}
	ret.OfcCT_Ink = *NewOfcCT_Ink()
	return ret
}

func (m *OfcInk) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:ink"
	return m.OfcCT_Ink.MarshalXML(e, start)
}

func (m *OfcInk) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_Ink = *NewOfcCT_Ink()
	for _, attr := range start.Attr {
		if attr.Name.Local == "i" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IAttr = &parsed
			continue
		}
		if attr.Name.Local == "annotation" {
			m.AnnotationAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "contentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentTypeAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcInk: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcInk and its children
func (m *OfcInk) Validate() error {
	return m.ValidateWithPath("OfcInk")
}

// ValidateWithPath validates the OfcInk and its children, prefixing error messages with path
func (m *OfcInk) ValidateWithPath(path string) error {
	if err := m.OfcCT_Ink.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
