package wml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_Placeholder struct {
	// Document Part Reference
	DocPart *CT_String
}

func NewCT_Placeholder() *CT_Placeholder {
	ret := &CT_Placeholder{}
	ret.DocPart = NewCT_String()
	return ret
}

func (m *CT_Placeholder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sedocPart := xml.StartElement{Name: xml.Name{Local: "w:docPart"}}
	e.EncodeElement(m.DocPart, sedocPart)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Placeholder) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.DocPart = NewCT_String()
lCT_Placeholder:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docPart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docPart"}:
				if err := d.DecodeElement(m.DocPart, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_Placeholder %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Placeholder
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Placeholder and its children
func (m *CT_Placeholder) Validate() error {
	return m.ValidateWithPath("CT_Placeholder")
}

// ValidateWithPath validates the CT_Placeholder and its children, prefixing error messages with path
func (m *CT_Placeholder) ValidateWithPath(path string) error {
	if err := m.DocPart.ValidateWithPath(path + "/DocPart"); err != nil {
		return err
	}
	return nil
}
