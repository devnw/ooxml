package diagram

import (
	"encoding/xml"
	"fmt"
)

type RelIds struct {
	CT_RelIds
}

func NewRelIds() *RelIds {
	ret := &RelIds{}
	ret.CT_RelIds = *NewCT_RelIds()
	return ret
}

func (m *RelIds) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:di"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "relIds"
	return m.CT_RelIds.MarshalXML(e, start)
}

func (m *RelIds) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_RelIds = *NewCT_RelIds()
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "dm" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "dm" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DmAttr = parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "lo" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "lo" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LoAttr = parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "qs" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "qs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.QsAttr = parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "cs" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "cs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CsAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing RelIds: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the RelIds and its children
func (m *RelIds) Validate() error {
	return m.ValidateWithPath("RelIds")
}

// ValidateWithPath validates the RelIds and its children, prefixing error messages with path
func (m *RelIds) ValidateWithPath(path string) error {
	if err := m.CT_RelIds.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
