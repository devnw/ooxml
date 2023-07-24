package dml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_GradientStopList struct {
	Gs []*CT_GradientStop
}

func NewCT_GradientStopList() *CT_GradientStopList {
	ret := &CT_GradientStopList{}
	return ret
}

func (m *CT_GradientStopList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	segs := xml.StartElement{Name: xml.Name{Local: "a:gs"}}
	for _, c := range m.Gs {
		e.EncodeElement(c, segs)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GradientStopList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GradientStopList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "gs"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "gs"}:
				tmp := NewCT_GradientStop()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Gs = append(m.Gs, tmp)
			default:
				office.Log("skipping unsupported element on CT_GradientStopList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GradientStopList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GradientStopList and its children
func (m *CT_GradientStopList) Validate() error {
	return m.ValidateWithPath("CT_GradientStopList")
}

// ValidateWithPath validates the CT_GradientStopList and its children, prefixing error messages with path
func (m *CT_GradientStopList) ValidateWithPath(path string) error {
	for i, v := range m.Gs {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Gs[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
