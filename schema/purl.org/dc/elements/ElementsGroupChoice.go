package elements

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type ElementsGroupChoice struct {
	Any []*Any
}

func NewElementsGroupChoice() *ElementsGroupChoice {
	ret := &ElementsGroupChoice{}
	return ret
}

func (m *ElementsGroupChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Any != nil {
		seany := xml.StartElement{Name: xml.Name{Local: "dc:any"}}
		for _, c := range m.Any {
			e.EncodeElement(c, seany)
		}
	}
	return nil
}

func (m *ElementsGroupChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lElementsGroupChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "any"}:
				tmp := NewAny()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Any = append(m.Any, tmp)
			default:
				office.Log("skipping unsupported element on ElementsGroupChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lElementsGroupChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ElementsGroupChoice and its children
func (m *ElementsGroupChoice) Validate() error {
	return m.ValidateWithPath("ElementsGroupChoice")
}

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (m *ElementsGroupChoice) ValidateWithPath(path string) error {
	for i, v := range m.Any {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Any[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
