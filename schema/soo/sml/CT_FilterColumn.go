package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_FilterColumn struct {
	// Filter Column Data
	ColIdAttr uint32
	// Hidden AutoFilter Button
	HiddenButtonAttr *bool
	// Show Filter Button
	ShowButtonAttr *bool
	// Filter Criteria
	Filters *CT_Filters
	// Top 10
	Top10 *CT_Top10
	// Custom Filters
	CustomFilters *CT_CustomFilters
	// Dynamic Filter
	DynamicFilter *CT_DynamicFilter
	// Color Filter Criteria
	ColorFilter *CT_ColorFilter
	// Icon Filter
	IconFilter *CT_IconFilter
	ExtLst     *CT_ExtensionList
}

func NewCT_FilterColumn() *CT_FilterColumn {
	ret := &CT_FilterColumn{}
	return ret
}

func (m *CT_FilterColumn) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "colId"},
		Value: fmt.Sprintf("%v", m.ColIdAttr)})
	if m.HiddenButtonAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hiddenButton"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenButtonAttr))})
	}
	if m.ShowButtonAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showButton"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowButtonAttr))})
	}
	e.EncodeToken(start)
	if m.Filters != nil {
		sefilters := xml.StartElement{Name: xml.Name{Local: "ma:filters"}}
		e.EncodeElement(m.Filters, sefilters)
	}
	if m.Top10 != nil {
		setop10 := xml.StartElement{Name: xml.Name{Local: "ma:top10"}}
		e.EncodeElement(m.Top10, setop10)
	}
	if m.CustomFilters != nil {
		secustomFilters := xml.StartElement{Name: xml.Name{Local: "ma:customFilters"}}
		e.EncodeElement(m.CustomFilters, secustomFilters)
	}
	if m.DynamicFilter != nil {
		sedynamicFilter := xml.StartElement{Name: xml.Name{Local: "ma:dynamicFilter"}}
		e.EncodeElement(m.DynamicFilter, sedynamicFilter)
	}
	if m.ColorFilter != nil {
		secolorFilter := xml.StartElement{Name: xml.Name{Local: "ma:colorFilter"}}
		e.EncodeElement(m.ColorFilter, secolorFilter)
	}
	if m.IconFilter != nil {
		seiconFilter := xml.StartElement{Name: xml.Name{Local: "ma:iconFilter"}}
		e.EncodeElement(m.IconFilter, seiconFilter)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FilterColumn) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "colId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ColIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "hiddenButton" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenButtonAttr = &parsed
			continue
		}
		if attr.Name.Local == "showButton" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowButtonAttr = &parsed
			continue
		}
	}
lCT_FilterColumn:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "filters"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "filters"}:
				m.Filters = NewCT_Filters()
				if err := d.DecodeElement(m.Filters, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "top10"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "top10"}:
				m.Top10 = NewCT_Top10()
				if err := d.DecodeElement(m.Top10, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "customFilters"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "customFilters"}:
				m.CustomFilters = NewCT_CustomFilters()
				if err := d.DecodeElement(m.CustomFilters, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "dynamicFilter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "dynamicFilter"}:
				m.DynamicFilter = NewCT_DynamicFilter()
				if err := d.DecodeElement(m.DynamicFilter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "colorFilter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "colorFilter"}:
				m.ColorFilter = NewCT_ColorFilter()
				if err := d.DecodeElement(m.ColorFilter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "iconFilter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "iconFilter"}:
				m.IconFilter = NewCT_IconFilter()
				if err := d.DecodeElement(m.IconFilter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_FilterColumn %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FilterColumn
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FilterColumn and its children
func (m *CT_FilterColumn) Validate() error {
	return m.ValidateWithPath("CT_FilterColumn")
}

// ValidateWithPath validates the CT_FilterColumn and its children, prefixing error messages with path
func (m *CT_FilterColumn) ValidateWithPath(path string) error {
	if m.Filters != nil {
		if err := m.Filters.ValidateWithPath(path + "/Filters"); err != nil {
			return err
		}
	}
	if m.Top10 != nil {
		if err := m.Top10.ValidateWithPath(path + "/Top10"); err != nil {
			return err
		}
	}
	if m.CustomFilters != nil {
		if err := m.CustomFilters.ValidateWithPath(path + "/CustomFilters"); err != nil {
			return err
		}
	}
	if m.DynamicFilter != nil {
		if err := m.DynamicFilter.ValidateWithPath(path + "/DynamicFilter"); err != nil {
			return err
		}
	}
	if m.ColorFilter != nil {
		if err := m.ColorFilter.ValidateWithPath(path + "/ColorFilter"); err != nil {
			return err
		}
	}
	if m.IconFilter != nil {
		if err := m.IconFilter.ValidateWithPath(path + "/IconFilter"); err != nil {
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
