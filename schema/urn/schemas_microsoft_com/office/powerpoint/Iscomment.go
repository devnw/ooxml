package powerpoint

import (
	"encoding/xml"
	"fmt"
)

type Iscomment struct {
	CT_Empty
}

func NewIscomment() *Iscomment {
	ret := &Iscomment{}
	ret.CT_Empty = *NewCT_Empty()
	return ret
}

func (m *Iscomment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:powerpoint"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "iscomment"
	return m.CT_Empty.MarshalXML(e, start)
}

func (m *Iscomment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Empty = *NewCT_Empty()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Iscomment: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Iscomment and its children
func (m *Iscomment) Validate() error {
	return m.ValidateWithPath("Iscomment")
}

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (m *Iscomment) ValidateWithPath(path string) error {
	if err := m.CT_Empty.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
