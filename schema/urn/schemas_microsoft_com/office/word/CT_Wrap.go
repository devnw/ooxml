package word

import (
	"encoding/xml"
	"fmt"
)

type CT_Wrap struct {
	TypeAttr    ST_WrapType
	SideAttr    ST_WrapSide
	AnchorxAttr ST_HorizontalAnchor
	AnchoryAttr ST_VerticalAnchor
}

func NewCT_Wrap() *CT_Wrap {
	ret := &CT_Wrap{}
	return ret
}

func (m *CT_Wrap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_WrapTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SideAttr != ST_WrapSideUnset {
		attr, err := m.SideAttr.MarshalXMLAttr(xml.Name{Local: "side"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AnchorxAttr != ST_HorizontalAnchorUnset {
		attr, err := m.AnchorxAttr.MarshalXMLAttr(xml.Name{Local: "anchorx"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AnchoryAttr != ST_VerticalAnchorUnset {
		attr, err := m.AnchoryAttr.MarshalXMLAttr(xml.Name{Local: "anchory"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Wrap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "side" {
			m.SideAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "anchorx" {
			m.AnchorxAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "anchory" {
			m.AnchoryAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Wrap: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Wrap and its children
func (m *CT_Wrap) Validate() error {
	return m.ValidateWithPath("CT_Wrap")
}

// ValidateWithPath validates the CT_Wrap and its children, prefixing error messages with path
func (m *CT_Wrap) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.SideAttr.ValidateWithPath(path + "/SideAttr"); err != nil {
		return err
	}
	if err := m.AnchorxAttr.ValidateWithPath(path + "/AnchorxAttr"); err != nil {
		return err
	}
	if err := m.AnchoryAttr.ValidateWithPath(path + "/AnchoryAttr"); err != nil {
		return err
	}
	return nil
}
