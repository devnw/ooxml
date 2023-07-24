package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_TextUnderlineLineFollowText struct {
}

func NewCT_TextUnderlineLineFollowText() *CT_TextUnderlineLineFollowText {
	ret := &CT_TextUnderlineLineFollowText{}
	return ret
}

func (m *CT_TextUnderlineLineFollowText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextUnderlineLineFollowText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextUnderlineLineFollowText: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextUnderlineLineFollowText and its children
func (m *CT_TextUnderlineLineFollowText) Validate() error {
	return m.ValidateWithPath("CT_TextUnderlineLineFollowText")
}

// ValidateWithPath validates the CT_TextUnderlineLineFollowText and its children, prefixing error messages with path
func (m *CT_TextUnderlineLineFollowText) ValidateWithPath(path string) error {
	return nil
}
