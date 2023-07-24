package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Tuple struct {
	// Field Index
	FldAttr *uint32
	// Hierarchy Index
	HierAttr *uint32
	// Item Index
	ItemAttr uint32
}

func NewCT_Tuple() *CT_Tuple {
	ret := &CT_Tuple{}
	return ret
}

func (m *CT_Tuple) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fld"},
			Value: fmt.Sprintf("%v", *m.FldAttr)})
	}
	if m.HierAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hier"},
			Value: fmt.Sprintf("%v", *m.HierAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "item"},
		Value: fmt.Sprintf("%v", m.ItemAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Tuple) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fld" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.FldAttr = &pt
			continue
		}
		if attr.Name.Local == "hier" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.HierAttr = &pt
			continue
		}
		if attr.Name.Local == "item" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ItemAttr = uint32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Tuple: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Tuple and its children
func (m *CT_Tuple) Validate() error {
	return m.ValidateWithPath("CT_Tuple")
}

// ValidateWithPath validates the CT_Tuple and its children, prefixing error messages with path
func (m *CT_Tuple) ValidateWithPath(path string) error {
	return nil
}
