package sml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_ProtectedRanges struct {
	// Protected Range
	ProtectedRange []*CT_ProtectedRange
}

func NewCT_ProtectedRanges() *CT_ProtectedRanges {
	ret := &CT_ProtectedRanges{}
	return ret
}

func (m *CT_ProtectedRanges) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seprotectedRange := xml.StartElement{Name: xml.Name{Local: "ma:protectedRange"}}
	for _, c := range m.ProtectedRange {
		e.EncodeElement(c, seprotectedRange)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ProtectedRanges) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ProtectedRanges:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "protectedRange"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "protectedRange"}:
				tmp := NewCT_ProtectedRange()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ProtectedRange = append(m.ProtectedRange, tmp)
			default:
				office.Log("skipping unsupported element on CT_ProtectedRanges %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ProtectedRanges
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ProtectedRanges and its children
func (m *CT_ProtectedRanges) Validate() error {
	return m.ValidateWithPath("CT_ProtectedRanges")
}

// ValidateWithPath validates the CT_ProtectedRanges and its children, prefixing error messages with path
func (m *CT_ProtectedRanges) ValidateWithPath(path string) error {
	for i, v := range m.ProtectedRange {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ProtectedRange[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
