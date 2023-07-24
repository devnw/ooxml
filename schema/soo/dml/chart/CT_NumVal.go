package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_NumVal struct {
	IdxAttr        uint32
	FormatCodeAttr *string
	V              string
}

func NewCT_NumVal() *CT_NumVal {
	ret := &CT_NumVal{}
	return ret
}

func (m *CT_NumVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idx"},
		Value: fmt.Sprintf("%v", m.IdxAttr)})
	if m.FormatCodeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formatCode"},
			Value: fmt.Sprintf("%v", *m.FormatCodeAttr)})
	}
	e.EncodeToken(start)
	sev := xml.StartElement{Name: xml.Name{Local: "c:v"}}
	office.AddPreserveSpaceAttr(&sev, m.V)
	e.EncodeElement(m.V, sev)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NumVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "idx" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdxAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "formatCode" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FormatCodeAttr = &parsed
			continue
		}
	}
lCT_NumVal:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "v"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "v"}:
				if err := d.DecodeElement(&m.V, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_NumVal %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumVal
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NumVal and its children
func (m *CT_NumVal) Validate() error {
	return m.ValidateWithPath("CT_NumVal")
}

// ValidateWithPath validates the CT_NumVal and its children, prefixing error messages with path
func (m *CT_NumVal) ValidateWithPath(path string) error {
	return nil
}
