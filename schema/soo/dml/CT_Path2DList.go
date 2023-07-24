package dml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_Path2DList struct {
	Path []*CT_Path2D
}

func NewCT_Path2DList() *CT_Path2DList {
	ret := &CT_Path2DList{}
	return ret
}

func (m *CT_Path2DList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Path != nil {
		sepath := xml.StartElement{Name: xml.Name{Local: "a:path"}}
		for _, c := range m.Path {
			e.EncodeElement(c, sepath)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Path2DList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Path2DList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "path"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "path"}:
				tmp := NewCT_Path2D()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Path = append(m.Path, tmp)
			default:
				office.Log("skipping unsupported element on CT_Path2DList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Path2DList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Path2DList and its children
func (m *CT_Path2DList) Validate() error {
	return m.ValidateWithPath("CT_Path2DList")
}

// ValidateWithPath validates the CT_Path2DList and its children, prefixing error messages with path
func (m *CT_Path2DList) ValidateWithPath(path string) error {
	for i, v := range m.Path {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Path[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
