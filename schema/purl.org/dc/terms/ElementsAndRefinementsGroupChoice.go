package terms

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/purl.org/dc/elements"
)

type ElementsAndRefinementsGroupChoice struct {
	Any []*elements.Any
}

func NewElementsAndRefinementsGroupChoice() *ElementsAndRefinementsGroupChoice {
	ret := &ElementsAndRefinementsGroupChoice{}
	return ret
}

func (m *ElementsAndRefinementsGroupChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Any != nil {
		seany := xml.StartElement{Name: xml.Name{Local: "dc:any"}}
		for _, c := range m.Any {
			e.EncodeElement(c, seany)
		}
	}
	return nil
}

func (m *ElementsAndRefinementsGroupChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lElementsAndRefinementsGroupChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "any"}:
				tmp := elements.NewAny()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Any = append(m.Any, tmp)
			default:
				office.Log("skipping unsupported element on ElementsAndRefinementsGroupChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lElementsAndRefinementsGroupChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (m *ElementsAndRefinementsGroupChoice) Validate() error {
	return m.ValidateWithPath("ElementsAndRefinementsGroupChoice")
}

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (m *ElementsAndRefinementsGroupChoice) ValidateWithPath(path string) error {
	for i, v := range m.Any {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Any[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
