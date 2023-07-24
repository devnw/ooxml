package wml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_DocPartBehaviors struct {
	// Entry Insertion Behavior
	Behavior []*CT_DocPartBehavior
}

func NewCT_DocPartBehaviors() *CT_DocPartBehaviors {
	ret := &CT_DocPartBehaviors{}
	return ret
}

func (m *CT_DocPartBehaviors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Behavior != nil {
		sebehavior := xml.StartElement{Name: xml.Name{Local: "w:behavior"}}
		for _, c := range m.Behavior {
			e.EncodeElement(c, sebehavior)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocPartBehaviors) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocPartBehaviors:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "behavior"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "behavior"}:
				tmp := NewCT_DocPartBehavior()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Behavior = append(m.Behavior, tmp)
			default:
				office.Log("skipping unsupported element on CT_DocPartBehaviors %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocPartBehaviors
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DocPartBehaviors and its children
func (m *CT_DocPartBehaviors) Validate() error {
	return m.ValidateWithPath("CT_DocPartBehaviors")
}

// ValidateWithPath validates the CT_DocPartBehaviors and its children, prefixing error messages with path
func (m *CT_DocPartBehaviors) ValidateWithPath(path string) error {
	for i, v := range m.Behavior {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Behavior[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
