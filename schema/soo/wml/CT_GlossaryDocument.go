package wml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_GlossaryDocument struct {
	// Document Background
	Background *CT_Background
	DocParts   *CT_DocParts
}

func NewCT_GlossaryDocument() *CT_GlossaryDocument {
	ret := &CT_GlossaryDocument{}
	return ret
}

func (m *CT_GlossaryDocument) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Background != nil {
		sebackground := xml.StartElement{Name: xml.Name{Local: "w:background"}}
		e.EncodeElement(m.Background, sebackground)
	}
	if m.DocParts != nil {
		sedocParts := xml.StartElement{Name: xml.Name{Local: "w:docParts"}}
		e.EncodeElement(m.DocParts, sedocParts)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GlossaryDocument) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GlossaryDocument:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "background"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "background"}:
				m.Background = NewCT_Background()
				if err := d.DecodeElement(m.Background, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docParts"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docParts"}:
				m.DocParts = NewCT_DocParts()
				if err := d.DecodeElement(m.DocParts, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_GlossaryDocument %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GlossaryDocument
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GlossaryDocument and its children
func (m *CT_GlossaryDocument) Validate() error {
	return m.ValidateWithPath("CT_GlossaryDocument")
}

// ValidateWithPath validates the CT_GlossaryDocument and its children, prefixing error messages with path
func (m *CT_GlossaryDocument) ValidateWithPath(path string) error {
	if m.Background != nil {
		if err := m.Background.ValidateWithPath(path + "/Background"); err != nil {
			return err
		}
	}
	if m.DocParts != nil {
		if err := m.DocParts.ValidateWithPath(path + "/DocParts"); err != nil {
			return err
		}
	}
	return nil
}
