package pml

import (
	"encoding/xml"
	"fmt"
)

type CT_SplitTransition struct {
	// Orientation
	OrientAttr ST_Direction
	// Direction
	DirAttr ST_TransitionInOutDirectionType
}

func NewCT_SplitTransition() *CT_SplitTransition {
	ret := &CT_SplitTransition{}
	return ret
}

func (m *CT_SplitTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OrientAttr != ST_DirectionUnset {
		attr, err := m.OrientAttr.MarshalXMLAttr(xml.Name{Local: "orient"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DirAttr != ST_TransitionInOutDirectionTypeUnset {
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

func (m *CT_SplitTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "orient" {
			m.OrientAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dir" {
			m.DirAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SplitTransition: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SplitTransition and its children
func (m *CT_SplitTransition) Validate() error {
	return m.ValidateWithPath("CT_SplitTransition")
}

// ValidateWithPath validates the CT_SplitTransition and its children, prefixing error messages with path
func (m *CT_SplitTransition) ValidateWithPath(path string) error {
	if err := m.OrientAttr.ValidateWithPath(path + "/OrientAttr"); err != nil {
		return err
	}
	if err := m.DirAttr.ValidateWithPath(path + "/DirAttr"); err != nil {
		return err
	}
	return nil
}
