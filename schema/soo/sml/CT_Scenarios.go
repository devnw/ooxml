package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_Scenarios struct {
	// Current Scenario
	CurrentAttr *uint32
	// Last Shown Scenario
	ShowAttr *uint32
	// Sequence of References
	SqrefAttr *ST_Sqref
	// Scenario
	Scenario []*CT_Scenario
}

func NewCT_Scenarios() *CT_Scenarios {
	ret := &CT_Scenarios{}
	return ret
}

func (m *CT_Scenarios) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CurrentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "current"},
			Value: fmt.Sprintf("%v", *m.CurrentAttr)})
	}
	if m.ShowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "show"},
			Value: fmt.Sprintf("%v", *m.ShowAttr)})
	}
	if m.SqrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sqref"},
			Value: fmt.Sprintf("%v", *m.SqrefAttr)})
	}
	e.EncodeToken(start)
	sescenario := xml.StartElement{Name: xml.Name{Local: "ma:scenario"}}
	for _, c := range m.Scenario {
		e.EncodeElement(c, sescenario)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Scenarios) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "current" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CurrentAttr = &pt
			continue
		}
		if attr.Name.Local == "show" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ShowAttr = &pt
			continue
		}
		if attr.Name.Local == "sqref" {
			parsed, err := ParseSliceST_Sqref(attr.Value)
			if err != nil {
				return err
			}
			m.SqrefAttr = &parsed
			continue
		}
	}
lCT_Scenarios:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "scenario"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "scenario"}:
				tmp := NewCT_Scenario()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Scenario = append(m.Scenario, tmp)
			default:
				office.Log("skipping unsupported element on CT_Scenarios %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Scenarios
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Scenarios and its children
func (m *CT_Scenarios) Validate() error {
	return m.ValidateWithPath("CT_Scenarios")
}

// ValidateWithPath validates the CT_Scenarios and its children, prefixing error messages with path
func (m *CT_Scenarios) ValidateWithPath(path string) error {
	for i, v := range m.Scenario {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Scenario[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
