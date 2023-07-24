package chart

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_BandFmts struct {
	BandFmt []*CT_BandFmt
}

func NewCT_BandFmts() *CT_BandFmts {
	ret := &CT_BandFmts{}
	return ret
}

func (m *CT_BandFmts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.BandFmt != nil {
		sebandFmt := xml.StartElement{Name: xml.Name{Local: "c:bandFmt"}}
		for _, c := range m.BandFmt {
			e.EncodeElement(c, sebandFmt)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BandFmts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BandFmts:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "bandFmt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "bandFmt"}:
				tmp := NewCT_BandFmt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BandFmt = append(m.BandFmt, tmp)
			default:
				office.Log("skipping unsupported element on CT_BandFmts %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BandFmts
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BandFmts and its children
func (m *CT_BandFmts) Validate() error {
	return m.ValidateWithPath("CT_BandFmts")
}

// ValidateWithPath validates the CT_BandFmts and its children, prefixing error messages with path
func (m *CT_BandFmts) ValidateWithPath(path string) error {
	for i, v := range m.BandFmt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BandFmt[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
