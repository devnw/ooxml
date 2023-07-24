package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_AudioCD struct {
	St     *CT_AudioCDTime
	End    *CT_AudioCDTime
	ExtLst *CT_OfficeArtExtensionList
}

func NewCT_AudioCD() *CT_AudioCD {
	ret := &CT_AudioCD{}
	ret.St = NewCT_AudioCDTime()
	ret.End = NewCT_AudioCDTime()
	return ret
}

func (m *CT_AudioCD) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sest := xml.StartElement{Name: xml.Name{Local: "a:st"}}
	e.EncodeElement(m.St, sest)
	seend := xml.StartElement{Name: xml.Name{Local: "a:end"}}
	e.EncodeElement(m.End, seend)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AudioCD) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.St = NewCT_AudioCDTime()
	m.End = NewCT_AudioCDTime()
lCT_AudioCD:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "st"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "st"}:
				if err := d.DecodeElement(m.St, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "end"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "end"}:
				if err := d.DecodeElement(m.End, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_AudioCD %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AudioCD
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AudioCD and its children
func (m *CT_AudioCD) Validate() error {
	return m.ValidateWithPath("CT_AudioCD")
}

// ValidateWithPath validates the CT_AudioCD and its children, prefixing error messages with path
func (m *CT_AudioCD) ValidateWithPath(path string) error {
	if err := m.St.ValidateWithPath(path + "/St"); err != nil {
		return err
	}
	if err := m.End.ValidateWithPath(path + "/End"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
