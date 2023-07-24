package diagram

import (
	"encoding/xml"
	"fmt"
)

type CT_Parameter struct {
	TypeAttr ST_ParameterId
	ValAttr  ST_ParameterVal
}

func NewCT_Parameter() *CT_Parameter {
	ret := &CT_Parameter{}
	ret.TypeAttr = ST_ParameterId(1)
	return ret
}

func (m *CT_Parameter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Parameter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_ParameterId(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_ParameterVal(attr.Value)
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
			return fmt.Errorf("parsing CT_Parameter: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Parameter and its children
func (m *CT_Parameter) Validate() error {
	return m.ValidateWithPath("CT_Parameter")
}

// ValidateWithPath validates the CT_Parameter and its children, prefixing error messages with path
func (m *CT_Parameter) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_ParameterIdUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
