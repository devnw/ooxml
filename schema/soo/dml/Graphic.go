package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type Graphic struct {
	CT_GraphicalObject
}

func NewGraphic() *Graphic {
	ret := &Graphic{}
	ret.CT_GraphicalObject = *NewCT_GraphicalObject()
	return ret
}

func (m *Graphic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_GraphicalObject.MarshalXML(e, start)
}

func (m *Graphic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_GraphicalObject = *NewCT_GraphicalObject()
lGraphic:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "graphicData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "graphicData"}:
				if err := d.DecodeElement(m.GraphicData, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on Graphic %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lGraphic
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Graphic and its children
func (m *Graphic) Validate() error {
	return m.ValidateWithPath("Graphic")
}

// ValidateWithPath validates the Graphic and its children, prefixing error messages with path
func (m *Graphic) ValidateWithPath(path string) error {
	if err := m.CT_GraphicalObject.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
