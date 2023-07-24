package diagram

import (
	"encoding/xml"
	"fmt"
)

type CT_SDName struct {
	LangAttr *string
	ValAttr  string
}

func NewCT_SDName() *CT_SDName {
	ret := &CT_SDName{}
	return ret
}

func (m *CT_SDName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LangAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lang"},
			Value: fmt.Sprintf("%v", *m.LangAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SDName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "lang" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LangAttr = &parsed
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SDName: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SDName and its children
func (m *CT_SDName) Validate() error {
	return m.ValidateWithPath("CT_SDName")
}

// ValidateWithPath validates the CT_SDName and its children, prefixing error messages with path
func (m *CT_SDName) ValidateWithPath(path string) error {
	return nil
}
