package pml

import (
	"encoding/xml"
	"fmt"
)

type CT_SideDirectionTransition struct {
	// Direction
	DirAttr ST_TransitionSideDirectionType
}

func NewCT_SideDirectionTransition() *CT_SideDirectionTransition {
	ret := &CT_SideDirectionTransition{}
	return ret
}

func (m *CT_SideDirectionTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DirAttr != ST_TransitionSideDirectionTypeUnset {
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

func (m *CT_SideDirectionTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_SideDirectionTransition: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SideDirectionTransition and its children
func (m *CT_SideDirectionTransition) Validate() error {
	return m.ValidateWithPath("CT_SideDirectionTransition")
}

// ValidateWithPath validates the CT_SideDirectionTransition and its children, prefixing error messages with path
func (m *CT_SideDirectionTransition) ValidateWithPath(path string) error {
	if err := m.DirAttr.ValidateWithPath(path + "/DirAttr"); err != nil {
		return err
	}
	return nil
}
