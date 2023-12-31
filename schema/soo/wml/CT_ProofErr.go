package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_ProofErr struct {
	// Proofing Error Anchor Type
	TypeAttr ST_ProofErr
}

func NewCT_ProofErr() *CT_ProofErr {
	ret := &CT_ProofErr{}
	ret.TypeAttr = ST_ProofErr(1)
	return ret
}

func (m *CT_ProofErr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "w:type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ProofErr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_ProofErr(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ProofErr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ProofErr and its children
func (m *CT_ProofErr) Validate() error {
	return m.ValidateWithPath("CT_ProofErr")
}

// ValidateWithPath validates the CT_ProofErr and its children, prefixing error messages with path
func (m *CT_ProofErr) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_ProofErrUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
