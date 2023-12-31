package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_NumRestart struct {
	// Automatic Numbering Restart Value
	ValAttr ST_RestartNumber
}

func NewCT_NumRestart() *CT_NumRestart {
	ret := &CT_NumRestart{}
	ret.ValAttr = ST_RestartNumber(1)
	return ret
}

func (m *CT_NumRestart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NumRestart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_RestartNumber(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_NumRestart: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_NumRestart and its children
func (m *CT_NumRestart) Validate() error {
	return m.ValidateWithPath("CT_NumRestart")
}

// ValidateWithPath validates the CT_NumRestart and its children, prefixing error messages with path
func (m *CT_NumRestart) ValidateWithPath(path string) error {
	if m.ValAttr == ST_RestartNumberUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
