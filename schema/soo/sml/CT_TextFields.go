package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_TextFields struct {
	// Count of Fields
	CountAttr *uint32
	// Text Import Field Settings
	TextField []*CT_TextField
}

func NewCT_TextFields() *CT_TextFields {
	ret := &CT_TextFields{}
	return ret
}

func (m *CT_TextFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	setextField := xml.StartElement{Name: xml.Name{Local: "ma:textField"}}
	for _, c := range m.TextField {
		e.EncodeElement(c, setextField)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextFields) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_TextFields:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "textField"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "textField"}:
				tmp := NewCT_TextField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TextField = append(m.TextField, tmp)
			default:
				office.Log("skipping unsupported element on CT_TextFields %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextFields
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextFields and its children
func (m *CT_TextFields) Validate() error {
	return m.ValidateWithPath("CT_TextFields")
}

// ValidateWithPath validates the CT_TextFields and its children, prefixing error messages with path
func (m *CT_TextFields) ValidateWithPath(path string) error {
	for i, v := range m.TextField {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TextField[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
