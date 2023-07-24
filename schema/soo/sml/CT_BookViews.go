package sml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_BookViews struct {
	// Workbook View
	WorkbookView []*CT_BookView
}

func NewCT_BookViews() *CT_BookViews {
	ret := &CT_BookViews{}
	return ret
}

func (m *CT_BookViews) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seworkbookView := xml.StartElement{Name: xml.Name{Local: "ma:workbookView"}}
	for _, c := range m.WorkbookView {
		e.EncodeElement(c, seworkbookView)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BookViews) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BookViews:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "workbookView"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "workbookView"}:
				tmp := NewCT_BookView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.WorkbookView = append(m.WorkbookView, tmp)
			default:
				office.Log("skipping unsupported element on CT_BookViews %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BookViews
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BookViews and its children
func (m *CT_BookViews) Validate() error {
	return m.ValidateWithPath("CT_BookViews")
}

// ValidateWithPath validates the CT_BookViews and its children, prefixing error messages with path
func (m *CT_BookViews) ValidateWithPath(path string) error {
	for i, v := range m.WorkbookView {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/WorkbookView[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
