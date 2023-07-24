package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_PivotHierarchy struct {
	// Outline New Levels
	OutlineAttr *bool
	// Multiple Field Filters
	MultipleItemSelectionAllowedAttr *bool
	// New Levels Subtotals At Top
	SubtotalTopAttr *bool
	// Show In Field List
	ShowInFieldListAttr *bool
	// Drag To Row
	DragToRowAttr *bool
	// Drag To Column
	DragToColAttr *bool
	// Drag to Page
	DragToPageAttr *bool
	// Drag To Data
	DragToDataAttr *bool
	// Drag Off
	DragOffAttr *bool
	// Inclusive Manual Filter
	IncludeNewItemsInFilterAttr *bool
	// Hierarchy Caption
	CaptionAttr *string
	// OLAP Member Properties
	Mps *CT_MemberProperties
	// Members
	Members []*CT_Members
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_PivotHierarchy() *CT_PivotHierarchy {
	ret := &CT_PivotHierarchy{}
	return ret
}

func (m *CT_PivotHierarchy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OutlineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "outline"},
			Value: fmt.Sprintf("%d", b2i(*m.OutlineAttr))})
	}
	if m.MultipleItemSelectionAllowedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "multipleItemSelectionAllowed"},
			Value: fmt.Sprintf("%d", b2i(*m.MultipleItemSelectionAllowedAttr))})
	}
	if m.SubtotalTopAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "subtotalTop"},
			Value: fmt.Sprintf("%d", b2i(*m.SubtotalTopAttr))})
	}
	if m.ShowInFieldListAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showInFieldList"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowInFieldListAttr))})
	}
	if m.DragToRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToRow"},
			Value: fmt.Sprintf("%d", b2i(*m.DragToRowAttr))})
	}
	if m.DragToColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToCol"},
			Value: fmt.Sprintf("%d", b2i(*m.DragToColAttr))})
	}
	if m.DragToPageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToPage"},
			Value: fmt.Sprintf("%d", b2i(*m.DragToPageAttr))})
	}
	if m.DragToDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToData"},
			Value: fmt.Sprintf("%d", b2i(*m.DragToDataAttr))})
	}
	if m.DragOffAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragOff"},
			Value: fmt.Sprintf("%d", b2i(*m.DragOffAttr))})
	}
	if m.IncludeNewItemsInFilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "includeNewItemsInFilter"},
			Value: fmt.Sprintf("%d", b2i(*m.IncludeNewItemsInFilterAttr))})
	}
	if m.CaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "caption"},
			Value: fmt.Sprintf("%v", *m.CaptionAttr)})
	}
	e.EncodeToken(start)
	if m.Mps != nil {
		semps := xml.StartElement{Name: xml.Name{Local: "ma:mps"}}
		e.EncodeElement(m.Mps, semps)
	}
	if m.Members != nil {
		semembers := xml.StartElement{Name: xml.Name{Local: "ma:members"}}
		for _, c := range m.Members {
			e.EncodeElement(c, semembers)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotHierarchy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "dragToData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToDataAttr = &parsed
			continue
		}
		if attr.Name.Local == "outline" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OutlineAttr = &parsed
			continue
		}
		if attr.Name.Local == "subtotalTop" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SubtotalTopAttr = &parsed
			continue
		}
		if attr.Name.Local == "showInFieldList" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowInFieldListAttr = &parsed
			continue
		}
		if attr.Name.Local == "dragToRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToRowAttr = &parsed
			continue
		}
		if attr.Name.Local == "dragToCol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToColAttr = &parsed
			continue
		}
		if attr.Name.Local == "dragToPage" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToPageAttr = &parsed
			continue
		}
		if attr.Name.Local == "multipleItemSelectionAllowed" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MultipleItemSelectionAllowedAttr = &parsed
			continue
		}
		if attr.Name.Local == "dragOff" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragOffAttr = &parsed
			continue
		}
		if attr.Name.Local == "includeNewItemsInFilter" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IncludeNewItemsInFilterAttr = &parsed
			continue
		}
		if attr.Name.Local == "caption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CaptionAttr = &parsed
			continue
		}
	}
lCT_PivotHierarchy:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "mps"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "mps"}:
				m.Mps = NewCT_MemberProperties()
				if err := d.DecodeElement(m.Mps, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "members"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "members"}:
				tmp := NewCT_Members()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Members = append(m.Members, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_PivotHierarchy %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotHierarchy
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotHierarchy and its children
func (m *CT_PivotHierarchy) Validate() error {
	return m.ValidateWithPath("CT_PivotHierarchy")
}

// ValidateWithPath validates the CT_PivotHierarchy and its children, prefixing error messages with path
func (m *CT_PivotHierarchy) ValidateWithPath(path string) error {
	if m.Mps != nil {
		if err := m.Mps.ValidateWithPath(path + "/Mps"); err != nil {
			return err
		}
	}
	for i, v := range m.Members {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Members[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
