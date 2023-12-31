package pml

import (
	"encoding/xml"
	"fmt"
)

type CT_OrientationTransition struct {
	// Transition Direction
	DirAttr ST_Direction
}

func NewCT_OrientationTransition() *CT_OrientationTransition {
	ret := &CT_OrientationTransition{}
	return ret
}

func (m *CT_OrientationTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DirAttr != ST_DirectionUnset {
		attr, err := m.DirAttr.MarshalXMLAttr(xml.Name{Local: "dir"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OrientationTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "dir" {
			m.DirAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_OrientationTransition: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OrientationTransition and its children
func (m *CT_OrientationTransition) Validate() error {
	return m.ValidateWithPath("CT_OrientationTransition")
}

// ValidateWithPath validates the CT_OrientationTransition and its children, prefixing error messages with path
func (m *CT_OrientationTransition) ValidateWithPath(path string) error {
	if err := m.DirAttr.ValidateWithPath(path + "/DirAttr"); err != nil {
		return err
	}
	return nil
}
