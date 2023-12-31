package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcFill struct {
	OfcCT_Fill
}

func NewOfcFill() *OfcFill {
	ret := &OfcFill{}
	ret.OfcCT_Fill = *NewOfcCT_Fill()
	return ret
}

func (m *OfcFill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:fill"
	return m.OfcCT_Fill.MarshalXML(e, start)
}

func (m *OfcFill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_Fill = *NewOfcCT_Fill()
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcFill: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcFill and its children
func (m *OfcFill) Validate() error {
	return m.ValidateWithPath("OfcFill")
}

// ValidateWithPath validates the OfcFill and its children, prefixing error messages with path
func (m *OfcFill) ValidateWithPath(path string) error {
	if err := m.OfcCT_Fill.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
