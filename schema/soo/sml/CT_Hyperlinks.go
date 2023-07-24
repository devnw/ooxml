package sml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_Hyperlinks struct {
	// Hyperlink
	Hyperlink []*CT_Hyperlink
}

func NewCT_Hyperlinks() *CT_Hyperlinks {
	ret := &CT_Hyperlinks{}
	return ret
}

func (m *CT_Hyperlinks) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sehyperlink := xml.StartElement{Name: xml.Name{Local: "ma:hyperlink"}}
	for _, c := range m.Hyperlink {
		e.EncodeElement(c, sehyperlink)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Hyperlinks) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Hyperlinks:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "hyperlink"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "hyperlink"}:
				tmp := NewCT_Hyperlink()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Hyperlink = append(m.Hyperlink, tmp)
			default:
				office.Log("skipping unsupported element on CT_Hyperlinks %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Hyperlinks
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Hyperlinks and its children
func (m *CT_Hyperlinks) Validate() error {
	return m.ValidateWithPath("CT_Hyperlinks")
}

// ValidateWithPath validates the CT_Hyperlinks and its children, prefixing error messages with path
func (m *CT_Hyperlinks) ValidateWithPath(path string) error {
	for i, v := range m.Hyperlink {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Hyperlink[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
