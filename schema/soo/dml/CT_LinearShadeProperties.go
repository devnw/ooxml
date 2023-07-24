package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_LinearShadeProperties struct {
	AngAttr    *int32
	ScaledAttr *bool
}

func NewCT_LinearShadeProperties() *CT_LinearShadeProperties {
	ret := &CT_LinearShadeProperties{}
	return ret
}

func (m *CT_LinearShadeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AngAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ang"},
			Value: fmt.Sprintf("%v", *m.AngAttr)})
	}
	if m.ScaledAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "scaled"},
			Value: fmt.Sprintf("%d", b2i(*m.ScaledAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LinearShadeProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ang" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.AngAttr = &pt
			continue
		}
		if attr.Name.Local == "scaled" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ScaledAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LinearShadeProperties: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LinearShadeProperties and its children
func (m *CT_LinearShadeProperties) Validate() error {
	return m.ValidateWithPath("CT_LinearShadeProperties")
}

// ValidateWithPath validates the CT_LinearShadeProperties and its children, prefixing error messages with path
func (m *CT_LinearShadeProperties) ValidateWithPath(path string) error {
	if m.AngAttr != nil {
		if *m.AngAttr < 0 {
			return fmt.Errorf("%s/m.AngAttr must be >= 0 (have %v)", path, *m.AngAttr)
		}
		if *m.AngAttr >= 21600000 {
			return fmt.Errorf("%s/m.AngAttr must be < 21600000 (have %v)", path, *m.AngAttr)
		}
	}
	return nil
}
