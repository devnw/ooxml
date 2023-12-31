package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcComplex struct {
	OfcCT_Complex
}

func NewOfcComplex() *OfcComplex {
	ret := &OfcComplex{}
	ret.OfcCT_Complex = *NewOfcCT_Complex()
	return ret
}

func (m *OfcComplex) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:complex"
	return m.OfcCT_Complex.MarshalXML(e, start)
}

func (m *OfcComplex) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_Complex = *NewOfcCT_Complex()
	for _, attr := range start.Attr {
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcComplex: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcComplex and its children
func (m *OfcComplex) Validate() error {
	return m.ValidateWithPath("OfcComplex")
}

// ValidateWithPath validates the OfcComplex and its children, prefixing error messages with path
func (m *OfcComplex) ValidateWithPath(path string) error {
	if err := m.OfcCT_Complex.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
