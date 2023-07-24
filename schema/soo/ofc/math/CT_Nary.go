package math

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_Nary struct {
	NaryPr *CT_NaryPr
	Sub    *CT_OMathArg
	Sup    *CT_OMathArg
	E      *CT_OMathArg
}

func NewCT_Nary() *CT_Nary {
	ret := &CT_Nary{}
	ret.Sub = NewCT_OMathArg()
	ret.Sup = NewCT_OMathArg()
	ret.E = NewCT_OMathArg()
	return ret
}

func (m *CT_Nary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.NaryPr != nil {
		senaryPr := xml.StartElement{Name: xml.Name{Local: "m:naryPr"}}
		e.EncodeElement(m.NaryPr, senaryPr)
	}
	sesub := xml.StartElement{Name: xml.Name{Local: "m:sub"}}
	e.EncodeElement(m.Sub, sesub)
	sesup := xml.StartElement{Name: xml.Name{Local: "m:sup"}}
	e.EncodeElement(m.Sup, sesup)
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Nary) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Sub = NewCT_OMathArg()
	m.Sup = NewCT_OMathArg()
	m.E = NewCT_OMathArg()
lCT_Nary:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "naryPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "naryPr"}:
				m.NaryPr = NewCT_NaryPr()
				if err := d.DecodeElement(m.NaryPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sub"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sub"}:
				if err := d.DecodeElement(m.Sub, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sup"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sup"}:
				if err := d.DecodeElement(m.Sup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "e"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "e"}:
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_Nary %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Nary
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Nary and its children
func (m *CT_Nary) Validate() error {
	return m.ValidateWithPath("CT_Nary")
}

// ValidateWithPath validates the CT_Nary and its children, prefixing error messages with path
func (m *CT_Nary) ValidateWithPath(path string) error {
	if m.NaryPr != nil {
		if err := m.NaryPr.ValidateWithPath(path + "/NaryPr"); err != nil {
			return err
		}
	}
	if err := m.Sub.ValidateWithPath(path + "/Sub"); err != nil {
		return err
	}
	if err := m.Sup.ValidateWithPath(path + "/Sup"); err != nil {
		return err
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	return nil
}
