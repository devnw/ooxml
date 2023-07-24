package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_GeomRect struct {
	LAttr ST_AdjCoordinate
	TAttr ST_AdjCoordinate
	RAttr ST_AdjCoordinate
	BAttr ST_AdjCoordinate
}

func NewCT_GeomRect() *CT_GeomRect {
	ret := &CT_GeomRect{}
	return ret
}

func (m *CT_GeomRect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "l"},
		Value: fmt.Sprintf("%v", m.LAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "t"},
		Value: fmt.Sprintf("%v", m.TAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
		Value: fmt.Sprintf("%v", m.BAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GeomRect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "l" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.LAttr = parsed
			continue
		}
		if attr.Name.Local == "t" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.TAttr = parsed
			continue
		}
		if attr.Name.Local == "r" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.RAttr = parsed
			continue
		}
		if attr.Name.Local == "b" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.BAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_GeomRect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_GeomRect and its children
func (m *CT_GeomRect) Validate() error {
	return m.ValidateWithPath("CT_GeomRect")
}

// ValidateWithPath validates the CT_GeomRect and its children, prefixing error messages with path
func (m *CT_GeomRect) ValidateWithPath(path string) error {
	if err := m.LAttr.ValidateWithPath(path + "/LAttr"); err != nil {
		return err
	}
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	if err := m.RAttr.ValidateWithPath(path + "/RAttr"); err != nil {
		return err
	}
	if err := m.BAttr.ValidateWithPath(path + "/BAttr"); err != nil {
		return err
	}
	return nil
}
