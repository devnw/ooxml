package relationships

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_Relationships struct {
	Relationship []*Relationship
}

func NewCT_Relationships() *CT_Relationships {
	ret := &CT_Relationships{}
	return ret
}

func (m *CT_Relationships) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Relationship != nil {
		seRelationship := xml.StartElement{Name: xml.Name{Local: "Relationship"}}
		for _, c := range m.Relationship {
			e.EncodeElement(c, seRelationship)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Relationships) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Relationships:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/relationships", Local: "Relationship"}:
				tmp := NewRelationship()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Relationship = append(m.Relationship, tmp)
			default:
				office.Log("skipping unsupported element on CT_Relationships %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Relationships
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Relationships and its children
func (m *CT_Relationships) Validate() error {
	return m.ValidateWithPath("CT_Relationships")
}

// ValidateWithPath validates the CT_Relationships and its children, prefixing error messages with path
func (m *CT_Relationships) ValidateWithPath(path string) error {
	for i, v := range m.Relationship {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Relationship[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
