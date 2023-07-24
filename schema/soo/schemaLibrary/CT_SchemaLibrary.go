package schemaLibrary

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_SchemaLibrary struct {
	Schema []*CT_Schema
}

func NewCT_SchemaLibrary() *CT_SchemaLibrary {
	ret := &CT_SchemaLibrary{}
	return ret
}

func (m *CT_SchemaLibrary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Schema != nil {
		seschema := xml.StartElement{Name: xml.Name{Local: "ma:schema"}}
		for _, c := range m.Schema {
			e.EncodeElement(c, seschema)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SchemaLibrary) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SchemaLibrary:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/schemaLibrary/2006/main", Local: "schema"}:
				tmp := NewCT_Schema()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Schema = append(m.Schema, tmp)
			default:
				office.Log("skipping unsupported element on CT_SchemaLibrary %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SchemaLibrary
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SchemaLibrary and its children
func (m *CT_SchemaLibrary) Validate() error {
	return m.ValidateWithPath("CT_SchemaLibrary")
}

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (m *CT_SchemaLibrary) ValidateWithPath(path string) error {
	for i, v := range m.Schema {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Schema[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
