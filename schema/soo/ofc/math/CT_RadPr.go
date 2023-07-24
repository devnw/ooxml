package math

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_RadPr struct {
	DegHide *CT_OnOff
	CtrlPr  *CT_CtrlPr
}

func NewCT_RadPr() *CT_RadPr {
	ret := &CT_RadPr{}
	return ret
}

func (m *CT_RadPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DegHide != nil {
		sedegHide := xml.StartElement{Name: xml.Name{Local: "m:degHide"}}
		e.EncodeElement(m.DegHide, sedegHide)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RadPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_RadPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "degHide"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "degHide"}:
				m.DegHide = NewCT_OnOff()
				if err := d.DecodeElement(m.DegHide, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_RadPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RadPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RadPr and its children
func (m *CT_RadPr) Validate() error {
	return m.ValidateWithPath("CT_RadPr")
}

// ValidateWithPath validates the CT_RadPr and its children, prefixing error messages with path
func (m *CT_RadPr) ValidateWithPath(path string) error {
	if m.DegHide != nil {
		if err := m.DegHide.ValidateWithPath(path + "/DegHide"); err != nil {
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
