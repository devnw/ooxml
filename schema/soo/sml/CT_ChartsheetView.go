package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_ChartsheetView struct {
	// Sheet Tab Selected
	TabSelectedAttr *bool
	// Window Zoom Scale
	ZoomScaleAttr *uint32
	// Workbook View Id
	WorkbookViewIdAttr uint32
	// Zoom To Fit
	ZoomToFitAttr *bool
	ExtLst        *CT_ExtensionList
}

func NewCT_ChartsheetView() *CT_ChartsheetView {
	ret := &CT_ChartsheetView{}
	return ret
}

func (m *CT_ChartsheetView) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TabSelectedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tabSelected"},
			Value: fmt.Sprintf("%d", b2i(*m.TabSelectedAttr))})
	}
	if m.ZoomScaleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoomScale"},
			Value: fmt.Sprintf("%v", *m.ZoomScaleAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "workbookViewId"},
		Value: fmt.Sprintf("%v", m.WorkbookViewIdAttr)})
	if m.ZoomToFitAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoomToFit"},
			Value: fmt.Sprintf("%d", b2i(*m.ZoomToFitAttr))})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ChartsheetView) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "tabSelected" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TabSelectedAttr = &parsed
			continue
		}
		if attr.Name.Local == "zoomScale" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ZoomScaleAttr = &pt
			continue
		}
		if attr.Name.Local == "workbookViewId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.WorkbookViewIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "zoomToFit" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ZoomToFitAttr = &parsed
			continue
		}
	}
lCT_ChartsheetView:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_ChartsheetView %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ChartsheetView
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ChartsheetView and its children
func (m *CT_ChartsheetView) Validate() error {
	return m.ValidateWithPath("CT_ChartsheetView")
}

// ValidateWithPath validates the CT_ChartsheetView and its children, prefixing error messages with path
func (m *CT_ChartsheetView) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
