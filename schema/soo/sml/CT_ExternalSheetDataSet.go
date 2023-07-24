package sml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_ExternalSheetDataSet struct {
	// External Sheet Data Set
	SheetData []*CT_ExternalSheetData
}

func NewCT_ExternalSheetDataSet() *CT_ExternalSheetDataSet {
	ret := &CT_ExternalSheetDataSet{}
	return ret
}

func (m *CT_ExternalSheetDataSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sesheetData := xml.StartElement{Name: xml.Name{Local: "ma:sheetData"}}
	for _, c := range m.SheetData {
		e.EncodeElement(c, sesheetData)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ExternalSheetDataSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ExternalSheetDataSet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "sheetData"}:
				tmp := NewCT_ExternalSheetData()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SheetData = append(m.SheetData, tmp)
			default:
				office.Log("skipping unsupported element on CT_ExternalSheetDataSet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalSheetDataSet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ExternalSheetDataSet and its children
func (m *CT_ExternalSheetDataSet) Validate() error {
	return m.ValidateWithPath("CT_ExternalSheetDataSet")
}

// ValidateWithPath validates the CT_ExternalSheetDataSet and its children, prefixing error messages with path
func (m *CT_ExternalSheetDataSet) ValidateWithPath(path string) error {
	for i, v := range m.SheetData {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SheetData[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
