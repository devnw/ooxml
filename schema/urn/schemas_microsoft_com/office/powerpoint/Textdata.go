package powerpoint

import (
	"encoding/xml"
	"fmt"
)

type Textdata struct {
	CT_Rel
}

func NewTextdata() *Textdata {
	ret := &Textdata{}
	ret.CT_Rel = *NewCT_Rel()
	return ret
}

func (m *Textdata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:powerpoint"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "textdata"
	return m.CT_Rel.MarshalXML(e, start)
}

func (m *Textdata) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Rel = *NewCT_Rel()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Textdata: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Textdata and its children
func (m *Textdata) Validate() error {
	return m.ValidateWithPath("Textdata")
}

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (m *Textdata) ValidateWithPath(path string) error {
	if err := m.CT_Rel.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
