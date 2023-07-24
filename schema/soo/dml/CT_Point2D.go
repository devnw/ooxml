package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_Point2D struct {
	XAttr ST_Coordinate
	YAttr ST_Coordinate
}

func NewCT_Point2D() *CT_Point2D {
	ret := &CT_Point2D{}
	return ret
}

func (m *CT_Point2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "x"},
		Value: fmt.Sprintf("%v", m.XAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "y"},
		Value: fmt.Sprintf("%v", m.YAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Point2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "x" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.XAttr = parsed
			continue
		}
		if attr.Name.Local == "y" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.YAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Point2D: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Point2D and its children
func (m *CT_Point2D) Validate() error {
	return m.ValidateWithPath("CT_Point2D")
}

// ValidateWithPath validates the CT_Point2D and its children, prefixing error messages with path
func (m *CT_Point2D) ValidateWithPath(path string) error {
	if err := m.XAttr.ValidateWithPath(path + "/XAttr"); err != nil {
		return err
	}
	if err := m.YAttr.ValidateWithPath(path + "/YAttr"); err != nil {
		return err
	}
	return nil
}
