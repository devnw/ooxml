package word

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type Borderright struct {
	CT_Border
}

func NewBorderright() *Borderright {
	ret := &Borderright{}
	ret.CT_Border = *NewCT_Border()
	return ret
}

func (m *Borderright) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:word"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "borderright"
	return m.CT_Border.MarshalXML(e, start)
}

func (m *Borderright) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing Borderright: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Borderright and its children
func (m *Borderright) Validate() error {
	return m.ValidateWithPath("Borderright")
}

// ValidateWithPath validates the Borderright and its children, prefixing error messages with path
func (m *Borderright) ValidateWithPath(path string) error {
	if err := m.CT_Border.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
