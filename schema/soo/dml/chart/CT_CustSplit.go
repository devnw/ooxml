package chart

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_CustSplit struct {
	SecondPiePt []*CT_UnsignedInt
}

func NewCT_CustSplit() *CT_CustSplit {
	ret := &CT_CustSplit{}
	return ret
}

func (m *CT_CustSplit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SecondPiePt != nil {
		sesecondPiePt := xml.StartElement{Name: xml.Name{Local: "c:secondPiePt"}}
		for _, c := range m.SecondPiePt {
			e.EncodeElement(c, sesecondPiePt)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustSplit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustSplit:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "secondPiePt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "secondPiePt"}:
				tmp := NewCT_UnsignedInt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SecondPiePt = append(m.SecondPiePt, tmp)
			default:
				office.Log("skipping unsupported element on CT_CustSplit %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustSplit
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustSplit and its children
func (m *CT_CustSplit) Validate() error {
	return m.ValidateWithPath("CT_CustSplit")
}

// ValidateWithPath validates the CT_CustSplit and its children, prefixing error messages with path
func (m *CT_CustSplit) ValidateWithPath(path string) error {
	for i, v := range m.SecondPiePt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SecondPiePt[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
