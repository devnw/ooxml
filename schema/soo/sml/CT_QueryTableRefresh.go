package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_QueryTableRefresh struct {
	// Preserve Sort & Filter Layout
	PreserveSortFilterLayoutAttr *bool
	// Next Field Id Wrapped
	FieldIdWrappedAttr *bool
	// Headers In Last Refresh
	HeadersInLastRefreshAttr *bool
	// Minimum Refresh Version
	MinimumVersionAttr *uint8
	// Next field id
	NextIdAttr *uint32
	// Columns Left
	UnboundColumnsLeftAttr *uint32
	// Columns Right
	UnboundColumnsRightAttr *uint32
	// Query table fields
	QueryTableFields *CT_QueryTableFields
	// Deleted Fields
	QueryTableDeletedFields *CT_QueryTableDeletedFields
	// Sort State
	SortState *CT_SortState
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_QueryTableRefresh() *CT_QueryTableRefresh {
	ret := &CT_QueryTableRefresh{}
	ret.QueryTableFields = NewCT_QueryTableFields()
	return ret
}

func (m *CT_QueryTableRefresh) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PreserveSortFilterLayoutAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preserveSortFilterLayout"},
			Value: fmt.Sprintf("%d", b2i(*m.PreserveSortFilterLayoutAttr))})
	}
	if m.FieldIdWrappedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fieldIdWrapped"},
			Value: fmt.Sprintf("%d", b2i(*m.FieldIdWrappedAttr))})
	}
	if m.HeadersInLastRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headersInLastRefresh"},
			Value: fmt.Sprintf("%d", b2i(*m.HeadersInLastRefreshAttr))})
	}
	if m.MinimumVersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minimumVersion"},
			Value: fmt.Sprintf("%v", *m.MinimumVersionAttr)})
	}
	if m.NextIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "nextId"},
			Value: fmt.Sprintf("%v", *m.NextIdAttr)})
	}
	if m.UnboundColumnsLeftAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "unboundColumnsLeft"},
			Value: fmt.Sprintf("%v", *m.UnboundColumnsLeftAttr)})
	}
	if m.UnboundColumnsRightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "unboundColumnsRight"},
			Value: fmt.Sprintf("%v", *m.UnboundColumnsRightAttr)})
	}
	e.EncodeToken(start)
	sequeryTableFields := xml.StartElement{Name: xml.Name{Local: "ma:queryTableFields"}}
	e.EncodeElement(m.QueryTableFields, sequeryTableFields)
	if m.QueryTableDeletedFields != nil {
		sequeryTableDeletedFields := xml.StartElement{Name: xml.Name{Local: "ma:queryTableDeletedFields"}}
		e.EncodeElement(m.QueryTableDeletedFields, sequeryTableDeletedFields)
	}
	if m.SortState != nil {
		sesortState := xml.StartElement{Name: xml.Name{Local: "ma:sortState"}}
		e.EncodeElement(m.SortState, sesortState)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_QueryTableRefresh) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.QueryTableFields = NewCT_QueryTableFields()
	for _, attr := range start.Attr {
		if attr.Name.Local == "preserveSortFilterLayout" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveSortFilterLayoutAttr = &parsed
			continue
		}
		if attr.Name.Local == "fieldIdWrapped" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FieldIdWrappedAttr = &parsed
			continue
		}
		if attr.Name.Local == "headersInLastRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HeadersInLastRefreshAttr = &parsed
			continue
		}
		if attr.Name.Local == "minimumVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.MinimumVersionAttr = &pt
			continue
		}
		if attr.Name.Local == "nextId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NextIdAttr = &pt
			continue
		}
		if attr.Name.Local == "unboundColumnsLeft" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.UnboundColumnsLeftAttr = &pt
			continue
		}
		if attr.Name.Local == "unboundColumnsRight" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.UnboundColumnsRightAttr = &pt
			continue
		}
	}
lCT_QueryTableRefresh:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "queryTableFields"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "queryTableFields"}:
				if err := d.DecodeElement(m.QueryTableFields, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "queryTableDeletedFields"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "queryTableDeletedFields"}:
				m.QueryTableDeletedFields = NewCT_QueryTableDeletedFields()
				if err := d.DecodeElement(m.QueryTableDeletedFields, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sortState"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "sortState"}:
				m.SortState = NewCT_SortState()
				if err := d.DecodeElement(m.SortState, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_QueryTableRefresh %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_QueryTableRefresh
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_QueryTableRefresh and its children
func (m *CT_QueryTableRefresh) Validate() error {
	return m.ValidateWithPath("CT_QueryTableRefresh")
}

// ValidateWithPath validates the CT_QueryTableRefresh and its children, prefixing error messages with path
func (m *CT_QueryTableRefresh) ValidateWithPath(path string) error {
	if err := m.QueryTableFields.ValidateWithPath(path + "/QueryTableFields"); err != nil {
		return err
	}
	if m.QueryTableDeletedFields != nil {
		if err := m.QueryTableDeletedFields.ValidateWithPath(path + "/QueryTableDeletedFields"); err != nil {
			return err
		}
	}
	if m.SortState != nil {
		if err := m.SortState.ValidateWithPath(path + "/SortState"); err != nil {
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
