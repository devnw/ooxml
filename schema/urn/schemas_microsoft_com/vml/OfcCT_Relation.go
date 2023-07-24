package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcCT_Relation struct {
	IdsrcAttr  *string
	IddestAttr *string
	IdcntrAttr *string
	ExtAttr    ST_Ext
}

func NewOfcCT_Relation() *OfcCT_Relation {
	ret := &OfcCT_Relation{}
	return ret
}

func (m *OfcCT_Relation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdsrcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idsrc"},
			Value: fmt.Sprintf("%v", *m.IdsrcAttr)})
	}
	if m.IddestAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "iddest"},
			Value: fmt.Sprintf("%v", *m.IddestAttr)})
	}
	if m.IdcntrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idcntr"},
			Value: fmt.Sprintf("%v", *m.IdcntrAttr)})
	}
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_Relation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "idsrc" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdsrcAttr = &parsed
			continue
		}
		if attr.Name.Local == "iddest" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IddestAttr = &parsed
			continue
		}
		if attr.Name.Local == "idcntr" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdcntrAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_Relation: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_Relation and its children
func (m *OfcCT_Relation) Validate() error {
	return m.ValidateWithPath("OfcCT_Relation")
}

// ValidateWithPath validates the OfcCT_Relation and its children, prefixing error messages with path
func (m *OfcCT_Relation) ValidateWithPath(path string) error {
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
