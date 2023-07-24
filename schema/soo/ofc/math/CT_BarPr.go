package math

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_BarPr struct {
	Pos    *CT_TopBot
	CtrlPr *CT_CtrlPr
}

func NewCT_BarPr() *CT_BarPr {
	ret := &CT_BarPr{}
	return ret
}

func (m *CT_BarPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Pos != nil {
		sepos := xml.StartElement{Name: xml.Name{Local: "m:pos"}}
		e.EncodeElement(m.Pos, sepos)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BarPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BarPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "pos"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "pos"}:
				m.Pos = NewCT_TopBot()
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_BarPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BarPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BarPr and its children
func (m *CT_BarPr) Validate() error {
	return m.ValidateWithPath("CT_BarPr")
}

// ValidateWithPath validates the CT_BarPr and its children, prefixing error messages with path
func (m *CT_BarPr) ValidateWithPath(path string) error {
	if m.Pos != nil {
		if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
