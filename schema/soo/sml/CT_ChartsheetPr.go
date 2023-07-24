package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_ChartsheetPr struct {
	// Published
	PublishedAttr *bool
	// Code Name
	CodeNameAttr *string
	TabColor     *CT_Color
}

func NewCT_ChartsheetPr() *CT_ChartsheetPr {
	ret := &CT_ChartsheetPr{}
	return ret
}

func (m *CT_ChartsheetPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PublishedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "published"},
			Value: fmt.Sprintf("%d", b2i(*m.PublishedAttr))})
	}
	if m.CodeNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "codeName"},
			Value: fmt.Sprintf("%v", *m.CodeNameAttr)})
	}
	e.EncodeToken(start)
	if m.TabColor != nil {
		setabColor := xml.StartElement{Name: xml.Name{Local: "ma:tabColor"}}
		e.EncodeElement(m.TabColor, setabColor)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ChartsheetPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "published" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PublishedAttr = &parsed
			continue
		}
		if attr.Name.Local == "codeName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CodeNameAttr = &parsed
			continue
		}
	}
lCT_ChartsheetPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tabColor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "tabColor"}:
				m.TabColor = NewCT_Color()
				if err := d.DecodeElement(m.TabColor, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_ChartsheetPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ChartsheetPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ChartsheetPr and its children
func (m *CT_ChartsheetPr) Validate() error {
	return m.ValidateWithPath("CT_ChartsheetPr")
}

// ValidateWithPath validates the CT_ChartsheetPr and its children, prefixing error messages with path
func (m *CT_ChartsheetPr) ValidateWithPath(path string) error {
	if m.TabColor != nil {
		if err := m.TabColor.ValidateWithPath(path + "/TabColor"); err != nil {
			return err
		}
	}
	return nil
}
