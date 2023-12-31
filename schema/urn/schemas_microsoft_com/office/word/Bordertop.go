package word

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type Bordertop struct {
	CT_Border
}

func NewBordertop() *Bordertop {
	ret := &Bordertop{}
	ret.CT_Border = *NewCT_Border()
	return ret
}

func (m *Bordertop) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:word"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "bordertop"
	return m.CT_Border.MarshalXML(e, start)
}

func (m *Bordertop) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Border = *NewCT_Border()
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "width" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.WidthAttr = &pt
			continue
		}
		if attr.Name.Local == "shadow" {
			m.ShadowAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Bordertop: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Bordertop and its children
func (m *Bordertop) Validate() error {
	return m.ValidateWithPath("Bordertop")
}

// ValidateWithPath validates the Bordertop and its children, prefixing error messages with path
func (m *Bordertop) ValidateWithPath(path string) error {
	if err := m.CT_Border.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
