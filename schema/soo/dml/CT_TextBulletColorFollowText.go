package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_TextBulletColorFollowText struct {
}

func NewCT_TextBulletColorFollowText() *CT_TextBulletColorFollowText {
	ret := &CT_TextBulletColorFollowText{}
	return ret
}

func (m *CT_TextBulletColorFollowText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextBulletColorFollowText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextBulletColorFollowText: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextBulletColorFollowText and its children
func (m *CT_TextBulletColorFollowText) Validate() error {
	return m.ValidateWithPath("CT_TextBulletColorFollowText")
}

// ValidateWithPath validates the CT_TextBulletColorFollowText and its children, prefixing error messages with path
func (m *CT_TextBulletColorFollowText) ValidateWithPath(path string) error {
	return nil
}
