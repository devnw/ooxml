package math

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_ManualBreak struct {
	AlnAtAttr *int64
}

func NewCT_ManualBreak() *CT_ManualBreak {
	ret := &CT_ManualBreak{}
	return ret
}

func (m *CT_ManualBreak) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AlnAtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "m:alnAt"},
			Value: fmt.Sprintf("%v", *m.AlnAtAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ManualBreak) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "alnAt" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.AlnAtAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ManualBreak: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ManualBreak and its children
func (m *CT_ManualBreak) Validate() error {
	return m.ValidateWithPath("CT_ManualBreak")
}

// ValidateWithPath validates the CT_ManualBreak and its children, prefixing error messages with path
func (m *CT_ManualBreak) ValidateWithPath(path string) error {
	if m.AlnAtAttr != nil {
		if *m.AlnAtAttr < 1 {
			return fmt.Errorf("%s/m.AlnAtAttr must be >= 1 (have %v)", path, *m.AlnAtAttr)
		}
		if *m.AlnAtAttr > 255 {
			return fmt.Errorf("%s/m.AlnAtAttr must be <= 255 (have %v)", path, *m.AlnAtAttr)
		}
	}
	return nil
}
