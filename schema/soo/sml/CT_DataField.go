package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_DataField struct {
	// Data Field Name
	NameAttr *string
	// Field
	FldAttr uint32
	// Subtotal
	SubtotalAttr ST_DataConsolidateFunction
	// Show Data As Display Format
	ShowDataAsAttr ST_ShowDataAs
	// 'Show Data As' Base Field
	BaseFieldAttr *int32
	// 'Show Data As' Base Setting
	BaseItemAttr *uint32
	// Number Format Id
	NumFmtIdAttr *uint32
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_DataField() *CT_DataField {
	ret := &CT_DataField{}
	return ret
}

func (m *CT_DataField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fld"},
		Value: fmt.Sprintf("%v", m.FldAttr)})
	if m.SubtotalAttr != ST_DataConsolidateFunctionUnset {
		attr, err := m.SubtotalAttr.MarshalXMLAttr(xml.Name{Local: "subtotal"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ShowDataAsAttr != ST_ShowDataAsUnset {
		attr, err := m.ShowDataAsAttr.MarshalXMLAttr(xml.Name{Local: "showDataAs"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BaseFieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "baseField"},
			Value: fmt.Sprintf("%v", *m.BaseFieldAttr)})
	}
	if m.BaseItemAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "baseItem"},
			Value: fmt.Sprintf("%v", *m.BaseItemAttr)})
	}
	if m.NumFmtIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numFmtId"},
			Value: fmt.Sprintf("%v", *m.NumFmtIdAttr)})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DataField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
		if attr.Name.Local == "fld" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.FldAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "subtotal" {
			m.SubtotalAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "showDataAs" {
			m.ShowDataAsAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "baseField" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.BaseFieldAttr = &pt
			continue
		}
		if attr.Name.Local == "baseItem" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.BaseItemAttr = &pt
			continue
		}
		if attr.Name.Local == "numFmtId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NumFmtIdAttr = &pt
			continue
		}
	}
lCT_DataField:
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
				office.Log("skipping unsupported element on CT_DataField %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DataField
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DataField and its children
func (m *CT_DataField) Validate() error {
	return m.ValidateWithPath("CT_DataField")
}

// ValidateWithPath validates the CT_DataField and its children, prefixing error messages with path
func (m *CT_DataField) ValidateWithPath(path string) error {
	if err := m.SubtotalAttr.ValidateWithPath(path + "/SubtotalAttr"); err != nil {
		return err
	}
	if err := m.ShowDataAsAttr.ValidateWithPath(path + "/ShowDataAsAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
