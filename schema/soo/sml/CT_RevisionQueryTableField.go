package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_RevisionQueryTableField struct {
	// Sheet Id
	SheetIdAttr uint32
	// QueryTable Reference
	RefAttr string
	// Field Id
	FieldIdAttr uint32
}

func NewCT_RevisionQueryTableField() *CT_RevisionQueryTableField {
	ret := &CT_RevisionQueryTableField{}
	return ret
}

func (m *CT_RevisionQueryTableField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
		Value: fmt.Sprintf("%v", m.SheetIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
		Value: fmt.Sprintf("%v", m.RefAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fieldId"},
		Value: fmt.Sprintf("%v", m.FieldIdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionQueryTableField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SheetIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = parsed
			continue
		}
		if attr.Name.Local == "fieldId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.FieldIdAttr = uint32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RevisionQueryTableField: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_RevisionQueryTableField and its children
func (m *CT_RevisionQueryTableField) Validate() error {
	return m.ValidateWithPath("CT_RevisionQueryTableField")
}

// ValidateWithPath validates the CT_RevisionQueryTableField and its children, prefixing error messages with path
func (m *CT_RevisionQueryTableField) ValidateWithPath(path string) error {
	return nil
}
