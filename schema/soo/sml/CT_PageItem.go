package sml

import (
	"encoding/xml"
	"fmt"
)

type CT_PageItem struct {
	// Page Item Name
	NameAttr string
}

func NewCT_PageItem() *CT_PageItem {
	ret := &CT_PageItem{}
	return ret
}

func (m *CT_PageItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PageItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PageItem: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PageItem and its children
func (m *CT_PageItem) Validate() error {
	return m.ValidateWithPath("CT_PageItem")
}

// ValidateWithPath validates the CT_PageItem and its children, prefixing error messages with path
func (m *CT_PageItem) ValidateWithPath(path string) error {
	return nil
}
