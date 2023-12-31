package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_ColorSchemeMapping struct {
	// Background 1 Theme Color Mapping
	Bg1Attr ST_WmlColorSchemeIndex
	// Text 1 Theme Color Mapping
	T1Attr ST_WmlColorSchemeIndex
	// Background 2 Theme Color Mapping
	Bg2Attr ST_WmlColorSchemeIndex
	// Text 2 Theme Color Mapping
	T2Attr ST_WmlColorSchemeIndex
	// Accent 1 Theme Color Mapping
	Accent1Attr ST_WmlColorSchemeIndex
	// Accent 2 Theme Color Mapping
	Accent2Attr ST_WmlColorSchemeIndex
	// Accent3 Theme Color Mapping
	Accent3Attr ST_WmlColorSchemeIndex
	// Accent4 Theme Color Mapping
	Accent4Attr ST_WmlColorSchemeIndex
	// Accent5 Theme Color Mapping
	Accent5Attr ST_WmlColorSchemeIndex
	// Accent6 Theme Color Mapping
	Accent6Attr ST_WmlColorSchemeIndex
	// Hyperlink Theme Color Mapping
	HyperlinkAttr ST_WmlColorSchemeIndex
	// Followed Hyperlink Theme Color Mapping
	FollowedHyperlinkAttr ST_WmlColorSchemeIndex
}

func NewCT_ColorSchemeMapping() *CT_ColorSchemeMapping {
	ret := &CT_ColorSchemeMapping{}
	return ret
}

func (m *CT_ColorSchemeMapping) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Bg1Attr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.Bg1Attr.MarshalXMLAttr(xml.Name{Local: "w:bg1"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.T1Attr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.T1Attr.MarshalXMLAttr(xml.Name{Local: "w:t1"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.Bg2Attr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.Bg2Attr.MarshalXMLAttr(xml.Name{Local: "w:bg2"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.T2Attr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.T2Attr.MarshalXMLAttr(xml.Name{Local: "w:t2"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.Accent1Attr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.Accent1Attr.MarshalXMLAttr(xml.Name{Local: "w:accent1"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.Accent2Attr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.Accent2Attr.MarshalXMLAttr(xml.Name{Local: "w:accent2"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.Accent3Attr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.Accent3Attr.MarshalXMLAttr(xml.Name{Local: "w:accent3"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.Accent4Attr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.Accent4Attr.MarshalXMLAttr(xml.Name{Local: "w:accent4"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.Accent5Attr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.Accent5Attr.MarshalXMLAttr(xml.Name{Local: "w:accent5"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.Accent6Attr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.Accent6Attr.MarshalXMLAttr(xml.Name{Local: "w:accent6"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HyperlinkAttr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.HyperlinkAttr.MarshalXMLAttr(xml.Name{Local: "w:hyperlink"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FollowedHyperlinkAttr != ST_WmlColorSchemeIndexUnset {
		attr, err := m.FollowedHyperlinkAttr.MarshalXMLAttr(xml.Name{Local: "w:followedHyperlink"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorSchemeMapping) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "bg1" {
			m.Bg1Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "t1" {
			m.T1Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "bg2" {
			m.Bg2Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "t2" {
			m.T2Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent1" {
			m.Accent1Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent2" {
			m.Accent2Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent3" {
			m.Accent3Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent4" {
			m.Accent4Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent5" {
			m.Accent5Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent6" {
			m.Accent6Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "hyperlink" {
			m.HyperlinkAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "followedHyperlink" {
			m.FollowedHyperlinkAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ColorSchemeMapping: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ColorSchemeMapping and its children
func (m *CT_ColorSchemeMapping) Validate() error {
	return m.ValidateWithPath("CT_ColorSchemeMapping")
}

// ValidateWithPath validates the CT_ColorSchemeMapping and its children, prefixing error messages with path
func (m *CT_ColorSchemeMapping) ValidateWithPath(path string) error {
	if err := m.Bg1Attr.ValidateWithPath(path + "/Bg1Attr"); err != nil {
		return err
	}
	if err := m.T1Attr.ValidateWithPath(path + "/T1Attr"); err != nil {
		return err
	}
	if err := m.Bg2Attr.ValidateWithPath(path + "/Bg2Attr"); err != nil {
		return err
	}
	if err := m.T2Attr.ValidateWithPath(path + "/T2Attr"); err != nil {
		return err
	}
	if err := m.Accent1Attr.ValidateWithPath(path + "/Accent1Attr"); err != nil {
		return err
	}
	if err := m.Accent2Attr.ValidateWithPath(path + "/Accent2Attr"); err != nil {
		return err
	}
	if err := m.Accent3Attr.ValidateWithPath(path + "/Accent3Attr"); err != nil {
		return err
	}
	if err := m.Accent4Attr.ValidateWithPath(path + "/Accent4Attr"); err != nil {
		return err
	}
	if err := m.Accent5Attr.ValidateWithPath(path + "/Accent5Attr"); err != nil {
		return err
	}
	if err := m.Accent6Attr.ValidateWithPath(path + "/Accent6Attr"); err != nil {
		return err
	}
	if err := m.HyperlinkAttr.ValidateWithPath(path + "/HyperlinkAttr"); err != nil {
		return err
	}
	if err := m.FollowedHyperlinkAttr.ValidateWithPath(path + "/FollowedHyperlinkAttr"); err != nil {
		return err
	}
	return nil
}
