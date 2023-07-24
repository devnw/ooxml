package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_LineJoinBevel struct {
}

func NewCT_LineJoinBevel() *CT_LineJoinBevel {
	ret := &CT_LineJoinBevel{}
	return ret
}

func (m *CT_LineJoinBevel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LineJoinBevel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LineJoinBevel: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LineJoinBevel and its children
func (m *CT_LineJoinBevel) Validate() error {
	return m.ValidateWithPath("CT_LineJoinBevel")
}

// ValidateWithPath validates the CT_LineJoinBevel and its children, prefixing error messages with path
func (m *CT_LineJoinBevel) ValidateWithPath(path string) error {
	return nil
}
