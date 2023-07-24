package wml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_FontsList struct {
	// Properties for a Single Font
	Font []*CT_Font
}

func NewCT_FontsList() *CT_FontsList {
	ret := &CT_FontsList{}
	return ret
}

func (m *CT_FontsList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Font != nil {
		sefont := xml.StartElement{Name: xml.Name{Local: "w:font"}}
		for _, c := range m.Font {
			e.EncodeElement(c, sefont)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FontsList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FontsList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "font"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "font"}:
				tmp := NewCT_Font()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Font = append(m.Font, tmp)
			default:
				office.Log("skipping unsupported element on CT_FontsList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FontsList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FontsList and its children
func (m *CT_FontsList) Validate() error {
	return m.ValidateWithPath("CT_FontsList")
}

// ValidateWithPath validates the CT_FontsList and its children, prefixing error messages with path
func (m *CT_FontsList) ValidateWithPath(path string) error {
	for i, v := range m.Font {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Font[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
