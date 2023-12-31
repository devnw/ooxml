package diagram

import (
	"encoding/xml"
	"fmt"
)

type CT_ResizeHandles struct {
	ValAttr ST_ResizeHandlesStr
}

func NewCT_ResizeHandles() *CT_ResizeHandles {
	ret := &CT_ResizeHandles{}
	return ret
}

func (m *CT_ResizeHandles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_ResizeHandlesStrUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ResizeHandles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ResizeHandles: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ResizeHandles and its children
func (m *CT_ResizeHandles) Validate() error {
	return m.ValidateWithPath("CT_ResizeHandles")
}

// ValidateWithPath validates the CT_ResizeHandles and its children, prefixing error messages with path
func (m *CT_ResizeHandles) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
