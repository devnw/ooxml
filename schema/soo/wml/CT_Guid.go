package wml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml/schema/soo/ofc/sharedTypes"
)

type CT_Guid struct {
	// GUID Value
	ValAttr *string
}

func NewCT_Guid() *CT_Guid {
	ret := &CT_Guid{}
	return ret
}

func (m *CT_Guid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Guid) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Guid: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Guid and its children
func (m *CT_Guid) Validate() error {
	return m.ValidateWithPath("CT_Guid")
}

// ValidateWithPath validates the CT_Guid and its children, prefixing error messages with path
func (m *CT_Guid) ValidateWithPath(path string) error {
	if m.ValAttr != nil {
		if !sharedTypes.ST_GuidPatternRe.MatchString(*m.ValAttr) {
			return fmt.Errorf(`%s/m.ValAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, *m.ValAttr)
		}
	}
	return nil
}
