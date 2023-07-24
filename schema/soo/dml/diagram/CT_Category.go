package diagram

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Category struct {
	TypeAttr string
	PriAttr  uint32
}

func NewCT_Category() *CT_Category {
	ret := &CT_Category{}
	return ret
}

func (m *CT_Category) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"},
		Value: fmt.Sprintf("%v", m.TypeAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pri"},
		Value: fmt.Sprintf("%v", m.PriAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Category) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypeAttr = parsed
			continue
		}
		if attr.Name.Local == "pri" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.PriAttr = uint32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Category: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Category and its children
func (m *CT_Category) Validate() error {
	return m.ValidateWithPath("CT_Category")
}

// ValidateWithPath validates the CT_Category and its children, prefixing error messages with path
func (m *CT_Category) ValidateWithPath(path string) error {
	return nil
}
