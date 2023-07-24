package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_PageSetUpPr struct {
	// Show Auto Page Breaks
	AutoPageBreaksAttr *bool
	// Fit To Page
	FitToPageAttr *bool
}

func NewCT_PageSetUpPr() *CT_PageSetUpPr {
	ret := &CT_PageSetUpPr{}
	return ret
}

func (m *CT_PageSetUpPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AutoPageBreaksAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoPageBreaks"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoPageBreaksAttr))})
	}
	if m.FitToPageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fitToPage"},
			Value: fmt.Sprintf("%d", b2i(*m.FitToPageAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PageSetUpPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "autoPageBreaks" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoPageBreaksAttr = &parsed
			continue
		}
		if attr.Name.Local == "fitToPage" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FitToPageAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PageSetUpPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PageSetUpPr and its children
func (m *CT_PageSetUpPr) Validate() error {
	return m.ValidateWithPath("CT_PageSetUpPr")
}

// ValidateWithPath validates the CT_PageSetUpPr and its children, prefixing error messages with path
func (m *CT_PageSetUpPr) ValidateWithPath(path string) error {
	return nil
}
