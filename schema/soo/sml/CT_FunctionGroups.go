package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_FunctionGroups struct {
	// Built-in Function Group Count
	BuiltInGroupCountAttr *uint32
	// Function Group
	FunctionGroup []*CT_FunctionGroup
}

func NewCT_FunctionGroups() *CT_FunctionGroups {
	ret := &CT_FunctionGroups{}
	return ret
}

func (m *CT_FunctionGroups) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BuiltInGroupCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "builtInGroupCount"},
			Value: fmt.Sprintf("%v", *m.BuiltInGroupCountAttr)})
	}
	e.EncodeToken(start)
	if m.FunctionGroup != nil {
		sefunctionGroup := xml.StartElement{Name: xml.Name{Local: "ma:functionGroup"}}
		for _, c := range m.FunctionGroup {
			e.EncodeElement(c, sefunctionGroup)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FunctionGroups) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "builtInGroupCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.BuiltInGroupCountAttr = &pt
			continue
		}
	}
lCT_FunctionGroups:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "functionGroup"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "functionGroup"}:
				tmp := NewCT_FunctionGroup()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FunctionGroup = append(m.FunctionGroup, tmp)
			default:
				office.Log("skipping unsupported element on CT_FunctionGroups %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FunctionGroups
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FunctionGroups and its children
func (m *CT_FunctionGroups) Validate() error {
	return m.ValidateWithPath("CT_FunctionGroups")
}

// ValidateWithPath validates the CT_FunctionGroups and its children, prefixing error messages with path
func (m *CT_FunctionGroups) ValidateWithPath(path string) error {
	for i, v := range m.FunctionGroup {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FunctionGroup[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
