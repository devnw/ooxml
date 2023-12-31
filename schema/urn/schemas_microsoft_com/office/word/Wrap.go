package word

import (
	"encoding/xml"
	"fmt"
)

type Wrap struct {
	CT_Wrap
}

func NewWrap() *Wrap {
	ret := &Wrap{}
	ret.CT_Wrap = *NewCT_Wrap()
	return ret
}

func (m *Wrap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:word"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "wrap"
	return m.CT_Wrap.MarshalXML(e, start)
}

func (m *Wrap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Wrap = *NewCT_Wrap()
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "side" {
			m.SideAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "anchorx" {
			m.AnchorxAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "anchory" {
			m.AnchoryAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Wrap: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Wrap and its children
func (m *Wrap) Validate() error {
	return m.ValidateWithPath("Wrap")
}

// ValidateWithPath validates the Wrap and its children, prefixing error messages with path
func (m *Wrap) ValidateWithPath(path string) error {
	if err := m.CT_Wrap.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
