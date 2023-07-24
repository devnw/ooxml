package diagram

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_ColorTransformHeaderLst struct {
	ColorsDefHdr []*CT_ColorTransformHeader
}

func NewCT_ColorTransformHeaderLst() *CT_ColorTransformHeaderLst {
	ret := &CT_ColorTransformHeaderLst{}
	return ret
}

func (m *CT_ColorTransformHeaderLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ColorsDefHdr != nil {
		secolorsDefHdr := xml.StartElement{Name: xml.Name{Local: "colorsDefHdr"}}
		for _, c := range m.ColorsDefHdr {
			e.EncodeElement(c, secolorsDefHdr)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorTransformHeaderLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ColorTransformHeaderLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "colorsDefHdr"}:
				tmp := NewCT_ColorTransformHeader()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ColorsDefHdr = append(m.ColorsDefHdr, tmp)
			default:
				office.Log("skipping unsupported element on CT_ColorTransformHeaderLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorTransformHeaderLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorTransformHeaderLst and its children
func (m *CT_ColorTransformHeaderLst) Validate() error {
	return m.ValidateWithPath("CT_ColorTransformHeaderLst")
}

// ValidateWithPath validates the CT_ColorTransformHeaderLst and its children, prefixing error messages with path
func (m *CT_ColorTransformHeaderLst) ValidateWithPath(path string) error {
	for i, v := range m.ColorsDefHdr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ColorsDefHdr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
