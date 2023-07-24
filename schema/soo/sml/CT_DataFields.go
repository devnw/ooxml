package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_DataFields struct {
	// Data Items Count
	CountAttr *uint32
	// Data Field Item
	DataField []*CT_DataField
}

func NewCT_DataFields() *CT_DataFields {
	ret := &CT_DataFields{}
	return ret
}

func (m *CT_DataFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sedataField := xml.StartElement{Name: xml.Name{Local: "ma:dataField"}}
	for _, c := range m.DataField {
		e.EncodeElement(c, sedataField)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DataFields) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_DataFields:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "dataField"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "dataField"}:
				tmp := NewCT_DataField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DataField = append(m.DataField, tmp)
			default:
				office.Log("skipping unsupported element on CT_DataFields %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DataFields
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DataFields and its children
func (m *CT_DataFields) Validate() error {
	return m.ValidateWithPath("CT_DataFields")
}

// ValidateWithPath validates the CT_DataFields and its children, prefixing error messages with path
func (m *CT_DataFields) ValidateWithPath(path string) error {
	for i, v := range m.DataField {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DataField[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
