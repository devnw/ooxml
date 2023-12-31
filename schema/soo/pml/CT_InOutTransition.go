package pml

import (
	"encoding/xml"
	"fmt"
)

type CT_InOutTransition struct {
	// Direction
	DirAttr ST_TransitionInOutDirectionType
}

func NewCT_InOutTransition() *CT_InOutTransition {
	ret := &CT_InOutTransition{}
	return ret
}

func (m *CT_InOutTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

func (m *CT_InOutTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_InOutTransition: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_InOutTransition and its children
func (m *CT_InOutTransition) Validate() error {
	return m.ValidateWithPath("CT_InOutTransition")
}

// ValidateWithPath validates the CT_InOutTransition and its children, prefixing error messages with path
func (m *CT_InOutTransition) ValidateWithPath(path string) error {
	if err := m.DirAttr.ValidateWithPath(path + "/DirAttr"); err != nil {
		return err
	}
	return nil
}
