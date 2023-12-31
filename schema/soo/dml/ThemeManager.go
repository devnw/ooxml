package dml

import (
	"encoding/xml"
	"fmt"
)

type ThemeManager struct {
	CT_EmptyElement
}

func NewThemeManager() *ThemeManager {
	ret := &ThemeManager{}
	ret.CT_EmptyElement = *NewCT_EmptyElement()
	return ret
}

func (m *ThemeManager) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "a:themeManager"
	return m.CT_EmptyElement.MarshalXML(e, start)
}

func (m *ThemeManager) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_EmptyElement = *NewCT_EmptyElement()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing ThemeManager: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the ThemeManager and its children
func (m *ThemeManager) Validate() error {
	return m.ValidateWithPath("ThemeManager")
}

// ValidateWithPath validates the ThemeManager and its children, prefixing error messages with path
func (m *ThemeManager) ValidateWithPath(path string) error {
	if err := m.CT_EmptyElement.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
