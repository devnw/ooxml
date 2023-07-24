package sml

import (
	"encoding/xml"
	"strconv"

	"go.devnw.com/ooxml"
)

type PivotCacheRecords struct {
	CT_PivotCacheRecords
}

func NewPivotCacheRecords() *PivotCacheRecords {
	ret := &PivotCacheRecords{}
	ret.CT_PivotCacheRecords = *NewCT_PivotCacheRecords()
	return ret
}

func (m *PivotCacheRecords) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:pivotCacheRecords"
	return m.CT_PivotCacheRecords.MarshalXML(e, start)
}

func (m *PivotCacheRecords) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_PivotCacheRecords = *NewCT_PivotCacheRecords()
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lPivotCacheRecords:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "r"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "r"}:
				tmp := NewCT_Record()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.R = append(m.R, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on PivotCacheRecords %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lPivotCacheRecords
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the PivotCacheRecords and its children
func (m *PivotCacheRecords) Validate() error {
	return m.ValidateWithPath("PivotCacheRecords")
}

// ValidateWithPath validates the PivotCacheRecords and its children, prefixing error messages with path
func (m *PivotCacheRecords) ValidateWithPath(path string) error {
	if err := m.CT_PivotCacheRecords.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
