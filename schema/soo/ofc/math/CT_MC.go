package math

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_MC struct {
	McPr *CT_MCPr
}

func NewCT_MC() *CT_MC {
	ret := &CT_MC{}
	return ret
}

func (m *CT_MC) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.McPr != nil {
		semcPr := xml.StartElement{Name: xml.Name{Local: "m:mcPr"}}
		e.EncodeElement(m.McPr, semcPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MC) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MC:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "mcPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "mcPr"}:
				m.McPr = NewCT_MCPr()
				if err := d.DecodeElement(m.McPr, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_MC %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MC
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MC and its children
func (m *CT_MC) Validate() error {
	return m.ValidateWithPath("CT_MC")
}

// ValidateWithPath validates the CT_MC and its children, prefixing error messages with path
func (m *CT_MC) ValidateWithPath(path string) error {
	if m.McPr != nil {
		if err := m.McPr.ValidateWithPath(path + "/McPr"); err != nil {
			return err
		}
	}
	return nil
}
