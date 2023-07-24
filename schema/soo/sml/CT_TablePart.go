package sml

import (
	"encoding/xml"
	"fmt"
)

type CT_TablePart struct {
	IdAttr string
}

func NewCT_TablePart() *CT_TablePart {
	ret := &CT_TablePart{}
	return ret
}

func (m *CT_TablePart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TablePart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TablePart: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TablePart and its children
func (m *CT_TablePart) Validate() error {
	return m.ValidateWithPath("CT_TablePart")
}

// ValidateWithPath validates the CT_TablePart and its children, prefixing error messages with path
func (m *CT_TablePart) ValidateWithPath(path string) error {
	return nil
}
