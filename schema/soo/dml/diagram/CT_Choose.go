package diagram

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_Choose struct {
	NameAttr *string
	If       []*CT_When
	Else     *CT_Otherwise
}

func NewCT_Choose() *CT_Choose {
	ret := &CT_Choose{}
	return ret
}

func (m *CT_Choose) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	e.EncodeToken(start)
	seif := xml.StartElement{Name: xml.Name{Local: "if"}}
	for _, c := range m.If {
		e.EncodeElement(c, seif)
	}
	if m.Else != nil {
		seelse := xml.StartElement{Name: xml.Name{Local: "else"}}
		e.EncodeElement(m.Else, seelse)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Choose) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
	}
lCT_Choose:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "if"}:
				tmp := NewCT_When()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.If = append(m.If, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "else"}:
				m.Else = NewCT_Otherwise()
				if err := d.DecodeElement(m.Else, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_Choose %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Choose
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Choose and its children
func (m *CT_Choose) Validate() error {
	return m.ValidateWithPath("CT_Choose")
}

// ValidateWithPath validates the CT_Choose and its children, prefixing error messages with path
func (m *CT_Choose) ValidateWithPath(path string) error {
	for i, v := range m.If {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/If[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.Else != nil {
		if err := m.Else.ValidateWithPath(path + "/Else"); err != nil {
			return err
		}
	}
	return nil
}
