package sml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_Sheets struct {
	// Sheet Information
	Sheet []*CT_Sheet
}

func NewCT_Sheets() *CT_Sheets {
	ret := &CT_Sheets{}
	return ret
}

func (m *CT_Sheets) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sesheet := xml.StartElement{Name: xml.Name{Local: "ma:sheet"}}
	for _, c := range m.Sheet {
		e.EncodeElement(c, sesheet)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Sheets) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Sheets:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheet"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "sheet"}:
				tmp := NewCT_Sheet()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Sheet = append(m.Sheet, tmp)
			default:
				office.Log("skipping unsupported element on CT_Sheets %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Sheets
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Sheets and its children
func (m *CT_Sheets) Validate() error {
	return m.ValidateWithPath("CT_Sheets")
}

// ValidateWithPath validates the CT_Sheets and its children, prefixing error messages with path
func (m *CT_Sheets) ValidateWithPath(path string) error {
	for i, v := range m.Sheet {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Sheet[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
