package sml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_ExternalLink struct {
	Choice *CT_ExternalLinkChoice
	ExtLst *CT_ExtensionList
}

func NewCT_ExternalLink() *CT_ExternalLink {
	ret := &CT_ExternalLink{}
	return ret
}

func (m *CT_ExternalLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ExternalLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ExternalLink:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "externalBook"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "externalBook"}:
				m.Choice = NewCT_ExternalLinkChoice()
				if err := d.DecodeElement(&m.Choice.ExternalBook, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "ddeLink"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "ddeLink"}:
				m.Choice = NewCT_ExternalLinkChoice()
				if err := d.DecodeElement(&m.Choice.DdeLink, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "oleLink"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "oleLink"}:
				m.Choice = NewCT_ExternalLinkChoice()
				if err := d.DecodeElement(&m.Choice.OleLink, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_ExternalLink %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalLink
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ExternalLink and its children
func (m *CT_ExternalLink) Validate() error {
	return m.ValidateWithPath("CT_ExternalLink")
}

// ValidateWithPath validates the CT_ExternalLink and its children, prefixing error messages with path
func (m *CT_ExternalLink) ValidateWithPath(path string) error {
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
