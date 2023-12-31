package word

import (
	"encoding/xml"
	"fmt"
)

type Anchorlock struct {
	CT_AnchorLock
}

func NewAnchorlock() *Anchorlock {
	ret := &Anchorlock{}
	ret.CT_AnchorLock = *NewCT_AnchorLock()
	return ret
}

func (m *Anchorlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:word"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "anchorlock"
	return m.CT_AnchorLock.MarshalXML(e, start)
}

func (m *Anchorlock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_AnchorLock = *NewCT_AnchorLock()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Anchorlock: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Anchorlock and its children
func (m *Anchorlock) Validate() error {
	return m.ValidateWithPath("Anchorlock")
}

// ValidateWithPath validates the Anchorlock and its children, prefixing error messages with path
func (m *Anchorlock) ValidateWithPath(path string) error {
	if err := m.CT_AnchorLock.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
