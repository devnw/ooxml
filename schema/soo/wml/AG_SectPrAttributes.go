package wml

import (
	"encoding/xml"
	"fmt"
)

type AG_SectPrAttributes struct {
	RsidRPrAttr  *string
	RsidDelAttr  *string
	RsidRAttr    *string
	RsidSectAttr *string
}

func NewAG_SectPrAttributes() *AG_SectPrAttributes {
	ret := &AG_SectPrAttributes{}
	return ret
}

func (m *AG_SectPrAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RsidRPrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRPr"},
			Value: fmt.Sprintf("%v", *m.RsidRPrAttr)})
	}
	if m.RsidDelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidDel"},
			Value: fmt.Sprintf("%v", *m.RsidDelAttr)})
	}
	if m.RsidRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidR"},
			Value: fmt.Sprintf("%v", *m.RsidRAttr)})
	}
	if m.RsidSectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidSect"},
			Value: fmt.Sprintf("%v", *m.RsidSectAttr)})
	}
	return nil
}

func (m *AG_SectPrAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rsidRPr" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidRPrAttr = &parsed
			continue
		}
		if attr.Name.Local == "rsidDel" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidDelAttr = &parsed
			continue
		}
		if attr.Name.Local == "rsidR" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidRAttr = &parsed
			continue
		}
		if attr.Name.Local == "rsidSect" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidSectAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_SectPrAttributes: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_SectPrAttributes and its children
func (m *AG_SectPrAttributes) Validate() error {
	return m.ValidateWithPath("AG_SectPrAttributes")
}

// ValidateWithPath validates the AG_SectPrAttributes and its children, prefixing error messages with path
func (m *AG_SectPrAttributes) ValidateWithPath(path string) error {
	return nil
}
