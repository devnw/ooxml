package sml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_ExtensionList struct {
	// Extension
	Ext []*CT_Extension
}

func NewCT_ExtensionList() *CT_ExtensionList {
	ret := &CT_ExtensionList{}
	return ret
}

func (m *CT_ExtensionList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Ext != nil {
		seext := xml.StartElement{Name: xml.Name{Local: "ma:ext"}}
		for _, c := range m.Ext {
			e.EncodeElement(c, seext)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ExtensionList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ExtensionList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "ext"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "ext"}:
				tmp := NewCT_Extension()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ext = append(m.Ext, tmp)
			default:
				office.Log("skipping unsupported element on CT_ExtensionList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExtensionList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ExtensionList and its children
func (m *CT_ExtensionList) Validate() error {
	return m.ValidateWithPath("CT_ExtensionList")
}

// ValidateWithPath validates the CT_ExtensionList and its children, prefixing error messages with path
func (m *CT_ExtensionList) ValidateWithPath(path string) error {
	for i, v := range m.Ext {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ext[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
