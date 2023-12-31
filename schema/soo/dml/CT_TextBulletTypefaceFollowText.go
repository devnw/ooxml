package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_TextBulletTypefaceFollowText struct {
}

func NewCT_TextBulletTypefaceFollowText() *CT_TextBulletTypefaceFollowText {
	ret := &CT_TextBulletTypefaceFollowText{}
	return ret
}

func (m *CT_TextBulletTypefaceFollowText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextBulletTypefaceFollowText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextBulletTypefaceFollowText: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextBulletTypefaceFollowText and its children
func (m *CT_TextBulletTypefaceFollowText) Validate() error {
	return m.ValidateWithPath("CT_TextBulletTypefaceFollowText")
}

// ValidateWithPath validates the CT_TextBulletTypefaceFollowText and its children, prefixing error messages with path
func (m *CT_TextBulletTypefaceFollowText) ValidateWithPath(path string) error {
	return nil
}
