package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type AG_Password struct {
	AlgorithmNameAttr *string
	HashValueAttr     *string
	SaltValueAttr     *string
	SpinCountAttr     *int64
}

func NewAG_Password() *AG_Password {
	ret := &AG_Password{}
	return ret
}

func (m *AG_Password) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AlgorithmNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:algorithmName"},
			Value: fmt.Sprintf("%v", *m.AlgorithmNameAttr)})
	}
	if m.HashValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hashValue"},
			Value: fmt.Sprintf("%v", *m.HashValueAttr)})
	}
	if m.SaltValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:saltValue"},
			Value: fmt.Sprintf("%v", *m.SaltValueAttr)})
	}
	if m.SpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:spinCount"},
			Value: fmt.Sprintf("%v", *m.SpinCountAttr)})
	}
	return nil
}

func (m *AG_Password) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "algorithmName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgorithmNameAttr = &parsed
			continue
		}
		if attr.Name.Local == "hashValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HashValueAttr = &parsed
			continue
		}
		if attr.Name.Local == "saltValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SaltValueAttr = &parsed
			continue
		}
		if attr.Name.Local == "spinCount" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.SpinCountAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Password: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Password and its children
func (m *AG_Password) Validate() error {
	return m.ValidateWithPath("AG_Password")
}

// ValidateWithPath validates the AG_Password and its children, prefixing error messages with path
func (m *AG_Password) ValidateWithPath(path string) error {
	return nil
}
