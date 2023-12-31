package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcClippath struct {
	OfcCT_ClipPath
}

func NewOfcClippath() *OfcClippath {
	ret := &OfcClippath{}
	ret.OfcCT_ClipPath = *NewOfcCT_ClipPath()
	return ret
}

func (m *OfcClippath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:clippath"
	return m.OfcCT_ClipPath.MarshalXML(e, start)
}

func (m *OfcClippath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_ClipPath = *NewOfcCT_ClipPath()
	for _, attr := range start.Attr {
		if attr.Name.Local == "v" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcClippath: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcClippath and its children
func (m *OfcClippath) Validate() error {
	return m.ValidateWithPath("OfcClippath")
}

// ValidateWithPath validates the OfcClippath and its children, prefixing error messages with path
func (m *OfcClippath) ValidateWithPath(path string) error {
	if err := m.OfcCT_ClipPath.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
