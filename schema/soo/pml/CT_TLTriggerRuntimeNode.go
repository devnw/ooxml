package pml

import (
	"encoding/xml"
	"fmt"
)

type CT_TLTriggerRuntimeNode struct {
	// Value
	ValAttr ST_TLTriggerRuntimeNode
}

func NewCT_TLTriggerRuntimeNode() *CT_TLTriggerRuntimeNode {
	ret := &CT_TLTriggerRuntimeNode{}
	ret.ValAttr = ST_TLTriggerRuntimeNode(1)
	return ret
}

func (m *CT_TLTriggerRuntimeNode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTriggerRuntimeNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_TLTriggerRuntimeNode(1)
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
			return fmt.Errorf("parsing CT_TLTriggerRuntimeNode: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TLTriggerRuntimeNode and its children
func (m *CT_TLTriggerRuntimeNode) Validate() error {
	return m.ValidateWithPath("CT_TLTriggerRuntimeNode")
}

// ValidateWithPath validates the CT_TLTriggerRuntimeNode and its children, prefixing error messages with path
func (m *CT_TLTriggerRuntimeNode) ValidateWithPath(path string) error {
	if m.ValAttr == ST_TLTriggerRuntimeNodeUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
