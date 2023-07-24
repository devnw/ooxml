package schemaLibrary

import (
	"encoding/xml"
	"fmt"
)

type CT_Schema struct {
	UriAttr              *string
	ManifestLocationAttr *string
	SchemaLocationAttr   *string
	SchemaLanguageAttr   *string
}

func NewCT_Schema() *CT_Schema {
	ret := &CT_Schema{}
	return ret
}

func (m *CT_Schema) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UriAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ma:uri"},
			Value: fmt.Sprintf("%v", *m.UriAttr)})
	}
	if m.ManifestLocationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ma:manifestLocation"},
			Value: fmt.Sprintf("%v", *m.ManifestLocationAttr)})
	}
	if m.SchemaLocationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ma:schemaLocation"},
			Value: fmt.Sprintf("%v", *m.SchemaLocationAttr)})
	}
	if m.SchemaLanguageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ma:schemaLanguage"},
			Value: fmt.Sprintf("%v", *m.SchemaLanguageAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Schema) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uri" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UriAttr = &parsed
			continue
		}
		if attr.Name.Local == "manifestLocation" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ManifestLocationAttr = &parsed
			continue
		}
		if attr.Name.Local == "schemaLocation" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SchemaLocationAttr = &parsed
			continue
		}
		if attr.Name.Local == "schemaLanguage" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SchemaLanguageAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Schema: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Schema and its children
func (m *CT_Schema) Validate() error {
	return m.ValidateWithPath("CT_Schema")
}

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (m *CT_Schema) ValidateWithPath(path string) error {
	return nil
}
