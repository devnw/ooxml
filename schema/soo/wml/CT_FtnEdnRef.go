package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml/schema/soo/ofc/sharedTypes"
)

type CT_FtnEdnRef struct {
	// Suppress Footnote/Endnote Reference Mark
	CustomMarkFollowsAttr *sharedTypes.ST_OnOff
	// Footnote/Endnote ID Reference
	IdAttr int64
}

func NewCT_FtnEdnRef() *CT_FtnEdnRef {
	ret := &CT_FtnEdnRef{}
	return ret
}

func (m *CT_FtnEdnRef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CustomMarkFollowsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:customMarkFollows"},
			Value: fmt.Sprintf("%v", *m.CustomMarkFollowsAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FtnEdnRef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "customMarkFollows" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.CustomMarkFollowsAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FtnEdnRef: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_FtnEdnRef and its children
func (m *CT_FtnEdnRef) Validate() error {
	return m.ValidateWithPath("CT_FtnEdnRef")
}

// ValidateWithPath validates the CT_FtnEdnRef and its children, prefixing error messages with path
func (m *CT_FtnEdnRef) ValidateWithPath(path string) error {
	if m.CustomMarkFollowsAttr != nil {
		if err := m.CustomMarkFollowsAttr.ValidateWithPath(path + "/CustomMarkFollowsAttr"); err != nil {
			return err
		}
	}
	return nil
}
