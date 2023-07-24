package wml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_SdtComboBox struct {
	// Combo Box Last Saved Value
	LastValueAttr *string
	// Combo Box List Item
	ListItem []*CT_SdtListItem
}

func NewCT_SdtComboBox() *CT_SdtComboBox {
	ret := &CT_SdtComboBox{}
	return ret
}

func (m *CT_SdtComboBox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LastValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastValue"},
			Value: fmt.Sprintf("%v", *m.LastValueAttr)})
	}
	e.EncodeToken(start)
	if m.ListItem != nil {
		selistItem := xml.StartElement{Name: xml.Name{Local: "w:listItem"}}
		for _, c := range m.ListItem {
			e.EncodeElement(c, selistItem)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SdtComboBox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "lastValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LastValueAttr = &parsed
			continue
		}
	}
lCT_SdtComboBox:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "listItem"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "listItem"}:
				tmp := NewCT_SdtListItem()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ListItem = append(m.ListItem, tmp)
			default:
				office.Log("skipping unsupported element on CT_SdtComboBox %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SdtComboBox
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SdtComboBox and its children
func (m *CT_SdtComboBox) Validate() error {
	return m.ValidateWithPath("CT_SdtComboBox")
}

// ValidateWithPath validates the CT_SdtComboBox and its children, prefixing error messages with path
func (m *CT_SdtComboBox) ValidateWithPath(path string) error {
	for i, v := range m.ListItem {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ListItem[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
