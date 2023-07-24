package sml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_ExternalDefinedNames struct {
	// Defined Name
	DefinedName []*CT_ExternalDefinedName
}

func NewCT_ExternalDefinedNames() *CT_ExternalDefinedNames {
	ret := &CT_ExternalDefinedNames{}
	return ret
}

func (m *CT_ExternalDefinedNames) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DefinedName != nil {
		sedefinedName := xml.StartElement{Name: xml.Name{Local: "ma:definedName"}}
		for _, c := range m.DefinedName {
			e.EncodeElement(c, sedefinedName)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ExternalDefinedNames) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ExternalDefinedNames:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "definedName"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "definedName"}:
				tmp := NewCT_ExternalDefinedName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DefinedName = append(m.DefinedName, tmp)
			default:
				office.Log("skipping unsupported element on CT_ExternalDefinedNames %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalDefinedNames
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ExternalDefinedNames and its children
func (m *CT_ExternalDefinedNames) Validate() error {
	return m.ValidateWithPath("CT_ExternalDefinedNames")
}

// ValidateWithPath validates the CT_ExternalDefinedNames and its children, prefixing error messages with path
func (m *CT_ExternalDefinedNames) ValidateWithPath(path string) error {
	for i, v := range m.DefinedName {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DefinedName[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
