package custom_properties

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type Properties struct {
	CT_Properties
}

func NewProperties() *Properties {
	ret := &Properties{}
	ret.CT_Properties = *NewCT_Properties()
	return ret
}

func (m *Properties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/custom-properties"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:vt"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "Properties"
	return m.CT_Properties.MarshalXML(e, start)
}

func (m *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Properties = *NewCT_Properties()
lProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", Local: "property"}:
				tmp := NewCT_Property()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Property = append(m.Property, tmp)
			default:
				office.Log("skipping unsupported element on Properties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Properties and its children
func (m *Properties) Validate() error {
	return m.ValidateWithPath("Properties")
}

// ValidateWithPath validates the Properties and its children, prefixing error messages with path
func (m *Properties) ValidateWithPath(path string) error {
	if err := m.CT_Properties.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
