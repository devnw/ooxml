package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcCT_ClipPath struct {
	VAttr string
}

func NewOfcCT_ClipPath() *OfcCT_ClipPath {
	ret := &OfcCT_ClipPath{}
	return ret
}

func (m *OfcCT_ClipPath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "v"},
		Value: fmt.Sprintf("%v", m.VAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_ClipPath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "v" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_ClipPath: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_ClipPath and its children
func (m *OfcCT_ClipPath) Validate() error {
	return m.ValidateWithPath("OfcCT_ClipPath")
}

// ValidateWithPath validates the OfcCT_ClipPath and its children, prefixing error messages with path
func (m *OfcCT_ClipPath) ValidateWithPath(path string) error {
	return nil
}
